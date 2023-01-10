package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const asp = "./akademi_spellings.txt"

func loadAkademiSpellings(path string) []string {
	var result []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		result = append(result, s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return result

}

func MatchString(array []string, target string) int {
	if i, found := sort.Find(len(array), func(i int) int {
		return strings.Compare(target, array[i])

	}); found {
		return i
	}
	return 0
}

func main() {
	words := loadAkademiSpellings(asp)
	testword := "ভূত"
	a := MatchString(words, testword)

	fmt.Printf("Found the word '%s' in index %d\n\n", testword, a)

	if a != 0 {
		fmt.Printf("The found word at index %d is -> %s\n\n", a, words[a])
	}

	fmt.Printf("Found %d words\n", len(words))
	fmt.Println("BananJani")
}
