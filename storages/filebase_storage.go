package storages

import (
	"os"
	"fmt"
	"path/filepath"
)

type FileBaseStorage struct {
	fullPath string
}

func (f *FileBaseStorage) Name() string {
	return "filebase_storage"
}

func NewFileBaseStorage(rootDir, dirName, filename string) (*FileBaseStorage, error) {
	store := FileBaseStorage{}

	fileInfo, err := os.Stat(rootDir)
	if err != nil {
		return nil, &ErrStorageSetup{msg: fmt.Sprintf("FileBaseStorage Setup Fail (rootDir: %s): os.Stat fail: %s", rootDir, err.Error())}
	}

	if !fileInfo.IsDir() {
		return nil, &ErrStorageSetup{msg: fmt.Sprintf("FileBaseStorage Setup Fail (rootDir: %s): rootDir should be a directory", rootDir)}
	}

	fullPath := filepath.Join(rootDir, dirName)

	fileInfo, err = os.Stat(fullPath)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok {
			return nil, &ErrStorageSetup{msg: fmt.Sprintf("FileBaseStorage Setup Fail (rootDir: %s, dirName: %s, fullPath: %s): os.Stat fail: %s", rootDir, dirName, fullPath, err.Error())}
		}
		err = os.Mkdir(fullPath, 0o755)
		if err != nil {
			if pathErr, ok := err.(*os.PathError); ok {
				return nil, &ErrStorageSetup{msg: fmt.Sprintf("FileBaseStorage Setup Fail (fullPath: %s): os.Mkdir fail: %s", fullPath, pathErr.Err.Error())}
			}
			return nil, &ErrStorageSetup{msg: fmt.Sprintf("FileBaseStorage Setup Fail (fullPath: %s): os.Mkdir fail: %s", fullPath, err.Error())}
		}
	}

	fullFilePath := filepath.Join(fullPath, filename)
	fileInfo, err = os.Stat(fullFilePath)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); !ok {
			return nil, &ErrStorageSetup{msg: fmt.Sprintf("FileBaseStorage Setup Fail (full file path: %s): os.Stat fail: %s", fullFilePath, pathErr.Err.Error())}
		}
		_, err = os.Create(filepath.Join(fullPath, filename))
		if err != nil {
			if pathErr, ok := err.(*os.PathError); ok {
				return nil, &ErrStorageSetup{msg: fmt.Sprintf("FileBaseStorage Setup Fail (fullPath: %s, filename: %s): os.Create fail: %s", fullPath, filename, pathErr.Err.Error())}
			}
			return nil, &ErrStorageSetup{msg: fmt.Sprintf("FileBaseStorage Setup Fail (fullPath: %s, filename: %s): os.Create fail: %s", fullPath, filename, err.Error())}
		}
	}

	store.fullPath = fullFilePath 

	return &store, nil
}

func (s *FileBaseStorage) Notify(text string) error {
	file, err := os.OpenFile(s.fullPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			return  &ErrStorageNotify{msg: fmt.Sprintf("FileBaseStorage Notify Fail: os.OpenFile is fail: %s", pathErr.Err.Error())}
		}
		return  &ErrStorageNotify{msg: fmt.Sprintf("FileBaseStorage Notify Fail: os.OpenFile is fail: %s", err.Error())}
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return  &ErrStorageNotify{msg: fmt.Sprintf("FileBaseStorage Notify Fail: file.WriteString is fail: %s", err.Error())}
	}

	return nil
}
