package loader

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Loader load toml file and merge them
type Loader struct {
}

// NewLoader creates a new loader with an initialized map
func NewLoader() *Loader {
	return &Loader{}
}

// Load loads a toml data file and merege inherited data
func (g *Loader) Load(dir string, name string) (*Resume, error) {
	return g.load(dir, name, make([]string, 10))
}

// Load loads toml data file and merge into inherited data. It tracks loaded data
// to detect circular inheritance
func (g *Loader) load(dir string, name string, seen []string) (*Resume, error) {
	if contains(name, seen) {
		return nil, fmt.Errorf("Circular inheritance detected with file %s", name)
	}
	var resume *Resume

	//load the toml file
	filepath := filepath.Join(dir, name)
	tomlData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	// out the file content
	if _, err := toml.Decode(string(tomlData), &resume); err != nil {
		return nil, err
	}
	return resume, nil
}

func contains(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
