package usecase_test

import (
	"fmt"
	"testing"

	"logur.dev/logur"
	"logur.dev/logur/logtesting"

	"iseage/bank/internal/usecase"
)

func TestLoggerAdapter_Levels(t *testing.T) {
	tests := map[string]struct {
		logFunc func(logger *usecase.LoggerAdapter, msg string, fields ...map[string]interface{})
	}{
		"trace": {
			logFunc: (*usecase.LoggerAdapter).Trace,
		},
		"debug": {
			logFunc: (*usecase.LoggerAdapter).Debug,
		},
		"info": {
			logFunc: (*usecase.LoggerAdapter).Info,
		},
		"warn": {
			logFunc: (*usecase.LoggerAdapter).Warn,
		},
		"error": {
			logFunc: (*usecase.LoggerAdapter).Error,
		},
	}

	for name, test := range tests {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			testLogger := &logur.TestLoggerFacade{}
			logger := usecase.NewLoggerAdapter(testLogger)

			test.logFunc(logger, fmt.Sprintf("message: %s", name))

			level, _ := logur.ParseLevel(name)

			event := logur.LogEvent{
				Level: level,
				Line:  "message: " + name,
			}

			logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
		})
	}
}

func TestLoggerAdapter_WithFields(t *testing.T) {
	testLogger := &logur.TestLoggerFacade{}

	var logger usecase.Logger = usecase.NewLoggerAdapter(testLogger)

	fields := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	logger = logger.WithFields(fields)

	logger.Debug("message")

	event := logur.LogEvent{
		Level:  logur.Debug,
		Line:   "message",
		Fields: fields,
	}

	logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
}
