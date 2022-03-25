package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuildJSON(t *testing.T) {
	c := require.New(t)

	_, err := buildJSON()
	c.NoError(err)

	data, err := readFile()
	c.NoError(err)

	organization := createUsersOrganization(data)
	c.Len(organization, 2)

	main()
}

func TestGetLabel(t *testing.T) {
	c := require.New(t)

	labelRow := getLabel([]string{"organizacion", "usuario", "rol"})
	c.Len(labelRow, 3)
}
