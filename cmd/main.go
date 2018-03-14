package cmd

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

const (
	// AppName - the name of the application.
	AppName = "docker-prompt"
)

// Main starts application
func Main() {
	fmt.Printf("%s %s (rev-%s)\n", AppName, Version, ShortCommitID)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
	defer fmt.Println("Bye!")
	title := fmt.Sprintf("%s: interactive kubernetes client", AppName)
	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle(title),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
	)
	p.Run()
}
