package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/begmaroman/eth-services/types"
)

var _ types.Logger = (*Logrus)(nil)

type Logrus struct {
	logrus.FieldLogger
}

// NewLogrus creates a wrapped Logrus logger
func NewLogrus(logger logrus.FieldLogger) *Logrus {
	return &Logrus{
		FieldLogger: logger,
	}
}

// Trace is a shim stand-in for when we have real trace-level logging support
func (l *Logrus) Trace(args ...interface{}) {
	l.Debug(append([]interface{}{"TRACE: "}, args...))
}

// Tracef is a shim stand-in for when we have real trace-level logging support
func (l *Logrus) Tracef(format string, values ...interface{}) {
	l.Debugf("TRACE: %s", fmt.Sprintf(format, values...))
}

// Tracew is a shim stand-in for when we have real trace-level logging support
func (l *Logrus) Tracew(msg string, keysAndValues ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(keysAndValues); i += 2 {
		fields[fmt.Sprintf("%v", keysAndValues[i])] = keysAndValues[i+1]
	}

	l.WithFields(fields).Debug("TRACE: " + msg)
}

// Debugw is a shim stand-in for when we have real debug-level logging support
func (l *Logrus) Debugw(msg string, keysAndValues ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(keysAndValues); i += 2 {
		fields[fmt.Sprintf("%v", keysAndValues[i])] = keysAndValues[i+1]
	}

	l.WithFields(fields).Debug("DEBUG: " + msg)
}

// Infow is a shim stand-in for when we have real info-level logging support
func (l *Logrus) Infow(msg string, keysAndValues ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(keysAndValues); i += 2 {
		fields[fmt.Sprintf("%v", keysAndValues[i])] = keysAndValues[i+1]
	}

	l.WithFields(fields).Info("INFO: " + msg)
}

// Warnw is a shim stand-in for when we have real warn-level logging support
func (l *Logrus) Warnw(msg string, keysAndValues ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(keysAndValues); i += 2 {
		fields[fmt.Sprintf("%v", keysAndValues[i])] = keysAndValues[i+1]
	}

	l.WithFields(fields).Warn("WARN: " + msg)
}

// Errorw is a shim stand-in for when we have real error-level logging support
func (l *Logrus) Errorw(msg string, keysAndValues ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(keysAndValues); i += 2 {
		fields[fmt.Sprintf("%v", keysAndValues[i])] = keysAndValues[i+1]
	}

	l.WithFields(fields).Error("ERROR: " + msg)
}

// Panicw is a shim stand-in for when we have real panic-level logging support
func (l *Logrus) Panicw(msg string, keysAndValues ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(keysAndValues); i += 2 {
		fields[fmt.Sprintf("%v", keysAndValues[i])] = keysAndValues[i+1]
	}

	l.WithFields(fields).Panic("PANIC: " + msg)
}
