package safeMiniGame

import (
	"bufio"
	"fmt"
	"os"
)

func CodeTry() string {
	combination := "59"
	fmt.Println("Введіть комбінацію із двох цифр коду: ")
	scanner := bufio.NewScanner(os.Stdin)
	var userInput string

	for {
		if !scanner.Scan() {
			break
		}
		userInput = scanner.Text()
		if userInput == combination {
			fmt.Printf("Поздоровляю! Код вірний\n\n")
			break
		} else {
			fmt.Printf("Код не вірний. Спробуй ще\n\n")
		}
	}

	return userInput
}
