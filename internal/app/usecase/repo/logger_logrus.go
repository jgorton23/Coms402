package repo

import (
	"os"

	"github.com/sirupsen/logrus"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

// Pattern to verify loggerLogrusImplem conforms to the required interfaces
var (
	_assertLoggerLogrusImplem                    = &loggerLogrusImplem{}
	_                         usecase.LoggerRepo = _assertLoggerLogrusImplem
)

// LoggerLogrusBuilder creates a new loggerLogrusBuilder
func LoggerLogrusBuilder() *loggerLogrusBuilder {
	return &loggerLogrusBuilder{
		logger: &loggerLogrusImplem{},
		opts: struct {
			format string
			color  bool
			level  string
		}{
			format: "json",
			color:  false,
			level:  "info",
		}}
}

type loggerLogrusBuilder struct {
	logger *loggerLogrusImplem
	opts   struct {
		format string
		color  bool
		level  string
	}
}

// WithFormat allows the logging formatter to be specified
// The default formatter is json
func (builder *loggerLogrusBuilder) WithFormat(format string) *loggerLogrusBuilder {
	builder.opts.format = format
	return builder
}

// WithColor enables colored logging
// Color is disabled by default
func (builder *loggerLogrusBuilder) WithColor(color bool) *loggerLogrusBuilder {
	builder.opts.color = color
	return builder
}

// WithLoggerLogrusLevel sets the logging level
// The default log level is "info"
func (builder *loggerLogrusBuilder) WithLevel(level string) *loggerLogrusBuilder {
	builder.opts.level = level
	return builder
}

// New creates a new usecase.LoggerRepo using logrus based on the builder options
func (builder *loggerLogrusBuilder) New() (usecase.LoggerRepo, error) {
	l := logrus.New()
	l.SetOutput(os.Stdout)

	switch builder.opts.format {
	case "logfmt":
		l.SetFormatter(&logrus.TextFormatter{
			DisableColors:             builder.opts.color,
			EnvironmentOverrideColors: true,
		})
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{})
	}

	level, err := logrus.ParseLevel(builder.opts.level)

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
