package shared

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"unicode"
	
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomStr generates a random string with the specified length
func RandomStr(length int) (string, error) {
	if length <= 0 {
		return "", errors.WithStack(fmt.Errorf("length must be greater than 0"))
	}
	
	result := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))
	
	for i := 0; i < length; i++ {
		// Generate a random number from 0 to len(charset)-1
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", errors.WithStack(err)
		}
		result[i] = charset[randomIndex.Int64()]
	}
	
	return string(result), nil
}

// CheckValidEmailFormat checks if the provided email is in a valid format.
// It matches against a regex pattern that requires:
// - One or more letters, numbers, dots, underscores, percent signs, plus signs, or hyphens before the @ symbol
// - One or more letters, numbers, or hyphens after the @ symbol
// - A dot followed by at least 2 letters at the end
func CheckValidEmailFormat(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// CheckValidPasswordFormat checks if the provided password meets security requirements.
// Password must:
// - Be at least 8 characters long
// - Contain at least one uppercase letter
// - Contain at least one lowercase letter
// - Contain at least one digit
// - Contain at least one special character (non-alphanumeric)
func CheckValidPasswordFormat(password string) bool {
	if len(password) < 8 {
		return false
	}
	
	var hasLower, hasUpper, hasDigit bool
	
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case !strings.ContainsRune("@$!%*?&", char) && !unicode.IsLetter(char) && !unicode.IsDigit(char):
			return false
		}
	}
	
	return hasLower && hasUpper && hasDigit
}

// CheckPassword verifies if a password matches a bcrypt-generated hash
func CheckPassword(hashedPassword, password, salt string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))
}
