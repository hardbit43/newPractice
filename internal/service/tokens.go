package service

import (
	"crypto/rsa"
	"log"
	"my_api/internal/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type IDTokenCustomClaims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

type RefreshTokenCustomClaims struct {
	UID uuid.UUID
	jwt.StandardClaims
}

type RefreshToken struct {
	SS        string
	ID        string
	ExpiresIn time.Duration
}

func generateIDToken(u *model.User, key *rsa.PrivateKey) (string, error) {
	unixTime := time.Now().Unix()
	tokenExp := unixTime + 60*15 // 15 min from curent time

	claims := IDTokenCustomClaims{
		User: u,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  unixTime,
			ExpiresAt: tokenExp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(key)

	if err != nil {
		log.Println("Failed to sign id token string")
	}
	return ss, err
}

func generateRefreshToken(uid uuid.UUID, key string) (*RefreshToken, error) {
	curentTime := time.Now()
	tokenExp := curentTime.AddDate(0, 0, 3) // 3 days
	tokenID, err := uuid.NewRandom()

	if err != nil {
		log.Println("Failed to generate refresh token ID")
	}

	claims := RefreshTokenCustomClaims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  curentTime.Unix(),
			ExpiresAt: tokenExp.Unix(),
			Id:        tokenID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))

	if err != nil {
		log.Println("Failed to sign refresh token string")
	}
	return &RefreshToken{
		SS:        ss,
		ID:        tokenID.String(),
		ExpiresIn: tokenExp.Sub(curentTime),
	}, nil
}
