package loggo

import (
	"sync"
	"github.com/hashen47/loggo/storages"
	"github.com/hashen47/loggo/formatters"
	"github.com/hashen47/loggo/identifiers"
)

type LoggerI interface {
	Debug(vals map[string]string) []error
	Info(vals map[string]string) []error
	Warn(vals map[string]string) []error
	Error(vals map[string]string) []error
	SetLogLevel(level identifiers.LogLevel)
	RegisterStorage(level identifiers.LogLevel, store storages.StorageI)
	RegisterFormatter(formatter *formatters.Formatter)
}

type Logger struct {
	formatter      *formatters.Formatter
	Storages       map[identifiers.LogLevel]map[string]storages.StorageI
	level          identifiers.LogLevel
	storagesMutex  *sync.Mutex
	formatterMutex *sync.Mutex
	levelMutex     *sync.Mutex
}

func NewLogger() LoggerI {
	storages := make(map[identifiers.LogLevel]map[string]storages.StorageI, 0)
	l := Logger{Storages: storages, level: identifiers.Debug, storagesMutex: &sync.Mutex{}, formatterMutex: &sync.Mutex{}, levelMutex: &sync.Mutex{}}
	return &l
}

func (l *Logger) SetLogLevel(level identifiers.LogLevel) {
	l.levelMutex.Lock()
	l.level = level
	l.levelMutex.Unlock()
}

func (l *Logger) RegisterFormatter(formatter *formatters.Formatter) {
	l.formatterMutex.Lock()
	l.formatter = formatter
	l.formatterMutex.Unlock()
}

func (l *Logger) RegisterStorage(level identifiers.LogLevel, store storages.StorageI) {
	l.storagesMutex.Lock()
	if _, ok := l.Storages[level]; !ok {
		l.Storages[level] = make(map[string]storages.StorageI, 0)
	}
	l.Storages[level][store.Name()] = store
	l.storagesMutex.Unlock()
}

func (l *Logger) log(level identifiers.LogLevel, vals map[string]string) []error {
	errs      := make([]error, 0)

	switch l.level {
	case identifiers.Info:
		if level == identifiers.Debug {
			return errs
		}
	case identifiers.Warn:
		if level == identifiers.Debug || level == identifiers.Info {
			return errs
		}
	case identifiers.Error:
		if level != identifiers.Error {
			return errs
		}
	}

	text, err := l.formatter.GetText(&vals)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	if storageMap, ok := l.Storages[level]; ok {
		for _, store := range storageMap {
			err := store.Notify(text + "\n")
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errs
}

func (l *Logger) Debug(vals map[string]string) []error {
	level := identifiers.LogLevelIdentifier{}
	vals[level.Name()] = string(identifiers.Debug)
	return l.log(identifiers.Debug, vals)
}

func (l *Logger) Info(vals map[string]string) []error {
	level := identifiers.LogLevelIdentifier{}
	vals[level.Name()] = string(identifiers.Info)
	return l.log(identifiers.Info, vals)
}

func (l *Logger) Warn(vals map[string]string) []error {
	level := identifiers.LogLevelIdentifier{}
	vals[level.Name()] = string(identifiers.Warn)
	return l.log(identifiers.Warn, vals)
}

func (l *Logger) Error(vals map[string]string) []error {
	level := identifiers.LogLevelIdentifier{}
	vals[level.Name()] = string(identifiers.Error)
	return l.log(identifiers.Error, vals)
}
