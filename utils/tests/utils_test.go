package utils_test

import (
	"math/rand"
	"testing"
	utils "utils"

	"github.com/stretchr/testify/assert"
)

func TestCoalesce(t *testing.T) {
	var value *int
	const default_value = 50
	v := (rand.Int())
	value = &v
	coalesced := utils.Coalesce(value, default_value)
	assert.Equal(t, coalesced, *value, "Shouldn't have changed to the default value")
	value = nil
	coalesced = utils.Coalesce(value, default_value)
	assert.Equal(t, coalesced, default_value, "Should have changed to the default value")
}

func TestMap(t *testing.T) {
	values := []int{1, 2, 3, 4}
	mapped := utils.MapCar(values, func(v int) int {
		return v * 2
	})
	assert.Equal(t, mapped, []int{2, 4, 6, 8}, "Should have mapped all the itens to their new values")
}

func TestEncryptPassword(t *testing.T) {
	const password = "testing"
	crypted, err := utils.EncryptPassword(password)
	if err != nil {
		panic(err.Error())
	}
	assert.NotEqual(t, password, crypted, "The encrypted password should be diferent from the plain password")
	validation := utils.ValidatePassword(password, crypted)
	assert.True(t, validation, "Should validate the password correctly")
	validation = utils.ValidatePassword("password", crypted)
	assert.False(t, validation, "Should validate the password correctly")
}
