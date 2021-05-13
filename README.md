# sysvak

Simple utility for digging out vaccine data from SYSVAK.

## Install

    go get -u github.com/borud/sysvak/cmd/sysvak

## Command line options

```
Usage:
  sysvak [OPTIONS]

Application Options:
  -b, --from=                   From date (defaults to 1 week ago)
  -e, --to=                     To date (defaults to now)
  -d, --dose=[01|02]            Which doeses (default: 01,02)
  -m, --mu=                     Municipality code(s)
  -g, --gender=                 Genders (default: M,K)
  -a, --age=                    Age ranges, comma separated (default: 1,2,3,4,5,6,7)
  -f, --format=[json|csv|table] Output format (default: table)

Help Options:
  -h, --help                    Show this help message
```

To list municipalities use `-m ?` and to list age ranges `-a ?`.
