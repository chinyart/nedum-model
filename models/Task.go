package models

import (
	"fmt"

	db "github.com/chinyart/nedum-model/config"
	u "github.com/chinyart/nedum-model/utils"
	"go.mongodb.org/mongo-driver/mongo/primitive"
	"go.mongodb.org/mongo-driver/mongo/bson"
)

//Task to manage the task
type Task struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Body   string             `json:"body,omitempty" bson:"body,omitempty"`
	Date   string             `json:"date,omitempty" bson:"date,omitempty"`
	UserId uint               `json:"user_id,omitempty" bson:"user_id,omitempty"` //The user that this contact belongs to
}

func test() {
	db.DBConnect()
}

//Validate to validate the entry
func (task *Task) Validate() (map[string]interface{}, bool) {
	if task.Title == "" {
		return u.Message(false, "Title should be on the payload"), false
	}

	if task.Body == "" {
		return u.Message(false, "Body should be on the payload"), false
	}

	if task.Date == "" {
		return u.Message(false, "Date should be on the payload"), false
	}

	// if task.UserId <= 0 {
	// 	return u.Message(false, "User ID is not recognized"), false
	// }

	//All the required parameters are present
	return u.Message(true, "success"), true
}

//Create to create the task
func (task *Task) Create() map[string]interface{} {
	resp, ok := task.Validate()

	if ok == true {
		return resp
	}

	db.DBConnect()

	respon := u.Message(true, "success")
	respon["task"] = task
	return respon
}

//GetTask to get one task
func GetTask(id uint) *Task {
	task := &Task{}
	err := GetDB().Table("tasks").Where("id = ?", id).First(task).Error
	if err != nil {
		return nil
	}
	return task
}

//GetTasks to get all the tasks
func GetTasks(user uint) []*Task {

	task := make([]*Task, 0)
	err := GetDB().Table("tasks").Where("user_id = ?", user).Find(&task).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return task
}
