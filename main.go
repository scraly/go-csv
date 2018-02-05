package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

//Person reperesent an attendee
type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	csvFile, _ := os.Open("2018_bigdata.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	//1. Afficher la liste des participants
	// fmt.Println("Participants:")
	// fmt.Println("-----------------------")
	var people []Person
	people = RetrieveAttendees(people, reader)

	//2. Afficher le nombre de participants
	nbOfAttendees := GetNumberOfAttendees(people)
	fmt.Println("Nb of attendees: ", nbOfAttendees)

	//3. faire un random de 1 au nb de participants et afficher le gagnant
	rand.Seed(time.Now().UnixNano())
	winnerInt := rand.Intn(nbOfAttendees)
	fmt.Println("...And the winner is: ", people[winnerInt].Name)
}

func RetrieveAttendees(people []Person, reader *csv.Reader) []Person {
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(line[0])

		people = append(people, Person{
			Name:  line[0],
			Email: line[1],
		})
	}
	return people
}

func GetNumberOfAttendees(people []Person) int {
	return len(people)
}
