package users_test

import (
	"errors"
	"pkg"
	op "pkg/users"
	types "types/database/users"

	"context"
	"fmt"
	"math/rand"
	"testing"
	"utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

var (
	pool    *pgxpool.Pool
	conn    *pgxpool.Conn
	usingId int
	login   string
)

func TestMain(t *testing.T) {
	// before all equivalent
	var err error
	utils.MockTestsDatabase(t)
	pool, err = pkg.GeneratePool(utils.GenerateDatabaseConfig())
	if err != nil {
		panic(err.Error())
	}
	conn, err = pool.Acquire(context.Background())
	if err != nil {
		panic(err.Error())
	}
}

func TestAddUser(t *testing.T) {
	id := rand.Int()
	user := types.UsersRecord{
		Login:    fmt.Sprintf("mocker%d", id),
		Password: "testing",
	}
	response, err := op.AddUser(user, conn)
	assert.Nil(t, err, "Shouldn't have returned any errors")
	assert.NotNil(t, &response, "Should have a ID")

	usingId = int(*response)
	login = user.Login
}

func TestGetUserById(t *testing.T) {
	response, err := op.GetUserById(usingId, conn)
	assert.Nil(t, err, "Shouldn't have returned any errors")
	assert.Equal(t, response.Login, login, "Should have returned the actual login identification")
	assert.Equal(t, int(response.Id), usingId, "Should have returned the actual login identification")
}

func TestGetUserByLogin(t *testing.T) {
	response, err := op.GetUserByLogin(login, conn)
	assert.Nil(t, err, "Shouldn't have returned any errors")
	assert.Equal(t, response.Login, login, "Should have returned the actual login identification")
	assert.Equal(t, int(response.Id), usingId, "Should have returned the actual login identification")
}

func TestUpdateUser(t *testing.T) {
	newData := types.UsersRecord{
		Id:       usingId,
		Login:    fmt.Sprintf("modifiedtesting%d", usingId),
		Password: fmt.Sprintf("modifiedtesting%d", usingId),
	}
	err := op.UpdateUserById(newData, conn)
	assert.Nil(t, err, "Should have no errors")
	response, err := op.GetUserById(usingId, conn)
	assert.Nil(t, err, "Should have no errors")
	assert.Equal(t, int(response.Id), usingId, "Should have kept the id")
	assert.Equal(t, response.Login, newData.Login, "Should have returned the new login identification")
}

func TestDeleteUser(t *testing.T) {
	err := op.DeleteUserById(usingId, conn)
	assert.Nil(t, err, "Shouldn't return any errors")
	response, err := op.GetUserById(usingId, conn)
	assert.Nil(t, response, "Shouldn't return any response")
	assert.True(t, errors.Is(pgx.ErrNoRows, err), "Should return no rows error")
}
