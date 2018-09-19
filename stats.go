package main

import (
	"fmt"
	"io/ioutil"
	"log"
	//"regexp"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("There are", len(files), "files.")

	//var editors []string
	var authors []string

	authors, _ = parseFiles(files)

	authors_count := countDuplicates(authors)

	for k, v := range authors_count {
		fmt.Printf("Item : %s , Count : %d\n", k, v)
	}
}

func parseFiles(files []os.FileInfo) (authors []string, editors []string) {
	//var editors []string
	//var authors []string

	for _, file := range files {
		parts := strings.Split(file.Name(), ".")
		ext := parts[1]
		filename := parts[0]

		fmt.Println(filename, ext)

		info_parts := strings.Split(filename, " - ")

		if len(info_parts) < 2 {
			fmt.Println("Skipping ", file.Name(), "")
			continue
		}

		//date := info_parts[0]
		author_editor := info_parts[1]

		editor := strings.Split(author_editor, "-")[1]
		author := strings.Split(author_editor, "-")[0]

		fmt.Println("Author:", author)
		fmt.Println("Editor:", editor)

		//author_names := regexp.MustCompile("[, &]+").Split(author, -1)
		author_names := strings.FieldsFunc(author, Split)

		fmt.Println("Authors:", author_names)
		fmt.Println("There are", len(author_names), "authors")

		editors = append(editors, editor)
		authors = append(authors, author_names...)

		// author := info_parts[1].split('-')[0]
		// authors_names := re.split('[, &]', author)
		// print authors_names

		fmt.Println(file.Name())
	}

	return //authors, editors
}

func countDuplicates(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exists in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func Split(r rune) bool {
	return r == ' ' || r == '&' || r == ',' || r == '1' || r == '2'
}
