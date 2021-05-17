package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/borud/sysvak/pkg/sysvak"
	"github.com/jessevdk/go-flags"
)

var opt struct {
	From           string   `short:"f" long:"from"   description:"From date (defaults to 1 week ago)"`
	To             string   `short:"t" long:"to"     description:"To date (defaults to now)"`
	Doses          []string `short:"d" long:"dose"   description:"Which doeses"  default:"01,02"`
	Municipalities []string `short:"m" long:"municipality"     description:"Municipality code(s)"`
	Genders        []string `short:"g" long:"gender" description:"Genders"       default:"M,K"`
	Ages           []string `short:"a" long:"age"    description:"Age ranges, comma separated" default:"1,2,3,4,5,6,7"`
	Format         string   `short:"o" long:"output" description:"Output format" choice:"json" choice:"csv" choice:"table" default:"table"`
}

func main() {
	p := flags.NewParser(&opt, flags.Default)
	_, err := p.Parse()
	if err != nil {
		return
	}

	if len(opt.Municipalities) != 0 && opt.Municipalities[0] == "?" {
		for k, v := range sysvak.MunicipalityByCode {
			fmt.Printf("%s - %s\n", k, v.Name)
		}
		return
	}

	if len(opt.Ages) != 0 && opt.Ages[0] == "?" {
		for k, v := range sysvak.AgeRange {
			fmt.Printf("%d - %s\n", k, v)
		}
		return
	}

	if len(opt.Municipalities) == 0 {
		log.Fatal("Please list what municipalities you want data for")
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
		Doses:          opt.Doses,
		Municipalities: opt.Municipalities,
		Genders:        opt.Genders,
		Ages:           opt.Ages,
	}

	r, err := sysvak.Lookup(q)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("date range %s to %s", from.Format("2006-01-02"), to.Format("2006-01-02"))

	switch strings.ToLower(opt.Format) {
	case "json":
		printJSON(r)
	case "csv":
		printCSV(r)
	case "table":
		printTable(r)
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
	for _, r := range results {
		fmt.Printf("%s;%s;%d\n", r.Description, r.Where, r.Count)
	}
}

func printTable(results []sysvak.Result) {
	sum := 0
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	fmt.Println()
	fmt.Fprintf(w, "%s\t%s\t%s\n", "DESCRIPTION", "WHERE", " COUNT")
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\t%6d\n", r.Description, r.Where, r.Count)
		sum += int(r.Count)
	}
	fmt.Fprintf(w, "SUM\t%s\t%6d\n\n", "", sum)
	w.Flush()
}
