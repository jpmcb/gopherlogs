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
	logger, err := NewLogger(
		WithOutputWriter(writer),
		WithTty(false),
		WithLogVerbosity(5),
	)
    if err != nil {
        t.Errorf("Expected NewLogger to not error: %s", err.Error())
    }

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
	logger, err := NewLogger(
		WithOutputWriter(writer),
		WithTty(false),
		WithLogVerbosity(5),
	)
    if err != nil {
        t.Errorf("Expected NewLogger to not error: %s", err.Error())
    }

	logger.Info("Generating Message in Info")
	logger.Info("Generating Message in Info again")
	logger, err = NewLogger(
		WithOutputWriter(os.Stdout),
		WithTty(false),
		WithLogVerbosity(5),
	)
    if err != nil {
        t.Errorf("Expected NewLogger to not error: %s", err.Error())
    }

	logger.Info("This should not be in the buffer")
	logger, err = NewLogger(
		WithOutputWriter(io.Discard),
		WithTty(false),
		WithLogVerbosity(5),
	)
    if err != nil {
        t.Errorf("Expected NewLogger to not error: %s", err.Error())
    }
	logger.Info("This should be discarded")

	logger, err = NewLogger(
		WithOutputWriter(writer),
		WithTty(false),
		WithLogVerbosity(5),
	)
    if err != nil {
        t.Errorf("Expected NewLogger to not error: %s", err.Error())
    }

	logger.Info("This should be captured in buffer writer")
	writer.Flush()
	lineCount := len(strings.Split(strings.Trim(b.String(), "\n"), "\n"))
	if lineCount != 3 {
		t.Errorf("Expected 2 lines to be captured into buffer but found %d", lineCount)
	}
}
