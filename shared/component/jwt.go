package sharedcomponent

import (
	"fmt"
	"time"
	
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type JWTComp struct {
	secretKey string
	duration  time.Duration
}

func NewJWTComp(secretKey string, duration time.Duration) *JWTComp {
	return &JWTComp{
		secretKey: secretKey,
		duration:  duration,
	}
}

func (j *JWTComp) IssueToken(userID string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	now := time.Now()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(now.Add(j.duration)),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        fmt.Sprintf("%d", now.UnixNano()),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := claims.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", errors.WithStack(err)
	}
	
	return tokenString, nil
}

func (j *JWTComp) ExpiresIn() time.Duration {
	return j.duration
}

func (j *JWTComp) Validate(tokenString string) (*jwt.RegisteredClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token's signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	
	return nil, errors.New("invalid token")
}
