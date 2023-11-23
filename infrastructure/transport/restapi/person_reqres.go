package restapi

import "demoapp/application"

type PersonReqRes struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewPersonReqResFromPersonOutput(person *application.PersonOutput) *PersonReqRes {
	return (*PersonReqRes)(person)
}

func (p *PersonReqRes) ToPersonOutput() *application.PersonOutput {
	return (*application.PersonOutput)(p)
}
