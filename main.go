package main

import (
	"io/ioutil"
	"log"
	"os"

	markdown "github.com/WhoMeNope/k8b-render/markdown"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Requires one argument: path to a file to preview.")
	}

	// read file to preview
	path := os.Args[1]

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatalln(path + " : does not exist.")
	}
	if info.IsDir() {
		log.Fatalln(path + " : is a directory.")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// render and serve
	log.Print(string(data))

	rendered, err := markdown.Render(data)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(rendered.String())
}
