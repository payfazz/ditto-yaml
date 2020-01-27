package ditto_yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Get(version string) (string, error) {
	result := ""
	path := "type/" + version

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if nil != err {
			return result, err
		}
	}

	abstracts := make(map[string]bool)
	_ = filepath.Walk(path+"/abstract", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		name := strings.Split(info.Name(), ".")[0]
		abstracts[name] = true

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

	m2 := make(map[string]interface{})
	for k, v := range m {
		if _, ok := abstracts[k]; ok {
			continue
		}
		m2[k] = v
	}

	byt, err := yaml.Marshal(m2)
	if nil != err {
		return result, err
	}

	result = string(byt)

	return result, err
}
