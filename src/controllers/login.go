package controllers

import (
	"net/http"
	"social-app/src/utils"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplates(w, "login.html", nil)
}
