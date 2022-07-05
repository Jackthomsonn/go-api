package userService

import (
	"encoding/json"
	"errors"
	"jackthomson/go-api/models"
	fileSystemRepository "jackthomson/go-api/repositories/fileSystem"
)

func GetUsers() ([]models.User, error) {
	result, err := fileSystemRepository.GetFileContents("data/users.json")

	if err != nil { return nil, errors.New(err.Error()) }

	var users []models.User

	err = json.Unmarshal(result, &users)

	if err != nil { return nil, errors.New(err.Error()) }

	return users, nil
}

func GetUser(name string) (models.User, error) {
	result, err := fileSystemRepository.GetFileContents("data/users.json")

	if err != nil { return models.User{}, errors.New(err.Error()) }

	var users []models.User

	err = json.Unmarshal(result, &users)

	if err != nil { return models.User{}, errors.New(err.Error()) }

	user := models.User{}

	for _, u := range users {
		if u.Name == name {
			user = u
		}
	}

	if user == (models.User{}) {
		return user, errors.New("user not found")
	}

	return user, nil
}

func CreateUser(user []byte) (error) {
	result, err := fileSystemRepository.GetFileContents("data/users.json")

	if err != nil { return errors.New(err.Error()) }

	if err != nil { return errors.New(err.Error()) }

	users := []models.User{}
	json.Unmarshal(result, &users)

	newUser := models.User{}
	json.Unmarshal(user, &newUser)

	for _, u := range users {
		if u.Name == newUser.Name {
			return errors.New("user already exists")
		}
	}

	users = append(users, newUser)

	dataBytes, err := json.Marshal(users)

	if err != nil { return errors.New(err.Error()) }

	fileSystemRepository.WriteFileContents("data/users.json", dataBytes)

	return nil

}