package validator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const source = "./tests"

type test struct {
	Flag     string      `json:"flag"`
	Target   *string     `json:"target"`
	Expected interface{} `json:"expected"`
}

// Model defines an interface for FeatureConfig, Segment and Target types
type Model interface {
	FeatureConfig | Segment | Target
	GetIdentifier() string
}

// Entities is a slice of Models
type Entities[T Model] []T

func (e Entities[T]) Find(identifier string) int {
	for i, ent := range e {
		if ent.GetIdentifier() == identifier {
			return i
		}
	}
	return -1
}

type testFile struct {
	Flags    Entities[FeatureConfig] `json:"flags"`
	Segments Entities[Segment]       `json:"segments"`
	Targets  Entities[Target]        `json:"targets"`
	Tests    []test                  `json:"tests"`
}

func (t testFile) validate() error {
	for _, testCase := range t.Tests {
		if testCase.Target != nil && t.Targets.Find(*testCase.Target) == -1 {
			return fmt.Errorf("target %s not found", *testCase.Target)
		}

		if t.Flags.Find(testCase.Flag) == -1 {
			return fmt.Errorf("flag %s not found", testCase.Flag)
		}

		if err := t.validateFlags(); err != nil {
			return err
		}
	}
	return nil
}

func (t testFile) validateFlags() error {
	for _, config := range t.Flags {
		if config.Rules == nil {
			continue
		}
		for _, rule := range *config.Rules {
			for _, clause := range rule.Clauses {
				if clause.Values == nil || len(clause.Values) == 0 {
					return fmt.Errorf("clause value cannot be empty on flag %s", config.Feature)
				}
				if clause.Op == "segmentMatch" {
					if t.Segments.Find(clause.Values[0]) == -1 {
						return fmt.Errorf("flag segment %s not found in segments", clause.Values[0])
					}
				}
			}
		}

		if config.Prerequisites != nil && len(*config.Prerequisites) > 0 {
			for _, prereq := range *config.Prerequisites {
				if t.Flags.Find(prereq.Feature) == -1 {
					return fmt.Errorf("flag %s with prereq %s not found in flags", config.Feature, prereq.Feature)
				}
			}
		}
	}
	return nil
}

func verify(filename string) error {
	fp := filepath.Clean(filepath.Join(source, filename))
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}

	var test testFile
	err = json.Unmarshal(content, &test)
	if err != nil {
		return err
	}
	return test.validate()
}

func validate() error {
	files, err := ioutil.ReadDir(source)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}
		err = verify(file.Name())
		if err != nil {
			return fmt.Errorf("error loading file %s with err %w", file.Name(), err)
		}

	}
	return nil
}
