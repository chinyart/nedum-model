package controllers

import "github.com/chinyart/nedum-model/models"

//DisplayPerson to display person details
func DisplayPersonName(personBio models.PersonBio) (string, error) {
	//personBio.Name = "chidi"
	return personBio.Name, nil
}
