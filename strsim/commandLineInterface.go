package main

import (
	"bitbucket.org/dbarbera/wdeai_ex5/stringSimilarity"
	"fmt"
	"os"
)

const (
	argJ  = "jaro"
	argJw = "jaroWinkler"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("3 arguments required (metric, string1, string2)")
		return
	}
	var out float64
	s1 := args[2]
	s2 := args[3]
	switch args[1] {
	case argJ:
		out = stringSimilarity.Jaro(s1, s2)
	case argJw:
		out = stringSimilarity.JaroWinkler(s1, s2)
	default:
		fmt.Println("invalid metric")
		return
	}

	fmt.Println(out)
}
