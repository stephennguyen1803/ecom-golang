package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashByBytes := hash.Sum(nil)
	return hex.EncodeToString(hashByBytes)
}

// GeneralSalt generate random salt
func GeneralSalt(length int) (string, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

// HashPassword generate hash password
func HashPassword(password, salt string) string {
	// concatenate password and salt
	saltedPassword := password + salt
	// hash the combined password
	hashPass := sha256.Sum256([]byte(saltedPassword))
	return hex.EncodeToString(hashPass[:])
}

// MatchingPassword compare password
func MatchingPassword(password, salt, storeHash string) bool {
	hashPass := HashPassword(password, salt)
	return hashPass == storeHash
}
