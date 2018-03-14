package cmd

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

func getPreviousOption(d prompt.Document) (cmd, option string, found bool) {
	args := strings.Split(d.TextBeforeCursor(), " ")
	l := len(args)
	if l >= 2 {
		option = args[l-2]
	}
	if strings.HasPrefix(option, "-") {
		return args[0], option, true
	}
	return "", "", false
}

func completeOptionArguments(d prompt.Document) ([]prompt.Suggest, bool) {
	cmd, option, found := getPreviousOption(d)
	if !found {
		return []prompt.Suggest{}, false
	}
	switch cmd {
	case "build":
		switch option {
		case "-f", "--file":
			return fileCompleter(d), true
		}
	case "deploy":
		switch option {
		case "--bundle-file", "-c", "--compose-file":
			return fileCompleter(d), true
		}
	case "export":
		switch option {
		case "-o", "--output":
			return fileCompleter(d), true
		}
	case "save":
		switch option {
		case "-o", "--output":
			return fileCompleter(d), true
		}
	}
	return []prompt.Suggest{}, false
}

func fileCompleter(d prompt.Document) []prompt.Suggest {
	path := d.GetWordBeforeCursor()
	if strings.HasPrefix(path, "./") {
		path = path[2:]
	}
	dir := filepath.Dir(path)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Print("ERROR catch error " + err.Error())
		return []prompt.Suggest{}
	}
	suggests := make([]prompt.Suggest, 0, len(files))
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		suggests = append(suggests, prompt.Suggest{Text: filepath.Join(dir, f.Name())})
	}
	return prompt.FilterHasPrefix(suggests, path, false)
}
