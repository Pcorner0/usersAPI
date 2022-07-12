package controller

import (
	"github.com/pcorner0/usersAPI/database"
	"github.com/pcorner0/usersAPI/models"
	"github.com/pcorner0/usersAPI/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	User := &models.Users{
		Username: utils.RandomString(6),
		Email:    utils.RandomEmail(),
		Rol:      utils.RandomString(3),
		Password: utils.RandomString(8),
	}

	err := models.CreateUser(database.InitDB(), User)
	require.NoError(t, err)

}