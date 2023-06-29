package iactools

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func GetCommand() (method string, err error) {
	getCommandPrompt := promptui.Select{
		Label:        "Select command",
		Items:        []string{"Create"},
		HideSelected: true,
	}

	_, command, err := getCommandPrompt.Run()

	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}

	return command, nil
}
