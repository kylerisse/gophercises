package main

import (
	"bufio"
	"fmt"
	"os"

	"./urlparse"
)

func main() {
	fileList := []string{"ex1.html", "ex2.html", "ex3.html", "ex4.html", "ex5.html"}
	for i := range fileList {
		file, err := os.Open(fileList[i])
		if err != nil {
			fmt.Println("cannot open file " + fileList[i])
			os.Exit(1)
		}
		fmt.Println("[" + fileList[i] + "]")
		buf := bufio.NewReader(file)
		var links []urlparse.Link
		links, err = urlparse.ExtractLinks(buf)
		if err != nil {
			fmt.Println(err.Error())
		}
		for j := range links {
			fmt.Println("Href: '" + links[j].Href + "' Text: '" + links[j].Text + "'")
		}
	}
	
}