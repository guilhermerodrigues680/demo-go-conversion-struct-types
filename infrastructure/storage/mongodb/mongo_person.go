package mongodb

import "demoapp"

type MongoPerson struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func NewMongoPersonFromPerson(person *demoapp.Person) *MongoPerson {
	return (*MongoPerson)(person)
}

func (p *MongoPerson) ToPerson() *demoapp.Person {
	return (*demoapp.Person)(p)
}
