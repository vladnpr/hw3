package safeMiniGame

import (
	"bufio"
	"fmt"
	"go_course/hw3/textOutput"
	"os"
)

type Safe struct {
	Output textOutput.Output
}

func (s Safe) CodeTry() string {
	combination := "59"
	s.Output.TextOutput(fmt.Sprintln("Введіть комбінацію із двох цифр коду: "))
	scanner := bufio.NewScanner(os.Stdin)
	var userInput string

	for {
		if !scanner.Scan() {
			break
		}
		userInput = scanner.Text()
		if userInput == combination {
			s.Output.TextOutput(fmt.Sprintf("Поздоровляю! Код вірний\n\n"))
			break
		} else {
			s.Output.TextOutput(fmt.Sprintf("Код не вірний. Спробуй ще\n\n"))
		}
	}

	return userInput
}
