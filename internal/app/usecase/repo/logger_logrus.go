package repo

import (
	"os"

	"github.com/samber/do"
	"github.com/sirupsen/logrus"

	"github.com/MatthewBehnke/apis/internal/app/domain"
	"github.com/MatthewBehnke/apis/internal/app/usecase"
)

// Pattern to verify loggerLogrusImplem conforms to the required interfaces
var (
	_assertLoggerRepoImplem                    = &loggerLogrusImplem{}
	_                       usecase.LoggerRepo = _assertLoggerRepoImplem
)

func NewLoggerRepo(i *do.Injector) (usecase.LoggerRepo, error) {
	config := do.MustInvoke[*domain.Config](i)

	l := logrus.New()

	l.SetOutput(os.Stdout)

	switch config.Format {
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{})
	case "logfmt":
		l.SetFormatter(&logrus.TextFormatter{
			DisableColors:             config.NoColor,
			EnvironmentOverrideColors: true,
		})
	}

	level, err := logrus.ParseLevel(config.Level)

	if err != nil {
		return nil, err
	}

	l.SetLevel(level)

	return &loggerLogrusImplem{
		entry: logrus.NewEntry(l),
	}, nil
}

type loggerLogrusImplem struct {
	entry *logrus.Entry
}

// Trace implements the Logur Logger interface.
func (l *loggerLogrusImplem) Trace(msg string, fields ...map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.TraceLevel) {
		return
	}

	var entry = l.entry
	if len(fields) > 0 {
		entry = entry.WithFields(fields[0])
	}

	entry.Trace(msg)
}

// Debug implements the Logur Logger interface.
func (l *loggerLogrusImplem) Debug(msg string, fields ...map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.DebugLevel) {
		return
	}

	var entry = l.entry
	if len(fields) > 0 {
		entry = entry.WithFields(fields[0])
	}

	entry.Debug(msg)
}

// Info implements the Logur Logger interface.
func (l *loggerLogrusImplem) Info(msg string, fields ...map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.InfoLevel) {
		return
	}

	var entry = l.entry
	if len(fields) > 0 {
		entry = entry.WithFields(fields[0])
	}

	entry.Info(msg)
}

// Warn implements the Logur Logger interface.
func (l *loggerLogrusImplem) Warn(msg string, fields ...map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.WarnLevel) {
		return
	}

	var entry = l.entry
	if len(fields) > 0 {
		entry = entry.WithFields(fields[0])
	}

	entry.Warn(msg)
}

// Error implements the Logur Logger interface.
func (l *loggerLogrusImplem) Error(msg string, fields ...map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.ErrorLevel) {
		return
	}

	var entry = l.entry
	if len(fields) > 0 {
		entry = entry.WithFields(fields[0])
	}

	entry.Error(msg)
}
