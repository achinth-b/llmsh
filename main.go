package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	model := flag.String("model", "world", "Default model that we want to use for llmsh")
	input := flag.String("input", "You are a helpful assistant", "Input string that wishes to be sent to the model")

	flag.Parse()
	fmt.Printf(os.Args[0])

	if flag.NArg() == 0 {
		fmt.Printf("Hello, %s!\n", *model)
		fmt.Printf("Hello, %s!\n", *input)

	} else if flag.Arg(0) == "model" {

		if *model == "text-embeddings-3-small" {

			// HTTP Post request
			fmt.Printf(os.Getenv("apiKey"))

		}

	} else {
		fmt.Printf("Hello, %s!\n", *model)
	}

}
