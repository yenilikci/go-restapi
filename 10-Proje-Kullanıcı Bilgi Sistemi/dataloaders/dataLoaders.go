package dataloaders

import (
	. "../models" //alter: m "../models", m.User...
	util "../utils"
	"encoding/json"
)

//json -> go object
func LoadUsers() []User {
	bytes, _ := util.ReadFile("../json/users.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterests() []Interest {
	bytes, _ := util.ReadFile("../json/interests.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterestMapping() []InterestMapping {
	bytes, _ := util.ReadFile("../json/userInterestMapping.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
