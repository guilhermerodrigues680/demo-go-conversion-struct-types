package application

import "demoapp"

type PersonOutput struct {
	Name string
	Age  int
}

func NewPersonOutputFromPerson(person *demoapp.Person) *PersonOutput {
	return (*PersonOutput)(person)
}

func (p *PersonOutput) ToPerson() *demoapp.Person {
	return (*demoapp.Person)(p)
}
