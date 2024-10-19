package user

import (
	"backend/types"
	"backend/utils"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceStruct struct {
	dbRepo types.UserDBRepoInterface // repo
}

func NewUserServiceStruct(dbRepo types.UserDBRepoInterface) types.UserServiceInterface {
	return &UserServiceStruct{dbRepo: dbRepo}
}

func (s *UserServiceStruct) LoginUser(payload types.UserPayload) (string, error) {

	user, err := s.dbRepo.GetUserByEmail(payload.Email)
	if err != nil {
		if err == types.ErrNoRecord {
			return "", types.ErrInvalidUserCredentials
		}
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", types.ErrInvalidUserCredentials
	}
	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		return "", fmt.Errorf("Error making JWT. %w", err)
	}
	return jwt, nil
}

func (s *UserServiceStruct) RegisterUser(user types.User) error {

	// check if user already exists
	err := s.dbRepo.CheckEmail(user)
	if err != nil {
		return err
	}

	password := user.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}
	user.Password = string(hashedPassword)

	if err := s.dbRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *UserServiceStruct) RemoveUser(id int) error {
	err := s.dbRepo.DeleteUser(id)
	if err != nil {
		if err == types.ErrNoRecord {
			return types.ErrUserDoesNotExist
		}
		return err
	}
	return nil
}
