package fileSystemRepository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"jackthomson/go-api/models"
)

func GetFileContents(path string) ([]byte, error) {
	result, err := ioutil.ReadFile(path)

	if err != nil { return nil, errors.New("error reading file") }

	var users []models.User

	err = json.Unmarshal(result, &users)

	if err != nil { return nil, errors.New("error unmarshalling") }

	return result, nil
}

func WriteFileContents(path string, data []byte) error {
	err := ioutil.WriteFile("data/users.json", data, 0644)

	if err != nil { return errors.New("error writing file") }

	return nil
}