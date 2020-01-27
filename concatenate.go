package ditto_yaml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Get(version string) (string, error) {
	result := ""
	path := "type/" + version

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if nil != err {
			return result, err
		}
	}

	_ = filepath.Walk(path + "/abstract", func(path string, info os.FileInfo, err error) error {
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

	return result, err
}
