package sysvak

// AgeRangeToString ...
var AgeRangeToString = map[int]string{
	1: "0-15",
	2: "16-44",
	3: "45-54",
	4: "55-64",
	5: "65-74",
	6: "75-84",
	7: "> 85",
}

// StringToAgeRange ...
var StringToAgeRange = map[string]int{
	"0-15 år":    1,
	"16-44 år":   2,
	"45-54 år":   3,
	"55-64 år":   4,
	"65-74 år":   5,
	"75-84 år":   6,
	"85 og over": 7,
}

// StringToDose ...
var StringToDose = map[string]int{
	"Dose 1": 1,
	"Dose 2": 2,
}

// StringToGender ...
var StringToGender = map[string]int{
	"Kvinne": 1,
	"Mann":   2,
}

// GenderToString ...
var GenderToString = map[int]string{
	1: "Kvinne",
	2: "Mann",
}
