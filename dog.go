package dog

import "fmt"

func Jump() string {
	return "The dog jumps"
}

type Dog struct {
	Name string
	Race string
	Age int
}

func (d Dog) GetDogCredentials() string {
	creds := fmt.Sprintf("The dog's name is %s. It's race is %s and is %d years old.", d.Name, d.Race, d.Age)
	return creds
}