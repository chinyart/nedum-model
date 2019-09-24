package controllers

import (
	"github.com/chinyart/nedum-model/models"
)

//DisplayPerson to display person details
func DisplayPerson() (string, error) {
	var show = models.PersonBio{
		Name: "chidi",
	}
	return show.Name, nil
}
