package service

import "terminal/models"

func Hello() (string, bool) {
	resp := "Hello World"
	if err := models.DB.Row().Err(); err != nil {
		return resp, true
	} else {
		return resp, false
	}
}
