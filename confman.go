package confman

import (
	"os"
	"reflect"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

func GetThisFolder() string{

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	return dir
}

func StructToMap(aStruct interface{}) (map[string]interface{}) {

	newMap := make(map[string]interface{})

	thisStruct := reflect.ValueOf(aStruct)

	structType := thisStruct.Type()

	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()}

	for i := 0; i < thisStruct.NumField(); i++ {
		field := thisStruct.Field(i)
		//fmt.Printf("%d: %s %s = %v\n", i, structType.Field(i).Name, field.Type(), field.Interface())
		newMap[structType.Field(i).Name] = field.Interface()
	}
	
	return newMap
}

func LoadRaw(path string) (string, error) {

	buf, err := ioutil.ReadFile(path)

	return string(buf), err
}

func LoadJson(path string) (map[string]interface{}, error){

	buf, err := ioutil.ReadFile(path)

	if err != nil {return nil, err}

	newMap := make(map[string]interface{})

	json.Unmarshal(buf, &newMap)
	return newMap, err
}

func SaveJson(path string, alteredConfig map[string]interface{}) error {

	buf, err := json.MarshalIndent(alteredConfig, "", "  ")

	if err != nil {return err}

	ioutil.WriteFile(path, buf, 0644)

	return nil
}