package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

//import re

func main() {
	data, _ := os.ReadFile("001 miracle-in-the-andes.txt")
	ldata := string(data)

	//with open("001 miracle-in-the-andes.txt") as file:
	//	book = file.read()

	// fmt.Println(string(data))

	// re, _ := regexp.Compile("[A-Z]{1}[^.]*[^a-zA-Z]+love[^a-zA-Z]+[^.]*.")
	// sentences := re.FindAllString(ldata, -1)
	// fmt.Println(len(sentences))
	// for _, v := range sentences {
	// 	fmt.Println(v)
	// }

	wordsCount := make(map[string]int)

	re, _ := regexp.Compile("[a-z]+")
	words := re.FindAllString(strings.ToLower(ldata), -1)

	// 	pattern = re.compile("[a-z]+")
	// words = re.findall(pattern, book.lower())

	names := make([]string, 0, len(wordsCount))
	for _, word := range words {
		if _, p := wordsCount[word]; p == false {
			wordsCount[word] = 1
			names = append(names, word)
		} else {
			wordsCount[word] += 1
		}
	}

	sort.Slice(names, func(i, j int) bool {
		return wordsCount[names[i]] > wordsCount[names[j]]
	})

	// 	d = {}

	// for word in words:
	//     if word in d.keys():
	//         d[word] = d[word] + 1
	//     else:
	//         d[word] = 1

	// dlist = [(value,key) for (key, value) in d.items()]
	// sorted(dlist, reverse=True)

	for _, name := range names[:10] {
		fmt.Println(name, wordsCount[name])
	}

	// fmt.Printf("%+v", wordsCount)

}
