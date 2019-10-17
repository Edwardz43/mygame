package crypt_test

import (
	"log"
	"testing"

	crypt "github.com/Edwardz43/mygame/gameserver/lib/crypt"
	"github.com/stretchr/testify/assert"
)

func TestEncryptAndDecrypt(t *testing.T) {

	testString := []byte("Hello World")

	ciphertext := crypt.Encrypt(testString, "password")
	plaintext := crypt.Decrypt(ciphertext, "password")

	assert.Equal(t, testString, plaintext)
}

func TestGetToken(t *testing.T) {
	tt := crypt.GetToken(`test`)
	log.Println(tt)
	assert.NotNil(t, tt)
}
