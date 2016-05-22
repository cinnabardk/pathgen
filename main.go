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

type paths struct {
	name string
	p    []paths
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

	var call = func(root string, s []string) []string {
		for _, str := range s {
			output = append(output, root+"/"+str)
		}
		return output[len(output)-len(s):]
	}

	new := s[0]
	call("", new)

	spew.Dump(new)

	for i := 1; i < len(n); i++ {
		for _, str := range new {
			new = call(str, s[i])
		}
	}

	return output, nil
}

func possibleCombinations() {

}
