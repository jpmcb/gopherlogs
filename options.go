package gopherlogs

import (
	"io"
	"os"

	"github.com/jpmcb/gopherlogs/pkg/colors"
)

// LoggerOptions an option to be passed to the NewLogger constructor
type LoggerOptions interface {
	apply(*CMDLogger)
}

type loggerOptionAdapter func(*CMDLogger)

func (l loggerOptionAdapter) apply(o *CMDLogger) {
	l(o)
}

func LoggerWithTty(isTty bool) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) {
		l.tty = isTty
	})
}

func LoggerWithLogVerbosity(verbosity int) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) {
		l.verbosity = verbosity
	})
}

func LoggerWithLogLevel(logLevel int) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) {
		l.logLevel = logLevel
	})
}

func LoggerWithLeftPadIndent(indent int) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) {
		l.indent = indent
	})
}

func LoggerWithTermFileDescriptor(termFd int) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) {
		l.termFd = termFd
	})
}

func LoggerWithColor(color colors.Attribute) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) {
		l.logColor = color
	})
}

func LoggerWithOutputWriter(writer io.Writer) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) {
		l.output = writer
		l.defaultWriter = writer
	})
}

func LoggerWithLogFile(filePath string) LoggerOptions {
	return loggerOptionAdapter(func(l *CMDLogger) {
		logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			l.Warnf("Failed to open log file %q: %v", filePath, err)
			return
		}

		// Set the output to the new multiwriter while keeping the default writer
		l.output = io.MultiWriter(logFile, l.defaultWriter)
	})
}
