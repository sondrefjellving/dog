package dog

import "fmt"

func Jump() string {
	return "The dog jumps"
}

type Dog struct {
	name string
	race string
	age int
}

func (d Dog) GetDogCredentials() string {
	creds := fmt.Sprintf("The dog's name is %s. It's race is %s and is %d years old.", d.name, d.race, d.age)
	return creds
}