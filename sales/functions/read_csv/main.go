package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type UsersOrganization []company

type usersCompany struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

type company struct {
	Organization string         `json:"organization"`
	Users        []usersCompany `json:"users"`
}

type users struct {
	userRole map[string][]string
}

const (
	pathFileCSV        = "sample.csv"
	fieldRole          = 2
	fieldUser          = 1
	fieldOrganizations = 0
	maxRowItem         = 3
)

func createUsersOrganization(data [][]string) UsersOrganization {
	organizations := map[string]users{}

	for index, organization := range data {
		if index == 0 || len(organization) < maxRowItem {
			continue
		}

		organizations = marshalDocument(organization, organizations)
	}

	usersOrganization := UsersOrganization{}
	companyUser := []usersCompany{}

	for name, organization := range organizations {
		for userName, roles := range organization.userRole {
			companyUser = append(companyUser, usersCompany{Username: userName, Roles: roles})
		}

		usersOrganization = append(usersOrganization, company{Organization: name, Users: companyUser})
		companyUser = []usersCompany{}
	}

	return usersOrganization
}

func marshalDocument(organization []string, organizations map[string]users) map[string]users {
	username := organization[fieldUser]
	company := organization[fieldOrganizations]
	userRole := organization[fieldRole]

	if user, ok := organizations[company]; ok {
		if role, ok := user.userRole[username]; ok {
			organizations[company].userRole[username] = append(role, userRole)
			return organizations
		}

		organizations[company].userRole[username] = []string{userRole}
		return organizations
	}

	organizations[company] = users{userRole: map[string][]string{}}
	organizations[company].userRole[username] = append(organizations[company].userRole[username], userRole)

	return organizations
}

func main() {
	file, err := os.Open(pathFileCSV)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	usersOrganization := createUsersOrganization(data)

	organizer, err := json.Marshal(usersOrganization)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(organizer))
}
