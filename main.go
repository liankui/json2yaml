package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/json2yaml", json2yamlHandler)
	http.HandleFunc("/yaml2json", yaml2jsonHandler)
	fmt.Println("starting server on :21111")
	if err := http.ListenAndServe(":21111", nil); err != nil {
		log.Fatal(err)
	}
}

func json2yamlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read request body, error:%v", err)
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	bytes, err := json2yaml(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/x-yaml")
	_, _ = w.Write(bytes)
}

func yaml2jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read request body, error:%v", err)
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	bytes, err := yaml2json(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(bytes)
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
