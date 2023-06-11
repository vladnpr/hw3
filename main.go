package main

import (
	"go_course/hw3/character"
	"go_course/hw3/scenario"
)

func main() {
	mainChar := character.NewMainCharacter("Steven", 18)
	scenario.Run(mainChar)
}
