package main

import (
	"encoding/json"
	"log"
	"os"
)

type Storage interface {
	Save(tasks []*Task) error
	Read() ([]*Task, error)
}

type JsonStorage struct {
	filePath string
}

func NewStorage(filePath string) Storage {
	return &JsonStorage{
		filePath: filePath,
	}
}

func (s *JsonStorage) Save(tasks []*Task) error {
	file, err := os.Create(s.filePath)
	if err != nil {
		log.Println("create: ", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(tasks); err != nil {
		log.Println("encode:", err)
		return err
	} 

	return nil
}

func (s *JsonStorage) Read() ([]*Task, error) {
	if stat, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return make([]*Task, 0), nil
	} else if stat.Size() == 0 {
		return make([]*Task, 0), nil
	} else if err != nil {
		return nil, err
	}

	file, err := os.Open(s.filePath)
	if err != nil {
		log.Println("open: ", err)
		return nil, err
	}
	defer file.Close()

	var tasks []*Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		log.Println("decode: ", err)
		return nil, err
	}

	return tasks, nil
}
