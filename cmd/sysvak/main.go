package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jessevdk/go-flags"

	"github.com/borud/sysvak/pkg/sysvak"
)

var opt struct {
	From               string `short:"f" long:"from"   description:"from date (defaults to 1 week ago)"`
	To                 string `short:"t" long:"to"     description:"to date (defaults to now)"`
	Doses              string `short:"d" long:"dose"   description:"which doeses"  default:"1,2"`
	Municipalities     string `short:"m" long:"municipality"     description:"Municipality code(s)"`
	Genders            string `short:"g" long:"gender" description:"genders" default:"M,K"`
	Ages               string `short:"a" long:"age"    description:"age ranges, comma separated" default:"1,2,3,4,5,6,7"`
	Format             string `short:"o" long:"output" description:"output format" choice:"json" choice:"csv" choice:"table" choice:"markdown" choice:"html" default:"table"`
	NoColor            bool   `short:"n" long:"no-color" description:"turn off colors in table output"`
	ListMunicipalities bool   `short:"l" long:"list-muni" description:"list municipality codes"`
	ListAges           bool   `short:"x" long:"list-ages" description:"list age ranges"`
	Verbose            bool   `short:"v" description:"Verbose mode"`
}

func main() {
	p := flags.NewParser(&opt, flags.Default)
	_, err := p.Parse()
	if err != nil {
		return
	}

	if opt.ListMunicipalities {
		for k, v := range sysvak.MunicipalityByCode {
			fmt.Printf("%s - %s\n", k, v.Name)
		}
		return
	}

	if opt.ListAges {
		for k, v := range sysvak.AgeRangeToString {
			fmt.Printf("%d - %s\n", k, v)
		}
		return
	}

	if len(opt.Municipalities) == 0 {
		log.Print("Please list what municipalities you want data for")
		log.Print("(to list municipalities run with -l or --list-muni)")
		return
	}

	var from = time.Now().Add(-7 * 24 * time.Hour)
	var to = time.Now()

	if opt.From != "" {
		from, err = time.Parse("2006-01-02", opt.From)
		if err != nil {
			log.Fatalf("error parsing --from date: %v", err)
		}
	}

	if opt.To != "" {
		to, err = time.Parse("2006-01-02", opt.To)
		if err != nil {
			log.Fatalf("error parsing --to date: %v", err)
		}
	}

	q := sysvak.Query{
		From:           from,
		To:             to,
		Doses:          strings.Split(opt.Doses, ","),
		Municipalities: strings.Split(opt.Municipalities, ","),
		Genders:        strings.Split(opt.Genders, ","),
		Ages:           strings.Split(opt.Ages, ","),
	}

	if opt.Verbose {
		url, err := q.AsURL()
		log.Printf("URL = %s %v", url, err)
	}

	r, err := sysvak.Lookup(q)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	sort.Slice(r, func(i, j int) bool {
		ageComp := r[i].Age - r[j].Age
		switch {
		case ageComp < 0:
			return true
		case ageComp > 0:
			return false
		}

		genderComp := r[i].Gender - r[j].Gender
		switch genderComp {
		case -1:
			return true
		case 1:
			return false
		}
		return r[i].Dose < r[j].Dose
	})

	switch strings.ToLower(opt.Format) {
	case "json":
		printJSON(r)
	case "csv":
		printCSV(r)
	case "table":
		printTable(r, from, to)
	case "markdown":
		printTable(r, from, to)
	case "html":
		printTable(r, from, to)
	default:
		log.Printf("unknown output '%s'", opt.Format)
	}
}

func printJSON(results []sysvak.Result) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(results)
}

func printCSV(results []sysvak.Result) {
	fmt.Printf("%s,%s,%s,%s\n", "AGE", "DOSE", "GENDER", "COUNT")
	for _, r := range results {
		fmt.Printf("\"%s\",%d,\"%s\",%d\n", sysvak.AgeRangeToString[r.Age], r.Dose, sysvak.GenderToString[r.Gender], r.Count)
	}
}

func printTable(results []sysvak.Result, from time.Time, to time.Time) {
	sum := 0

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetIndexColumn(1)
	if !opt.NoColor {
		t.SetStyle(table.StyleColoredBlackOnRedWhite)
	}

	t.SetTitle("Covid-19 vaccines\nfrom %s\nto %s (%d days)",
		from.Format("2006-01-02"),
		to.Format("2006-01-02"),
		int(to.Sub(from).Hours()/24))

	t.AppendHeader(table.Row{"Age", "Municipality", "Dose", "Gender", "Count"})

	lastAge := 0
	for _, r := range results {
		age := sysvak.AgeRangeToString[r.Age]
		if r.Age == lastAge {
			age = ""
		}
		t.AppendRow([]interface{}{
			age, r.Where, r.Dose, sysvak.GenderToString[r.Gender], r.Count,
		})
		lastAge = r.Age
		sum += int(r.Count)
	}

	t.AppendSeparator()
	t.AppendFooter(table.Row{"", "", "Total", sum})

	switch opt.Format {
	case "html":
		t.RenderHTML()
	case "markdown":
		t.RenderMarkdown()
	default:
		fmt.Println()
		t.Render()
		fmt.Println()
	}

}
