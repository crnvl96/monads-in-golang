package operations

import "github.com/crnvl96/monads-in-golang/option_monad/monads"

type Person struct {
	name string
	age  int
}

var people = []Person{
	{"John", 30},
	{"Jane", 20},
	{"Bob", 15},
	{"Mary", 40},
}

type Pet struct {
	name  string
	owner Person
}

var pets = []Pet{
	{"Rex", people[0]},
	{"Lassie", people[1]},
	{"Snoopy", people[2]},
	{"Garfield", people[3]},
}

type PetAge struct {
	pet Pet
	age int
}

var petAges = []PetAge{
	{pets[0], 5},
	{pets[1], 10},
	{pets[2], 3},
	{pets[3], 7},
}

// These must return the monadic WrapperType
func GetPerson(person Person) monads.Option[Person] {
	for _, p := range people {
		if p == person {
			return monads.Option[Person]{Result: p, IsValid: true}
		}
	}
	return monads.Option[Person]{Result: Person{}, IsValid: false}
}

func GetPet(person Person) Pet {
	for _, p := range pets {
		if p.owner == person {
			return p
		}
	}
	return Pet{}
}

func GetPetAge(pet Pet) PetAge {
	for _, p := range petAges {
		if p.pet == pet {
			return p
		}
	}
	return PetAge{}
}
