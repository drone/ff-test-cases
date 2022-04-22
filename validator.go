package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Model interface {
	FeatureConfig | Segment | Target
}

func validate() error {
	_, err := loadFiles[Segment]("./segments")
	if err != nil {
		return err
	}
	_, err = loadFiles[Target]("./targets")
	if err != nil {
		return err
	}
	_, err = loadFiles[FeatureConfig]("./flags")
	if err != nil {
		return err
	}
	return nil
}

func loadFiles[T Model](dir string) (map[string]T, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	store := make(map[string]T)
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}
		value, err := verifyJSON[T](dir, file.Name())
		if err != nil {
			return nil, err
		}

		_, ok := store[file.Name()]
		if ok {
			return nil, fmt.Errorf("flag configuration already exist %v", file.Name())
		}
		store[file.Name()] = value
	}
	return store, nil
}

func verifyJSON[T Model](dir string, filename string) (T, error) {
	var result T
	fp := filepath.Clean(filepath.Join(dir, filename))
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(content, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
