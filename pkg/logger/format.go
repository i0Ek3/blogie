package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// JSONFormat formats the given information into a log field
func (l *Logger) JSONFormat(level Level, msg string) map[string]any {
	// plus 4 means to add four attributes of data
	data := make(Fields, len(l.fields)+4)
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
func (l *Logger) Trace(ctx context.Context, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelTrace, fmt.Sprint(v...))
}

// Tracef logs a message at Trace level
func (l *Logger) Tracef(ctx context.Context, format string, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelTrace, fmt.Sprintf(format, v...))
}

// Info logs a message at Info level
func (l *Logger) Info(ctx context.Context, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprint(v...))
}

// Infof logs a message at Info level
func (l *Logger) Infof(ctx context.Context, format string, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprintf(format, v...))
}

// Debug logs a message at Debug level
func (l *Logger) Debug(ctx context.Context, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprint(v...))
}

// Debugf logs a message at Debug level
func (l *Logger) Debugf(ctx context.Context, format string, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprintf(format, v...))
}

// Warn logs a message at Warn level
func (l *Logger) Warn(ctx context.Context, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelWarn, fmt.Sprint(v...))
}

// Warnf logs a message at Warn level
func (l *Logger) Warnf(ctx context.Context, format string, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelWarn, fmt.Sprintf(format, v...))
}

// Error logs a message at Error level
func (l *Logger) Error(ctx context.Context, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprint(v...))
}

// Errorf logs a message at Error level
func (l *Logger) Errorf(ctx context.Context, format string, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprintf(format, v...))
}

// Fatal logs a message at Fatal level
func (l *Logger) Fatal(ctx context.Context, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprint(v...))
}

// Fatalf logs a message at Fatal level
func (l *Logger) Fatalf(ctx context.Context, format string, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprintf(format, v...))
}

// Panic logs a message at Panic level
func (l *Logger) Panic(ctx context.Context, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprint(v...))
}

// Panicf logs a message at Panic level
func (l *Logger) Panicf(ctx context.Context, format string, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprintf(format, v...))
}
