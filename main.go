package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	fmt.Println(string(data))

	// render
	rendered, err := markdown.NewRenderer().Render(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("RENDERED AS:\n\n")
	fmt.Println(string(rendered))

	// serve
	http.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(rendered)
		},
	))

	fmt.Println("SERVING AT:\nhttp://localhost:5000/")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
