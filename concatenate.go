package ditto_yaml

import (
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func Get(version string) (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("caller not found")
	}

	result := ""
	path := filepath.Dir(filename) + "/type/" + version

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if nil != err {
			return result, err
		}
	}

	_ = filepath.Walk(path+"/abstract", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		b, err := ioutil.ReadFile(path)
		if nil != err {
			return err
		}

		result += fmt.Sprintf("\n%s", string(b))
		return err
	})

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if info.Name() == "abstract" {
				return filepath.SkipDir
			}
			return nil
		}

		b, err := ioutil.ReadFile(path)
		if nil != err {
			return err
		}

		result += fmt.Sprintf("\n%s", string(b))
		return err
	})

	if nil != err {
		return result, err
	}

	m := make(map[string]interface{})

	err = yaml.Unmarshal([]byte(result), &m)
	if nil != err {
		return result, err
	}

	byt, err := yaml.Marshal(m)
	if nil != err {
		return result, err
	}

	result = string(byt)

	return result, err
}
