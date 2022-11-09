package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

type film struct {
	Title string
	Year  int
}

type cinema struct {
	Films  []film
	People []string
	Header string
}

var myFc = template.FuncMap{
	"upper": strings.ToUpper,
	"inc":   increment,
}

func increment(n int) int {
	return n + 1
}

func init() {
	fmt.Println("When does it happen?")
}

func main() {

	// Create struct with data
	myCinema := cinema{
		Films: []film{
			{Title: "Godfather", Year: 1972},
			{Title: "The Dark Knight Rises", Year: 2012},
			{Title: "Matrix", Year: 1999},
		},
		People: []string{"John", "Peter", "Melinda", "Muddy"},
		Header: "The best films ever!!",
	}

	// Create a file for exporting template later
	myFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Unable to create a file")
	}
	defer myFile.Close()

	// Parse template with custom func and return a checked template
	myTemplate := template.Must(template.New("index.template").Funcs(myFc).ParseGlob("*.template"))

	// Pass the values to file
	err = myTemplate.Execute(myFile, myCinema)
	if err != nil {
		log.Fatalln("Unable to write to file:", err)
	}
	// Print them to standard output (since only there is one template)
	err = myTemplate.Execute(os.Stdout, myCinema)
	if err != nil {
		log.Fatalln("Unable to write to Stdout: ", err)
	}
}
