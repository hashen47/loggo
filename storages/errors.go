package storages

import (
	"fmt"
)

type ErrStorageSetup struct {
	msg string
}

func (e *ErrStorageSetup) Error() string {
	return fmt.Sprintf("Err: STORAGE_SETUP_ERROR: %s", e.msg)
}

type ErrStorageNotify struct {
	msg string
}

func (e *ErrStorageNotify) Error() string {
	return fmt.Sprintf("Err: STORAGE_NOTIFY_ERROR: %s", e.msg)
}

