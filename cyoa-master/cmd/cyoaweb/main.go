package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")

	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	fmt.Println(d)
}
