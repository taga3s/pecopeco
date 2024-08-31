package search

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/manifoldco/promptui"
	"github.com/taga3s/pecopeco-cli/api/model"
	uiutil "github.com/taga3s/pecopeco-cli/ui/util"
)

type searchConditionInput struct {
	City  string
	Genre string
}

func GetConditionInput(genreList []model.Genre) searchConditionInput {
	promptForCity := promptui.Prompt{
		Label: "> Which city? (Japanese only, ex.渋谷)",
		Validate: func(input string) error {
			if utf8.RuneCountInString(input) == 0 {
				return errors.New("please enter a city")
			}
			if strings.TrimSpace(input) == "" || strings.Contains(input, " ") {
				return errors.New("city cannot be only whitespace")
			}
			return nil
		},
		Templates: uiutil.DefaultPromptTemplate(),
	}
	city, err := promptForCity.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return searchConditionInput{}
	}

	genreMap := make(map[string]model.Genre)
	options := make([]string, 0, len(genreList))
	for _, v := range genreList {
		genreMap[v.Name] = v
		options = append(options, v.Name)
	}

	promptForGenre := promptui.Select{
		Label: "> What genre?",
		Items: options,
	}
	_, genre, err := promptForGenre.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return searchConditionInput{}
	}

	return searchConditionInput{City: strings.TrimSpace(city), Genre: strings.TrimSpace(genreMap[genre].Code)}
}
