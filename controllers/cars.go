package controllers

import (
	"encoding/json"
	"github.com/timereversal/keymaxweb/models"
	"net/http"
)

type Car struct {
	CarService *models.CarService
}

func (c *Car) Create(w http.ResponseWriter, r *http.Response) {
	var input struct {
		brand string `json:"brand"`
		model string `json:"model"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {

	}
}
