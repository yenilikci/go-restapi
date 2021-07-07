package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "../dataloaders"
	. "../models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	//data loaders
	users := LoadUsers()
	interests := LoadInterests()
	interestsMappings := LoadInterestMapping()
	//işlem
	var newUsers []User
	for _, user := range users {
		for _, interestMapping := range interestsMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
		fmt.Println(user.FirstName)
	}
	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel) //go object -> json
	w.Write([]byte(data))
}
