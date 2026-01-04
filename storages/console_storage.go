package storages

import (
	"os"
	"fmt"
)

type ConsoleStorage struct {}

func NewConsoleStorage() *ConsoleStorage {
	return &ConsoleStorage{} 
}

func (s *ConsoleStorage) Name() string {
	return "console_storage"
}

func (s *ConsoleStorage) Notify(text string) error {
	_, err := fmt.Fprint(os.Stdout, text)
	if err != nil {
		return &ErrStorageNotify{msg: fmt.Sprintf("ConsoleStorage Notify Fail: fmt.Fprintf fail: %s", err.Error())}
	}
	return nil
}
