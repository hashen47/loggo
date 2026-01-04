package loggo

import (
	"fmt"
	"github.com/hashen47/loggo/storages"
	"github.com/hashen47/loggo/formatters"
	"github.com/hashen47/loggo/identifiers"
)

func NewDefaultLogger(rootDir, dirName string) (LoggerI, error) {
	l := NewLogger() 

	dateTimeIdentifier := identifiers.DateTimeIdentifier{}
	logLevelIdentifier := identifiers.LogLevelIdentifier{}
	messageIdentifier  := identifiers.MessageIdentifier{}

	formatter := formatters.NewFormatter(
		fmt.Sprintf("[%%%s%%] [%%%s%%]: %%%s%%", dateTimeIdentifier.Name(), logLevelIdentifier.Name(), messageIdentifier.Name()),
		dateTimeIdentifier,
		logLevelIdentifier,
		messageIdentifier,
	)
	l.RegisterFormatter(formatter)

	consoleStore := storages.NewConsoleStorage()

	debugFilebaseStore, err := storages.NewFileBaseStorage(rootDir, dirName, "debug.log")
	if err != nil {
		return nil, err
	}

	infoFilebaseStore, err := storages.NewFileBaseStorage(rootDir, dirName, "info.log")
	if err != nil {
		return nil, err
	}

	warnFilebaseStore, err := storages.NewFileBaseStorage(rootDir, dirName, "warn.log")
	if err != nil {
		return nil, err
	}

	errorFilebaseStore, err := storages.NewFileBaseStorage(rootDir, dirName, "error.log")
	if err != nil {
		return nil, err
	}

	l.RegisterStorage(identifiers.Debug, consoleStore)
	l.RegisterStorage(identifiers.Info , consoleStore)
	l.RegisterStorage(identifiers.Warn , consoleStore)
	l.RegisterStorage(identifiers.Error, consoleStore)

	l.RegisterStorage(identifiers.Debug, debugFilebaseStore)
	l.RegisterStorage(identifiers.Info , infoFilebaseStore)
	l.RegisterStorage(identifiers.Warn , warnFilebaseStore)
	l.RegisterStorage(identifiers.Error, errorFilebaseStore)

	return l, nil
}
