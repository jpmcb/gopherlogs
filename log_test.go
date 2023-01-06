package gopherlogs

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestNewLoggerWithWriter(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	//logger := NewLogger(false, 5, writer)
	logger := NewLogger(
		LoggerWithOutputWriter(writer),
		LoggerWithTty(false),
		LoggerWithLogVerbosity(5),
	)

	logger.Info("Generating Message in Info")
	logger.Info("Generating Message in Info again")
	writer.Flush()
	lineCount := len(strings.Split(strings.Trim(b.String(), "\n"), "\n"))
	if lineCount != 2 {
		t.Errorf("Expected 2 lines to be captured into buffer but found %d", lineCount)
	}
}

func TestNewLoggerDifferentWriters(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	logger := NewLogger(
		LoggerWithOutputWriter(writer),
		LoggerWithTty(false),
		LoggerWithLogVerbosity(5),
	)
	logger.Info("Generating Message in Info")
	logger.Info("Generating Message in Info again")
	logger = NewLogger(
		LoggerWithOutputWriter(os.Stdout),
		LoggerWithTty(false),
		LoggerWithLogVerbosity(5),
	)
	logger.Info("This should not be in the buffer")
	logger = NewLogger(
		LoggerWithOutputWriter(io.Discard),
		LoggerWithTty(false),
		LoggerWithLogVerbosity(5),
	)
	logger.Info("This should be discarded")
	logger = NewLogger(
		LoggerWithOutputWriter(writer),
		LoggerWithTty(false),
		LoggerWithLogVerbosity(5),
	)
	logger.Info("This should be captured in buffer writer")
	writer.Flush()
	lineCount := len(strings.Split(strings.Trim(b.String(), "\n"), "\n"))
	if lineCount != 3 {
		t.Errorf("Expected 2 lines to be captured into buffer but found %d", lineCount)
	}
}
