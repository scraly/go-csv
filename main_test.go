package main

import "testing"

//pas besoin de redéclarer la struct Person car elle existe dans la clase main que l'on souhaite tester

func TestGetNumberOfAttendees(t *testing.T) {
	var people []Person
	people = InjectData()

	nb := GetNumberOfAttendees(people)

	if nb != 77 {
		t.Errorf("Nb of attendees is incorrect, got %d, want %d.", nb, 77)
	}

}

func InjectData() []Person {
	var people []Person
	max := 78
	i := 1
	for i < max {
		people = append(people, Person{
			Name:  "toto",
			Email: "toto@toto.com",
		})
		i++
	}
	return people
}
