package config_test

import (
	"github.com/Edwardz43/mygame/app/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGetDBConfigShouldReturnDBConfig(t *testing.T) {
// 	expectedResult := "root:root@tcp(localhost:13306)/MyGame"

// 	actualResult := config.GetDBConfig()

// 	assert.Equal(t, expectedResult, actualResult)
// }

// func TestGetDBConfigV2ShouldReturnDBConfig(t *testing.T) {
// 	expectedResult := "host=postgres port=5432 user=admin dbname=postgres password=test sslmode=disable"

// 	actualResult := config.GetDBConfigV2()

// 	assert.Equal(t, expectedResult, actualResult)
// }

// func TestGetLogstashShouldReturnConnectionString(t *testing.T) {
// 	expectedResult := "192.168.1.101:5000"

// 	actualResult := config.GetLogstashConfig()

// 	assert.Equal(t, expectedResult, actualResult)
// }

func TestGetGetELKShouldReturnIsEnable(t *testing.T) {
	expectedResult := false

	actualResult := config.GetELKConfig()

	assert.Equal(t, expectedResult, actualResult)
}
