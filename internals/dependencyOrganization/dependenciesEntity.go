package dependencyorganization

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
)

type dependencyStructureOrganization struct {
	Mapper map[string]string
}

func NewdependencyStructureOrganization() *dependencyStructureOrganization {
	return &dependencyStructureOrganization{
		Mapper: nil,
	}
}

func (d *dependencyStructureOrganization) ChoiceOne(args []string) string {

	choice := args

	var choiceType string

	prompt := &survey.Select{

		Message: fmt.Sprintf("Choose%s", choice),
		Options: choice,
	}

	err := survey.AskOne(prompt, &choiceType, nil)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return choiceType
}

func (d *dependencyStructureOrganization) ChoiceDependencies(keys, values []string, topicName string) []string {
	choice := keys

	var choiceType []string

	prompt := &survey.MultiSelect{
		Message: topicName,
		Options: choice,
	}

	err := survey.AskOne(prompt, &choiceType, nil)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return choiceType
}
