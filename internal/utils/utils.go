package utils

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GeneralCliTokenUUID(userId int) string {
	newUUID := uuid.New()
	// covert UUID to string and remove -
	uuidString := strings.ReplaceAll(newUUID.String(), "-", "")
	return fmt.Sprintf("%dranddnar%s", userId, uuidString)
}
