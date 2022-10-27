package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
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
type Fields map[string]any

const (
	// LevelTrace is a finer-grained log level
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

// NewLogger returns a logger instance with prefix and flag
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)

	return &Logger{logger: l}
}

// clone returns a logger from Logger struct to avoid
// affecting the use of other processes
func (l *Logger) clone() *Logger {
	logger := *l

	return &logger
}

// WithFields returns a logger with fields
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

// WithContext returns a logger with context
func (l *Logger) WithContext(ctx context.Context) *Logger {
	logger := l.clone()
	logger.ctx = ctx

	return logger
}

// WithCaller returns a logger with caller, but caller message only for one line
func (l *Logger) WithCaller(depth int) *Logger {
	logger := l.clone()
	pc, file, line, ok := runtime.Caller(depth)
	if ok {
		f := runtime.FuncForPC(pc)
		logger.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return logger
}

// WithCallersFrames returns a logger with whole callstack message
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1

	var callers []string
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

// WithTrace returns a logger with trace_id and span_id
func (l *Logger) WithTrace() *Logger {
	ginCtx, ok := l.ctx.(*gin.Context)
	if ok {
		return l.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Span-ID"),
		})
	}

	return l
}
