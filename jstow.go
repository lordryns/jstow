package jstow

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// the base struct, not to be used directly
// outside this package
type JstowBase[T any] struct {
	Path string
	Body map[string]T
}

func Jstow[T any](path string) *JstowBase[T] {
	// the underscore is for the error
	// no error handling yet, add later
	var payload, err = loadJson[T](path)
	if err != nil {
		file, err := os.Create(path)
		if err != nil {
			fmt.Println("Unable to create file!")
		}

		file.WriteString("{}")

		defer file.Close()
	}

	return &JstowBase[T]{Path: path, Body: payload}
}

// this function is simply for elegantly retuning every value there
func (j *JstowBase[T]) All() map[string]T {
	var body, err = loadJson[T](j.Path)
	var anError = make(map[string]T)
	if err != nil {
		return anError
	}
	return body
}

func (j *JstowBase[T]) Insert(data T) error {
	var err = insertJson(j.Path, data)
	if err != nil {
		return err
	}

	return nil
}

func (j *JstowBase[T]) Update(fieldName string, targetValue string, newData T) error {
	var err = updateData(j.Path, fieldName, targetValue, newData)
	if err != nil {
		return err
	}
	return nil
}

func (j *JstowBase[T]) Search(fieldName string, targetValue string) ([]T, error) {
	var response, _, err = searchData[T](j.Path, fieldName, targetValue)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (j *JstowBase[T]) Delete(fieldName string, targetValue string) error {
	var err = deleteRow[T](j.Path, fieldName, targetValue)
	if err != nil {
		return err
	}

	return nil
}

// ++++++++++++++ local functions +++++++++++++++++
//
// // this is a local function responsible for getting
// the data from the json file, parsing it into a map
func loadJson[T any](path string) (map[string]T, error) {
	var result = make(map[string]T)

	var bytes, err = os.ReadFile(path)
	// go's error handling is good idc
	if err != nil {
		return result, err
	}

	err2 := json.Unmarshal(bytes, &result) // convert to map here
	// error watch
	if err2 != nil {
		return result, err2
	}
	return result, nil
}

// insert new row
func insertJson[T any](path string, data T) error {
	currentData, err := loadJson[T](path)
	if err != nil {
		return errors.New("Unable to access json file!")
	}

	var keys []string
	for key, _ := range currentData {
		keys = append(keys, key)
	}

	var currentKey int
	var err2 error
	if len(keys) > 0 {
		currentKey, err2 = strconv.Atoi(keys[len(keys)-1])
		if err2 != nil {
			return errors.New("Invalid format!")
		}
	} else {
		currentKey = 0
	}

	currentData[strconv.Itoa(currentKey+1)] = data
	var err3 = writeToJson(path, currentData)

	if err3 != nil {
		return err3
	}

	return nil
}

// this rewrites the entire file
func writeToJson[T any](path string, data T) error {
	var stringData, err = json.Marshal(data)

	if err != nil {
		return errors.New("Unable to convert to json")
	}
	var byteData = []byte(stringData)
	var err2 = os.WriteFile(path, byteData, 0644)

	if err2 != nil {
		return errors.New("Unable to write to json")
	}

	return nil
}

func updateData[T any](path string, fieldName string, targetValue string, newData T) error {
	_, keys, err := searchData[T](path, fieldName, targetValue)
	if err != nil {
		return err
	}

	var dbData, err2 = loadJson[T](path)
	if err2 != nil {
		return err2
	}

	for _, key := range keys {
		dbData[key] = newData
	}

	var err3 = writeToJson(path, dbData)
	if err3 != nil {
		return err3
	}

	return nil
}

func searchData[T any](path string, fieldName string, targetValue string) ([]T, []string, error) {
	var response []T
	var keys []string
	var jsonData, err = loadJson[T](path)

	if err != nil {
		return response, keys, err
	}

	for key, row := range jsonData {
		var refl = reflect.ValueOf(row)
		if refl.Kind() == reflect.Struct {
			var field = refl.FieldByName(fieldName)
			if field.IsValid() {
				var value = field.Interface()
				if fmt.Sprint(value) == fmt.Sprint(targetValue) {
					response = append(response, row)
					keys = append(keys, key)
				}
			}
		}
	}

	return response, keys, nil

}

func deleteRow[T any](path string, fieldName string, targetValue string) error {
	var _, keys, err = searchData[T](path, fieldName, targetValue)
	var table, err2 = loadJson[T](path)
	if err != nil {
		return err
	}

	if err2 != nil {
		return err2
	}

	for _, key := range keys {
		delete(table, key)
	}

	writeToJson(path, table)

	return nil
}
