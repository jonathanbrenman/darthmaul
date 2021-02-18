package menue

import "fmt"

type Options interface {
	Show()
	ShowLogo()
}

type menue struct {
	AppLogo string
}

func NewMenu() Options {
	return &menue{
		AppLogo: "______           _   _    ___  ___            _ \n|  _  \\         | | | |   |  \\/  |           | |\n| | | |__ _ _ __| |_| |__ | .  . | __ _ _   _| |\n| | | / _` | '__| __| '_ \\| |\\/| |/ _` | | | | |\n| |/ / (_| | |  | |_| | | | |  | | (_| | |_| | |\n|___/ \\____|_|   \\__|_| |_\\_|  |_/\\____|\\____|_|\n                                                \n",
	}
}

func (m menue) Show() {
	fmt.Println("Version 1.0.0")
	fmt.Println("Generate command:")
	fmt.Println("generate controller <name>")
	fmt.Println("generate service <name>")
	fmt.Println("generate respository <name>")
	fmt.Println("\n")
}

func (m menue) ShowLogo() {
	fmt.Println(m.AppLogo)
}