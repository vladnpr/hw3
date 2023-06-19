package textOutput

import "fmt"

type Output struct {
}

func (o Output) TextOutput(text string) {
	fmt.Println(text)
}
