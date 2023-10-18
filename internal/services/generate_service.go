package services

import (
	"errors"
	"math/rand"
	"time"
)

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numbers = "0123456789"
var symbols = "!@#$%^&*()_+{}[]|:;<>,.?/~"

type GenerateCommand struct {
	Length  int8
	Letters bool
	Symbols bool
	Numbers bool
}

type GenerateService interface {
	GeneratePassword(command *GenerateCommand) (string, error)
}

type service struct {
}

func NewGenerateService() GenerateService {
	return &service{}
}

func (*service) GeneratePassword(command *GenerateCommand) (string, error) {
	if command.Length < 4 || command.Length > 40 {
		return "", errors.New("length must be between 4 and 40")
	}

	if !command.Letters && !command.Numbers && !command.Symbols {
		return "", errors.New("you must include at least one type of character in the password")
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	permitCharacters := ""
	if command.Letters {
		permitCharacters += letters
	}
	if command.Numbers {
		permitCharacters += numbers
	}
	if command.Symbols {
		permitCharacters += symbols
	}

	password := make([]byte, command.Length)
	for i := range password {
		r := rand.Float64()
		if r < 0.4 && command.Letters {
			password[i] = permitCharacters[rand.Intn(len(permitCharacters))]
		} else if r < 0.4+0.35 && command.Numbers {
			password[i] = permitCharacters[rand.Intn(len(permitCharacters))]
		} else if command.Symbols {
			password[i] = permitCharacters[rand.Intn(len(permitCharacters))]
		} else {
			i--
		}
	}

	for i := len(password) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		password[i], password[j] = password[j], password[i]
	}

	return string(password), nil
}
