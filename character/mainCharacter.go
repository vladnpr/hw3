package character

import "go_course/hw3/item"

type MainCharacter struct {
	Name      string
	Age       int
	Inventory []item.Item
}

func (mc *MainCharacter) AddToInventory(items ...item.Item) {
	mc.Inventory = append(mc.Inventory, items...)
}

func NewMainCharacter(name string, age int) MainCharacter {
	return MainCharacter{
		Name: name, Age: age,
	}
}
