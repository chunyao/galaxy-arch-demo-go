package utils

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

// NewUUID 获取UUID
func NewUUID() string {
	return strings.Replace(uuid.NewV1().String(), "-", "", -1)
}
