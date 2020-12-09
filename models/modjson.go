package models

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
)

type settings struct {
	Clients []Users
}

const settingsFilename = "settings.json"

func JsonFunk(database *sql.DB) {
	rawDataIn, err := ioutil.ReadFile(settingsFilename)
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}

	var settings settings
	err = json.Unmarshal(rawDataIn, &settings)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}

	newClient := Users{
		Id:       1,
		Admin:    true,
		Name:     "Jonibek",
		Surname:  "Mahmudov",
		Age:      22,
		Gender:   "man",
		Login:    "10071999",
		Password: "988777226",
		Remove:   false,
	}

	settings.Clients = append(settings.Clients, newClient)

	rawDataOut, err := json.MarshalIndent(&settings, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile(settingsFilename, rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}
}
