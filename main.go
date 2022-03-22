package main

import (
	"fmt"
	"log"
)

func main() {
	var userTerm string

	fmt.Printf("Enter the software you'd like to find threats for: ")
	_, err := fmt.Scan(&userTerm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Searching for", userTerm)
	//	down()
	//	read()
	//	write()
}

// ask for term (from software list options)
// Start with Golang, Rust, TypeScript, Java, JavaScript, Azure, +15 total
// search csv for term and return resulting key words
// have list of threat types??? choose per result? There's already systematic ways of doing this
