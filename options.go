package gopherlogs

import (
	"io"
	"os"

	"github.com/jpmcb/gopherlogs/pkg/colors"
)

// LoggerOptions an option to be passed to the NewLogger constructor
type LoggerOptions interface {
	apply(*CMDLogger) error
}

type loggerOptionAdapter func(*CMDLogger) error

func (l loggerOptionAdapter) apply(o *CMDLogger) error {
	err := l(o)
	return err
}

func WithTty(isTty bool) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) error {
		l.tty = isTty
		return nil
	})
}

func WithLogVerbosity(verbosity int) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) error {
		l.verbosity = verbosity
		return nil
	})
}

func WithLogLevel(logLevel int) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) error {
		l.logLevel = logLevel
		return nil
	})
}

func WithLeftPadIndent(indent int) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) error {
		l.indent = indent
		return nil
	})
}

func WithTermFileDescriptor(termFd int) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) error {
		l.termFd = termFd
		return nil
	})
}

func WithColor(color colors.Attribute) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) error {
		l.logColor = color
		return nil
	})
}

func WithOutputWriter(writer io.Writer) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) error {
		l.output = writer
		l.defaultWriter = writer
		return nil
	})
}

func WithLogFile(filePath string) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) error {
		logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			l.Warnf("Failed to open log file %q: %v", filePath, err)
			return err
		}

		// Set the output to the new multiwriter while keeping the default writer
		l.output = io.MultiWriter(logFile, l.defaultWriter)
		return nil
	})
}
