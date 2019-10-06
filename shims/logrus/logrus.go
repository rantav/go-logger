package logrus

import (
	originlog "github.com/InVisionApp/go-logger"
	log "github.com/rantav/go-logger"
	"github.com/sirupsen/logrus"
)

type shim struct {
	*logrus.Entry
}

// New can be used to override the default logger.
// Optionally pass in an existing logrus logger or pass in
// `nil` to use the default logger.
func New(logger *logrus.Logger) log.Logger {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	return &shim{logrus.NewEntry(logger)}
}

// WithFields will return a new logger based on the original logger with
// the additional supplied fields. Wrapper for logrus Entry.WithFields()
func (s *shim) WithFields(fields log.Fields) log.Logger {
	cp := &shim{
		s.Entry.WithFields(logrus.Fields(fields)),
	}
	return cp
}

func (s *shim) AsInvisionLogger() originlog.Logger {
	return &originshim{s.Entry}
}

type originshim struct {
	*logrus.Entry
}

func (s *originshim) WithFields(fields originlog.Fields) originlog.Logger {
	cp := &originshim{
		s.Entry.WithFields(logrus.Fields(fields)),
	}
	return cp
}
