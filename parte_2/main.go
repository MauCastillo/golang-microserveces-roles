package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	pathOutput         = "output.json"
	fieldRole          = "rol"
	fieldUser          = "usuario"
	fieldOrganizations = "organizacion"
	maxRowItem         = 3
)

func createUsersOrganization(data [][]string) UsersOrganization {
	organizations := map[string]users{}

	labels := map[string]int{}

	for index, organization := range data {
		if index == 0 || len(organization) < maxRowItem {
			labels = getLabel(organization)
			continue
		}

		organizations = marshalDocument(organization, organizations, labels)
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

func marshalDocument(organization []string, organizations map[string]users, labels map[string]int) map[string]users {
	username := organization[labels[fieldUser]]
	company := organization[labels[fieldOrganizations]]
	userRole := organization[labels[fieldRole]]

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

func readFile() ([][]string, error) {
	file, err := os.Open(pathFileCSV)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func usersOrganizationToString(usersOrganization UsersOrganization) (string, error) {
	organizer, err := json.MarshalIndent(usersOrganization, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(pathOutput, organizer, 0644)
	if err != nil {
		return "", err
	}
	return string(organizer), nil
}

func buildJSON() (string, error) {
	data, err := readFile()
	if err != nil {
		return "", err
	}

	usersOrganization := createUsersOrganization(data)

	return usersOrganizationToString(usersOrganization)
}

func getLabel(labelRow []string) map[string]int {
	labels := make(map[string]int)

	for index, value := range labelRow {
		labels[value] = index
	}

	return labels
}

func main() {
	document, err := buildJSON()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(document)
}
