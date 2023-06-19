package main

import (
	"fmt"
	"go_course/hw3/character"
	"go_course/hw3/newWorldScenario"
)

func main() {
	var name string

	fmt.Println("Введіть ім'я персонажу: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err.Error())
	}
	mainChar := character.NewMainCharacter(name)
	scenario := newWorldScenario.NWScenario{MainChar: mainChar}
	scenario.Run()
}
