package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	embeddingCmd := flag.NewFlagSet("embedding", flag.ExitOnError)
	embModelName := embeddingCmd.String("model", "text-embeddings-3-small", "Choose the appropriate embedding model")
	embInput := embeddingCmd.String("input", "Break down this string into some embedded representation", "Input text to generate embeddings")

	chatCmd := flag.NewFlagSet("chat", flag.ExitOnError)
	chatModelName := chatCmd.String("model", "text-embeddings-3-small", "Choose Embeddings Model")
	chatInput := chatCmd.String("input", "", "Input chat string")

	if len(os.Args) < 2 {
		fmt.Printf("Expected the appropriate subcommand.")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "embedding":

		embeddingCmd.Parse(os.Args[2:])

		//embedding processing function file goes here
		fmt.Println(*embModelName)
		fmt.Println(*embInput)

	case "chat":

		chatCmd.Parse(os.Args[2:])
		// chat processing function .go file goes here
		fmt.Println(*chatModelName)
		fmt.Println(*chatInput)

	}

}
