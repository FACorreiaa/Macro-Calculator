package menu

import (
	"errors"
	"strconv"

	"fmt"

	"github.com/manifoldco/promptui"
)

func GetInputPrompt(label string) (float64, error) {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:       label,
		Default:     "",
		AllowEdit:   false,
		Validate:    validate,
		Mask:        0,
		HideEntered: false,
		Templates:   &promptui.PromptTemplates{Prompt: "{{ . }} ", Valid: "{{ . | bold }} ", Invalid: "{{ . | red | bold }} "},
		IsConfirm:   false,
		IsVimMode:   false,
		Stdin:       nil,
		Stdout:      nil,
	}

	result, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	data, err := strconv.ParseFloat(result, 64)
	if err != nil {
		return 0, err
	}

	return data, nil

}

func GetSelectMenu(label string, options []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: options,
		Size:  5,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U0001F336 {{ . | cyan }}",
			Inactive: "  {{ . | faint }}",
			Selected: "\U0001F336 {{ . | red | bold }}",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("You choose %q\n", result)
	return result
}
