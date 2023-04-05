package service

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func hashPassword(password string) (string, error) {
	salt := make([]byte, 32)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}

	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))
	return hashedPW, nil

}

func comparePasswords(stoerdPassword, string, suppliedPassword string) (bool, error) {
	pwsalt := strings.Split(stoerdPassword, ".")

	salt, err := hex.DecodeString(pwsalt[1])

	if err != nil {
		return false, fmt.Errorf("Unable to verify password")
	}

	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)

	return hex.EncodeToString(shash) == pwsalt[0], nil
}
