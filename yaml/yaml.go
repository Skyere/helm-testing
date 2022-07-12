package yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func GetValue[T any](path string) T {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var values T

	err = yaml.Unmarshal(yamlFile, &values)
	if err != nil {
		panic(err)
	}

	return values
}
