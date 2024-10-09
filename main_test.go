package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/json1.json
var json1 []byte

//go:embed testdata/json2.json
var json2 []byte

//go:embed testdata/json_all_type.json
var jsonAllType []byte

func Test_json2yaml(t *testing.T) {
	// got, err := json2yaml(json1)
	// got, err := json2yaml(json2)
	got, err := json2yaml(jsonAllType)
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	fmt.Println(string(got))
}
