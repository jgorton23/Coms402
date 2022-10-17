package usecase_test

// import (
// 	"fmt"
// 	"testing"

// 	"logur.dev/logur"
// 	"logur.dev/logur/logtesting"

// 	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
// )

// func TestLoggerUseCase_Levels(t *testing.T) {
// 	tests := map[string]struct {
// 		logFunc func(logger *usecase.LoggerUseCase, msg string, fields ...map[string]interface{})
// 	}{
// 		"trace": {
// 			logFunc: (*usecase.LoggerUseCase).Trace,
// 		},
// 		"debug": {
// 			logFunc: (*usecase.LoggerUseCase).Debug,
// 		},
// 		"info": {
// 			logFunc: (*usecase.LoggerUseCase).Info,
// 		},
// 		"warn": {
// 			logFunc: (*usecase.LoggerUseCase).Warn,
// 		},
// 		"error": {
// 			logFunc: (*usecase.LoggerUseCase).Error,
// 		},
// 	}

// 	for name, test := range tests {
// 		name, test := name, test

// 		t.Run(name, func(t *testing.T) {
// 			testLogger := &logur.TestLoggerFacade{}
// 			logger := usecase.NewLoggerUseCase(testLogger)

// 			test.logFunc(logger, fmt.Sprintf("message: %s", name))

// 			level, _ := logur.ParseLevel(name)

// 			event := logur.LogEvent{
// 				Level: level,
// 				Line:  "message: " + name,
// 			}

// 			logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
// 		})
// 	}
// }

// func TestLoggerUseCase_WithFields(t *testing.T) {
// 	testLogger := &logur.TestLoggerFacade{}

// 	var logger usecase.Logger = usecase.NewLoggerUseCase(testLogger)

// 	fields := map[string]interface{}{
// 		"key1": "value1",
// 		"key2": "value2",
// 	}

// 	logger = logger.WithFields(fields)

// 	logger.Debug("message")

// 	event := logur.LogEvent{
// 		Level:  logur.Debug,
// 		Line:   "message",
// 		Fields: fields,
// 	}

// 	logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
// }
