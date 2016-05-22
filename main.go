package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	strings = map[string]struct{}{}

	StrLen = len(strings)
)

func init() {
	for _, v := range []string{"/",
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
	s, err := GenPaths(2, 3, 20)
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

	var s = make([][]string, len(n))
	var output = []string{strings[0]}
	var usedStrings = make([]int, len(n))

	for i, num := range n {
		if num > StrLen-i {
			return output, errors.New(fmt.Sprintf("int no %d with a value of %d has exceeded the maximum of %d available strings", i, num, StrLen-i))
		}

	start:
		rnd := rand.Intn(StrLen)
		for _, u := range usedStrings {
			if rnd == u {
				if rnd > StrLen {
					goto start
				} else {
					rnd++
				}
			}
		}
		usedStrings[i] = rnd

		s[i] = append(s[i], strings[rnd])
	}

	for i := range s {
		word := s[i][0]
		for i2 := range s[i] {
			output = append(output, word+"/"+s[i][i2])
		}
	}

	return output, nil
}

func possibleCombinations() {

}
