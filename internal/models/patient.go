package models

import "fmt"

type Patient struct {
	ID      string
	Name    string
	Surname string
}

func (p Patient) GetFullName() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}
