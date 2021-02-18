package main

import (
	"darthmaul/commands"
	"darthmaul/menue"
	"fmt"
	"os"
)

var (
	allowedEntities = []string{"controller","service","repository"}
)

func main() {
	args := os.Args[1:]
	menue := menue.NewMenu()
	menue.ShowLogo()

	// check if i have some args.
	if len(args) == 0 {
		menue.Show()
		os.Exit(0)
	}

	// validate and execute command.
	var cmd commands.Command
	switch args[0] {
		case "generate":
			if len(args) < 3 || !containsString(args[1]){
				showError(menue, "missing argument should be: darthmaul generate <entity> <name>")
			}
			cmd = commands.NewGenerateCMD(args[1], args[2])
		case "create-app":
			if  len(args) < 2 {
				showError(menue, "missing argument should be: darthmaul create-app <name>")
			}
			cmd = commands.NewCreateAppCMD(args[1])
		default:
			showError(menue, "command "+args[0]+" not valid.")
	}

	if err := cmd.Execute(); err != nil {
		showError(menue, err)
	}

	fmt.Println("Done.")
}

func showError(menue menue.Options, message interface{}) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	menue.Show()
	fmt.Println(string(colorRed), message)
	fmt.Println(string(colorReset))
	os.Exit(1)
}

func containsString(argEntity string) bool {
	for _, entity := range allowedEntities {
		if entity == argEntity {
			return true
		}
	}
	return false
}