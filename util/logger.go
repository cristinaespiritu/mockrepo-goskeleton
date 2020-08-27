// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var formatLogFileName = "log-2006-01-02"
var runtimeGoosWindows = "windows"
var errorFailOpenMsg = "Failed to open log file %s for output: %s"
var fileExtension = ".json"
var tempFolderPath = "/tmp/"

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()
	var standardLogger = &StandardLogger{baseLogger}
	standardLogger.Formatter = &logrus.JSONFormatter{}

	var logfolder string
	if runtime.GOOS == runtimeGoosWindows {
		logfolder = "\\log\\"
	} else {
		logfolder = "/log/"
	}

	rootDir := RootDir()

	if !strings.Contains(rootDir, tempFolderPath) {
		finalLogPath := path.Join(rootDir + logfolder)

		if _, err := os.Stat(finalLogPath); os.IsNotExist(err) {
			err := os.Mkdir(finalLogPath, os.ModeDir)

			if err != nil {
				log.Fatal(err)
			}
		}

		fileID := time.Now().Format(formatLogFileName)
		logLocation := filepath.Join(finalLogPath, fileID+fileExtension)
		logFile, err := os.OpenFile(logLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf(errorFailOpenMsg, logLocation, err)
		}

		standardLogger.SetOutput(io.MultiWriter(os.Stderr, logFile))
	}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	invalidArgMessage      = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for arg: %s: %v"}
	missingArgMessage      = Event{3, "Missing arg: %s"}
)

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *StandardLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}

// Error is a standard error message
func (l *StandardLogger) Error(message string) {
	l.Errorf("Error : %s", message)
}

// Info is a standard info message
func (l *StandardLogger) Info(message string) {
	l.Infof("Informational : %s", message)
}

// Warning is a standard info message
func (l *StandardLogger) Warning(message string) {
	l.Warningf("Warning : %s", message)
}
