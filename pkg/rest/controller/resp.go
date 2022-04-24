package controller

import "github.com/radekkrejcirik01/radek-golang-backend/pkg/model/users"

type resp struct {
	Status  string       `json:""`
	Message string       `json:",omitempty"`
	Data    *[]users.USERS `json:",omitempty"`
}
