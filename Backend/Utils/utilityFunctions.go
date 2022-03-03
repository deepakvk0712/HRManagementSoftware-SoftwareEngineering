package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func HashPassword(password string) (string, int) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", 0
	}

	return string(hashedPassword), 1
}

func GenerateAccessToken(email string, role byte) (string, error) {
	claims := &models.JWTClaim{
		Email: email,
		StandardClaim: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(LoginJWT.LoginTimeout)).Unix(),
			Issuer:    LoginJWT.Issuer,
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(LoginJWT.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, err
}

func ValidateToken(signedToken string) (*models.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(signedToken, &models.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(LoginJWT.SecretKey), nil
	})

	if err != nil {
		return &models.JWTClaim{}, err
	}

	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		return &models.JWTClaim{}, errors.New("failed to parse")
	}

	if claims.StandardClaim.ExpiresAt < time.Now().Local().Unix() {
		return &models.JWTClaim{}, errors.New("expired")
	}

	return claims, nil
}

func GetUsername(email string) string {
	emailSlice := strings.Split(email, "@")

	return emailSlice[0]
}
