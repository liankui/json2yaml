package main

import (
	"log"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
)

func main() {
}

func json2yaml(data []byte) ([]byte, error) {
	var tmp map[string]any

	err := jsoniter.Unmarshal(data, &tmp)
	if err != nil {
		log.Println("json unmarshal error:", err)
		return nil, err
	}

	out, err := yaml.Marshal(&tmp)
	if err != nil {
		log.Println("yaml marshal error:", err)
		return nil, err
	}

	return out, nil
}
