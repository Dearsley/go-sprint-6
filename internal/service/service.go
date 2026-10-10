package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("convert fail: input is empty")
	}

	if strings.Trim(input, " .-") == "" {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}
