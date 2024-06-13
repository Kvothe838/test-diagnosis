package models

import "fmt"

type Patient struct {
	ID       string
	Name     string
	Surname  string
	Document Document
	Contacts []Contact
}

func (p Patient) GetFullName() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

type Document struct {
	ID   int
	Info string
	Type DocumentType
}

type DocumentType struct {
	ID   int
	Name string
}

type Contact struct {
	Type ContactType
	Info string
}

type ContactType struct {
	ID   int
	Name string
}
