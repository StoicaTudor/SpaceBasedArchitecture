package data_contracts

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type UserCreateDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (dto *UserCreateDTO) GetAction() Action {
	return CREATE
}

func (dto *UserCreateDTO) GetCommandType() CommandType {
	return CREATE_USER_DTO
}

type UserUpdateDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (dto *UserUpdateDTO) GetAction() Action {
	return UPDATE
}

func (dto *UserUpdateDTO) GetCommandType() CommandType {
	return UPDATE_USER_DTO
}

type UserDeleteDTO struct {
	ID string `json:"id"`
}

func (dto *UserDeleteDTO) GetAction() Action {
	return DELETE
}

func (dto *UserDeleteDTO) GetCommandType() CommandType {
	return DELETE_USER_DTO
}

func GetRandomCommand() Command {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch random.Intn(3) {
	case 0:
		return &UserCreateDTO{
			ID:      uuid.New().String(),
			Name:    fmt.Sprintf("User%d", random.Intn(100)),
			Balance: rand.Float64() * 1000,
		}
	case 1:
		return &UserUpdateDTO{
			ID:      uuid.New().String(),
			Name:    fmt.Sprintf("User%d", random.Intn(100)),
			Balance: random.Float64() * 1000,
		}
	case 2:
		return &UserDeleteDTO{
			ID: uuid.New().String(),
		}
	default:
		return nil
	}
}
