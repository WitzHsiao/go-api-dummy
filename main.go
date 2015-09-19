package main

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

const api_yml = "api.yml"

func main() {

	data, err := ioutil.ReadFile(api_yml)
	if err != nil {
		panic(err)
	}
	m := make(map[interface{}]interface{})
	yaml.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	spew.Dump(m)
}
