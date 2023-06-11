package character

type MainCharacter struct {
	Name string
	Age  int
}

func NewMainCharacter(name string) MainCharacter {
	return MainCharacter{
		Name: name,
	}
}
