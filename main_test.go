package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/json/json1.json
var json1 []byte

//go:embed testdata/json/json2.json
var json2 []byte

//go:embed testdata/json/json_all_type.json
var jsonAllType []byte

func Test_json2yaml(t *testing.T) {
	test := func(data []byte) {
		got, err := json2yaml(data)
		assert.NoError(t, err)
		assert.NotEmpty(t, got)
		fmt.Println(string(got))
	}

	test(json1)
	test(json2)
	test(jsonAllType)
}

//go:embed testdata/yaml/yaml1.yaml
var yaml1 []byte

//go:embed testdata/yaml/yaml2.yaml
var yaml2 []byte

//go:embed testdata/yaml/yaml_all_type.yaml
var yamlAllType []byte

func Test_yaml2json(t *testing.T) {
	test := func(data []byte) {
		got, err := yaml2json(data)
		assert.NoError(t, err)
		assert.NotEmpty(t, got)
		fmt.Println(string(got))
	}

	test(yaml1)
	test(yaml2)
	test(yamlAllType)
}
