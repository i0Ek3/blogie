package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

// Logger defines a logger struct
type Logger struct {
	// a logger instance
	logger *log.Logger
	// log context
	ctx context.Context
	// log fields
	fields Fields
	// call details
	callers []string
}

// Level denotes a logger level
type Level int8

// Fields denotes a logger field
type Fields map[string]interface{}

const (
	// Trace level is a finer-grained log level than Debug
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// If a type implements the String() method, then the String()
// is automatically called when using fmt.Xxxx() functions

// String returns the string represent of log level
func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

// NewLogger creates a logger instance with prefix and flag
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{logger: l}
}

// clone clones a new logger from Logger struct to avoid
// affecting the use of other processes
func (l *Logger) clone() *Logger {
	logger := *l
	return &logger
}

// WithFields creates a logger with fields
func (l *Logger) WithFields(f Fields) *Logger {
	logger := l.clone()
	if logger.fields == nil {
		logger.fields = make(Fields)
	}
	for k, v := range f {
		logger.fields[k] = v
	}
	return logger
}

// WithContext creates a logger with context
func (l *Logger) WithContext(ctx context.Context) *Logger {
	logger := l.clone()
	logger.ctx = ctx
	return logger
}

// WithCaller creates a logger with caller, but caller message only for one line
func (l *Logger) WithCaller(depth int) *Logger {
	logger := l.clone()
	pc, file, line, ok := runtime.Caller(depth)
	if ok {
		f := runtime.FuncForPC(pc)
		logger.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return logger
}

// WithCallersFrames creates a logger with whole callstack message
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1

	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}

	logger := l.clone()
	logger.callers = callers
	return logger
}

// JSONFormat formats the given information into a log field
func (l *Logger) JSONFormat(level Level, msg string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4) // why here need plus 4?
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = msg
	data["callers"] = l.callers

	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

// Output outputs corresponding message of log level
func (l *Logger) Output(level Level, msg string) {
	// call WithCaller(3) to generate a copy of the global Logger,
	// and refresh the call stack information in the copy
	body, _ := json.Marshal(l.WithCaller(3).JSONFormat(level, msg))
	content := string(body)
	switch level {
	case LevelTrace:
		l.logger.Print(content)
	case LevelDebug:
		l.logger.Print(content)
	case LevelInfo:
		l.logger.Print(content)
	case LevelWarn:
		l.logger.Print(content)
	case LevelError:
		l.logger.Print(content)
	case LevelFatal:
		l.logger.Fatal(content)
	case LevelPanic:
		l.logger.Panic(content)
	}
}

// Trace logs a message at Trace level
func (l *Logger) Trace(v ...interface{}) {
	l.Output(LevelTrace, fmt.Sprint(v...))
}

// Tracef logs a message at Trace level
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Output(LevelTrace, fmt.Sprintf(format, v...))
}

// Info logs a message at Info level
func (l *Logger) Info(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

// Infof logs a message at Info level
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

// Debug logs a message at Debug level
func (l *Logger) Debug(v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprint(v...))
}

// Debugf logs a message at Debug level
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprintf(format, v...))
}

// Warn logs a message at Warn level
func (l *Logger) Warn(v ...interface{}) {
	l.Output(LevelWarn, fmt.Sprint(v...))
}

// Warnf logs a message at Warn level
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(LevelWarn, fmt.Sprintf(format, v...))
}

// Error logs a message at Error level
func (l *Logger) Error(v ...interface{}) {
	l.Output(LevelError, fmt.Sprint(v...))
}

// Errorf logs a message at Error level
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}

// Fatal logs a message at Fatal level
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

// Fatalf logs a message at Fatal level
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}

// Panic logs a message at Panic level
func (l *Logger) Panic(v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprint(v...))
}

// Panicf logs a message at Panic level
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprintf(format, v...))
}
