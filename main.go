package main

import (
	//	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
)

var (
	strings = map[string]struct{}{}

	StrLen = len(strings)
)

func init() {
	for _, v := range []string{
		"butterfly",
		"hero",
		"action",
		"painting",
		"selfie",
		"nerd",
		"violet",
		"red",
		"blue",
		"summer",
		"rain",
		"neon",
		"sign",
		"life",
		"cool",
		"thoughtful",
		"trainstation",
		"cliffside",
		"midnight",
		"summercamp",
		"presidential",
		"eloquent",
	} {
		strings[v] = struct{}{}
	}
}

func main() {
	s, err := GenPaths(2, 3, 5)
	if err != nil {
		log.Println(err)
	}

	for _, str := range s {
		fmt.Println(str)
	}
	fmt.Printf("Total strings: %d", StrLen)
}

func GenPaths(n ...int) ([]string, error) {

	//rand.Seed(time.Now().UnixNano())

	var (
		s      = make([][]string, len(n))
		output = []string{"/"}
	)

	for i, num := range n {

		for i2 := 0; i2 < num; i2++ {
			for rndStr, _ := range strings {
				s[i] = append(s[i], rndStr)
				delete(strings, rndStr)
				break
			}
		}
	}

	spew.Dump(s)

	for i := range s {
		for i2 := range s[i] {
			output = append(output, word+"/"+s[i][i2])
		}
	}

	return output, nil
}

func possibleCombinations() {

}
