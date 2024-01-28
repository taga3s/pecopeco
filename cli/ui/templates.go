package ui

import "github.com/manifoldco/promptui"

func DefaultPromptTemplate() *promptui.PromptTemplates {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . }} ",
		Success: "{{ . | bold }} ",
	}
	return templates
}
