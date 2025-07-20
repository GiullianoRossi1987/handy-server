package utils

import (
	"os"
	types "types/config"

	crypt "golang.org/x/crypto/bcrypt"
)

func GenerateDatabaseConfig() types.PsConfig {
	return types.PsConfig{
		Host:     os.Getenv("HOSTNAME"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Db:       os.Getenv("DATABASE"),
	}
}

// Why GO doesnt ship without a '??' coalesce function? I gotta do this ¬~¬
func Coalesce[T any](value *T, default_val T) T {
	if value == nil {
		return default_val
	}
	return *value
}

// fucking hate the fact there there isn't FP in this language, why google why
func MapCar[T any, U any](ia []T, of func(item T) U) []U {
	mapped := make([]U, len(ia))
	for k, i := range ia {
		mapped[k] = of(i)
	}
	return mapped
}

// TODO implement this on user password
func EncryptPassword(password string) (string, error) {
	bpass := []byte(password)
	passwd, err := crypt.GenerateFromPassword(bpass, crypt.DefaultCost)
	if err != nil {
		return "nil", err
	}
	return string(passwd), nil
}

func ValidatePassword(incomming string, hashed string) bool {
	err := crypt.CompareHashAndPassword([]byte(hashed), []byte(incomming))
	return err == nil
}
