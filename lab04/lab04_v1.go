package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type DataReader interface {
	UnmarshalJSON(data []byte) error
}

type DataAdder interface {
	AddData(data interface{})
}

type Base struct {
	Offers    []Offer  `json:"offers"`
	People    []Person `json:"people"`
	Employees []Employee
}

type Offers struct {
	Offers []Offer `json:"offers"`
}

type Offer struct {
	Name      string    `json:"name"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	Education string    `json:"education"`
}

type People struct {
	People []Person `json:"people"`
}

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Education string `json:"education"`
}

type Employee struct {
	Person Person
	Offer  Offer
}

func (o *Offers) UnmarshalJSON(data []byte) error {
	var offers struct {
		Offers []Offer `json:"offers"`
	}
	err := json.Unmarshal(data, &offers)
	if err != nil {
		return err
	}
	o.Offers = offers.Offers
	return nil
}

func (p *People) UnmarshalJSON(data []byte) error {
	var people struct {
		People []Person `json:"people"`
	}
	err := json.Unmarshal(data, &people)
	if err != nil {
		return err
	}
	p.People = people.People
	return nil
}

func (b *Base) InitializeFromFile(fileName string, reader DataReader) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = reader.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	return nil
}

func (b *Base) AddOffer(offer Offer) {
	b.Offers = append(b.Offers, offer)
}

func (b *Base) AddPerson(person Person) {
	b.People = append(b.People, person)
}

func (b *Base) MatchOffersToPeople() {
	for _, person := range b.People {
		for _, offer := range b.Offers {
			if person.Education == offer.Education {
				employee := Employee{Person: person, Offer: offer}
				b.Employees = append(b.Employees, employee)
			}
		}
	}
}

func (b *Base) CalculateCost() {
	totalCost := 0
	fmt.Printf("Employees that we will hire:\n")
	for _, employee := range b.Employees {
		fmt.Printf("Name: %s\n", employee.Person.Name)
		totalCost += 1000 // Sample cost
	}
	fmt.Println("Total cost of employment:", totalCost)
}

func (b *Base) PrintBase() {
	fmt.Println("Offers:")
	for _, offer := range b.Offers {
		fmt.Printf("Name: %s, From: %s, To: %s, Education: %s\n", offer.Name, offer.From, offer.To, offer.Education)
	}

	fmt.Println("Persons:")
	for _, person := range b.People {
		fmt.Printf("Name: %s, Age: %d, Education: %s, Offer: %s\n", person.Name, person.Age, person.Education)
	}

	fmt.Println("Employees:")
	for _, employee := range b.Employees {
		fmt.Printf("Name: %s, Age: %d, Education: %s, Offer: %s\n", employee.Person.Name, employee.Person.Age, employee.Person.Education, employee.Offer.Name)
	}
}

func main() {
	base := Base{}

	offersReader := &Offers{}
	err := base.InitializeFromFile("Offers.json", offersReader)
	if err != nil {
		fmt.Println("Error initializing offers:", err)
		return
	}

	peopleReader := &People{}
	err = base.InitializeFromFile("People.json", peopleReader)
	if err != nil {
		fmt.Println("Error initializing people:", err)
		return
	}

	for _, offer := range offersReader.Offers {
		base.Offers = append(base.Offers, offer)
	}

	for _, person := range peopleReader.People {
		base.People = append(base.People, person)
		// for _, offer := range base.Offers {
		// 	if person.Education == offer.Education {
		// 		employee := Employee{Person: person, Offer: offer}
		// 		base.Employees = append(base.Employees, employee)
		// 	}
		// }
	}

	base.MatchOffersToPeople()
	base.PrintBase()

	// Przykład dodawania oferty
	newOffer := Offer{Name: "Tester", From: time.Now(), To: time.Now().AddDate(0, 1, 0), Education: "e"}
	base.AddOffer(newOffer)
	fmt.Println("\nAfter adding new offer:")
	base.PrintBase()

	// Przykład dodawania osoby
	newPerson := Person{Name: "John Doe", Age: 30, Education: "h"}
	base.AddPerson(newPerson)
	fmt.Println("\nAfter adding new person:")
	base.MatchOffersToPeople()
	base.PrintBase()

	// Przykład obliczania kosztu, stala kwota 1000
	base.CalculateCost()
}
