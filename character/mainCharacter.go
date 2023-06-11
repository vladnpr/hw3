package character

import "go_course/hw3/item"

type MainCharacter struct {
	Name      string
	Age       int
	Inventory [10]item.Item
}

func NewMainCharacter(name string, age int) MainCharacter {
	return MainCharacter{
		name, age, [10]item.Item{},
	}
}
