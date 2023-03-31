package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewUuid() string {
	return uuid.New().String()
}

func NewPureUuid() string {
	return strings.ReplaceAll(NewUuid(), "-", "")
}
