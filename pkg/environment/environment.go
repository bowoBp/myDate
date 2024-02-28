package environment

import (
	"log"
	"os"
	"strconv"
)

type (
	Environment interface {
		Get(key string) string
		GetUint(key string, defaultValue uint) uint
		CheckFlag(flag string) bool
	}

	environment struct {
	}
)

func NewEnvironment() Environment {
	return environment{}
}

func (environment) Get(key string) string {
	return os.Getenv(key)
}

func (e environment) CheckFlag(flag string) bool {
	str := os.Getenv(flag)
	status, err := strconv.ParseBool(str)
	if err != nil {
		log.Println(err)
		return false
	}

	return status
}

func (e environment) GetUint(key string, defaultValue uint) uint {
	str := os.Getenv(key)
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Println(err)
		log.Println("featureflag.environment.GetUint:", err)
		value = uint64(defaultValue)
	}

	return uint(value)
}
