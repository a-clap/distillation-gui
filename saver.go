package main

import (
	"os"
	"path"
)

type saver struct {
	path string
}

func (l *saver) Save(key string, data []byte) error {
	filePath := path.Join(l.path, key)
	return os.WriteFile(filePath, data, 0666)
}

func (l *saver) Load(key string) (data []byte, err error) {
	filePath := path.Join(l.path, key)
	f, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return f, nil
}
