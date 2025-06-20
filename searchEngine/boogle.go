package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var docs map[int]string = map[int]string{
	1: "Dogs, cats are pet animals",
	2: "cats are good thoug",
}

func maxVal(s []int) (top int) {
	top = s[0]
	for ind := range s {
		if s[ind] > top {
			top = s[ind]
		}
	}
	return

}

func add(doc map[int]string, text string) {
	var sliceKey []int
	for key := range doc {
		sliceKey = append(sliceKey, key)
	}
	maxKey := maxVal(sliceKey)
	doc[maxKey+1] = text
}

func QueryParser(query string) (sliceVal []string) {
	strSlice := strings.Split(query, " ")
	if len(strSlice) != 3 {
		fmt.Println("Query is not valid.")
		return

	}
	first_val := strings.ToLower(strSlice[0])
	bool_val := strings.ToLower(strSlice[1])
	last_val := strings.ToLower(strSlice[2])
	switch bool_val {
	case "and":
		for i := range docs {
			if strings.Contains(strings.ToLower(docs[i]), first_val) && strings.Contains(strings.ToLower(docs[i]), last_val) {
				sliceVal = append(sliceVal, docs[i])
			}
		}

	case "or":
		for i := range docs {
			if strings.Contains(strings.ToLower(docs[i]), first_val) || strings.Contains(strings.ToLower(docs[i]), last_val) {
				sliceVal = append(sliceVal, docs[i])
			}
		}
	case "not":
		for i := range docs {
			if strings.Contains(strings.ToLower(docs[i]), first_val) && !(strings.Contains(strings.ToLower(docs[i]), last_val)) {
				sliceVal = append(sliceVal, docs[i])
			}
		}
	default:
		for i := range docs {
			if !(strings.Contains(strings.ToLower(docs[i]), first_val)) || !(strings.Contains(strings.ToLower(docs[i]), last_val)) {
				sliceVal = append(sliceVal, docs[i])
			}
		}
	}
	return
}

func Search() {
	text := bufio.NewScanner(os.Stdin)
	fmt.Print("-> ")
	text.Scan()
	input := text.Text()

	val := QueryParser(input)
	for _, res := range val {
		fmt.Println(res)
	}
}

func main() {
	fmt.Println("Welcome to Boogle! 🪄")
	add(docs, "Python is a scripting language")
	add(docs, "Java and Pythonn are helpful")
	Search()
}
