package character

type MainCharacter struct {
	Name string
	Age  int
}

func NewMainCharacter(name string, age int) MainCharacter {
	return MainCharacter{
		Name: name, Age: age,
	}
}
