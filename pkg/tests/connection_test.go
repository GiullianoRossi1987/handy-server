package connection_test

import (
	pkg "pkg"
	"testing"
	utils "utils"

	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	_, err := pkg.GeneratePool(utils.GenerateDatabaseConfig())
	assert.Nil(t, err, "Shouldn't have returned any errors from creating the pool")
}
