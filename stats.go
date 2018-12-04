package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	authors_count := summarise(authors)

	for k, v := range authors_count {
		fmt.Printf("Item : %s , Count : %d\n", k, v)
	}
}

func parseFiles(files []os.FileInfo) (authors []string, editors []string) {
	// Iterate through each filename, making sure that the name matches
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
		author_names := strings.FieldsFunc(author, SplitNames)

		fmt.Println("Authors:", author_names)
		fmt.Println("There are", len(author_names), "authors")

		editors = append(editors, editor)
		authors = append(authors, author_names...)

		// author := info_parts[1].split('-')[0]
		// authors_names := re.split('[, &]', author)
		// print authors_names

		fmt.Println(file.Name())
	}

	return
}

func summarise(list []string) map[string]int {

	item_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exists in the item_frequency map

		_, exist := item_frequency[item]

		if exist {
			item_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			item_frequency[item] = 1 // else start counting from 1
		}
	}
	return item_frequency
}

func SplitNames(r rune) bool {
	// If the same translator translates two documents on their own on the same date,
	// the filenames will contain the author's name followed by a number,
	// which must be stripped away:
	//   Author 1
	//   Author 2
	//
	// Depending on how many translators there are, the names will be listed in
	// slightly different ways. If there are two authors, they will appear as:
	//   Author & Author
	//
	// If there are 3 or more authors, they will appear as:
	//   Author, Author & Author
	return r == ' ' || r == '&' || r == ',' || r == '1' || r == '2'
}
