package main

import (
	"flag"
	"fmt"
	"os"

	"llmsh/commands"
)

type Command struct {
	Description string
	FlagSet     *flag.FlagSet
}

func main() {

	// name -> (description, FlagSet)
	// needed for list of all commadns and list of flags per command available
	commandMap := make(map[string]Command)

	helpCmd := Command{
		FlagSet: flag.NewFlagSet("help", flag.ExitOnError),
	}

	healthCmd := Command{
		Description: "Check API Keys if they are appropriately set",
		FlagSet:     flag.NewFlagSet("check-api-keys", flag.ExitOnError),
	}
	commandMap["check-api-keys"] = healthCmd

	embeddingCmd := Command{
		Description: "Yield an embedding for a candidate input text",
		FlagSet:     flag.NewFlagSet("embedding", flag.ExitOnError),
	}
	embModelName := embeddingCmd.FlagSet.String("model", "text-embedding-3-small", "Choose the appropriate embedding model")
	embInput := embeddingCmd.FlagSet.String("input", "Break down this string into some embedded representation", "Input text to generate embeddings")
	commandMap["embedding"] = embeddingCmd

	chatCmd := Command{
		Description: "Engage in chat with your chosen model",
		FlagSet:     flag.NewFlagSet("chat", flag.ExitOnError),
	}
	chatModelName := chatCmd.FlagSet.String("model", "gpt-3.5-turbo-0125", "Choose Chat Model")
	chatInput := chatCmd.FlagSet.String("input", "", "Input chat string")
	commandMap["chat"] = chatCmd

	if len(os.Args) < 2 {

		fmt.Println("\n")

		fmt.Println("Expected the appropriate command. Please check from the available subcommands below.")
		for name, command := range commandMap {
			fmt.Printf("  %s: %s\n", name, command.Description)
		}
		fmt.Println("\n")

		os.Exit(1)
	}

	switch os.Args[1] {

	case "check-api-keys":

		showHelpIfNeeded(*healthCmd.FlagSet)

		healthCmd.FlagSet.Parse(os.Args[2:])
		fmt.Printf(os.LookupEnv("OPENAI_API_KEY"))

		fmt.Printf("Successfully added OpenAI API key!")

	case "embedding":

		embeddingCmd.FlagSet.Parse(os.Args[2:])
		showHelpIfNeeded(*embeddingCmd.FlagSet)

		//embedding processing function file goes here
		embedding, err := commands.Embedding(embModelName, embInput)
		if err != nil {
			fmt.Printf("Please check this error: %v\n", err)
		}

		fmt.Printf("%v", embedding)

	case "chat":

		showHelpIfNeeded(*chatCmd.FlagSet)

		chatCmd.FlagSet.Parse(os.Args[2:])
		// chat processing function .go file goes here
		fmt.Println(*chatModelName)
		fmt.Println(*chatInput)

	default:
		helpCmd.FlagSet.Parse(os.Args[2:])
		for name, command := range commandMap {
			fmt.Printf("  %s: %s\n", name, command.Description)
		}
		os.Exit(1)
	}

}

func showHelpIfNeeded(FlagSet flag.FlagSet) {

	if len(os.Args[2:]) == 0 {
		FlagSet.VisitAll(func(f *flag.Flag) {
			fmt.Printf("  -%s: %s (default: %q)\n", f.Name, f.Usage, f.DefValue)
		})

		os.Exit(1)

	}

}
