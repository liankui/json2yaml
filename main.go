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

func yaml2json(data []byte) ([]byte, error) {
	var tmp map[string]any

	err := yaml.Unmarshal(data, &tmp)
	if err != nil {
		log.Println("yaml unmarshal error:", err)
		return nil, err
	}

	out, err := jsoniter.MarshalIndent(&tmp, "", "  ")
	if err != nil {
		log.Println("json marshal error:", err)
		return nil, err
	}

	return out, nil
}
