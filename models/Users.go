package models

// PersonBio to manage person details
type PersonBio struct {
	Name string
	Age  int
}

// PersonRelationship to manage the relation of thr person
type PersonRelationship struct {
	Status   string
	Relation string
}

// Person the final person's information
type Person struct {
	Bio          PersonBio
	Relationship PersonRelationship
}
