package log

import (
	"fmt"
	"log"
)

const (
	LOG_LEVEL_NO = iota
	LOG_LEVEL_ERROR
	LOG_LEVEL_WARN
	LOG_LEVEL_NOTICE
	LOG_LEVEL_LOG
	LOG_LEVEL_DEBUG
	LOG_LEVEL_PANIC
)

type Logger struct {
	*log.Logger
	calldepth int
	level     int
}

func New(logger *log.Logger, calldepth int, level int) *Logger {
	return &Logger{Logger: logger, calldepth: calldepth, level: level}
}

func (l *Logger) GetLevel() int {
	return l.level
}

func (l *Logger) SetLevel(level int) {
	l.level = level
}

func (l *Logger) GetCalldepth() int {
	return l.calldepth
}

func (l *Logger) SetCalldepth(calldepth int) {
	l.calldepth = calldepth
}

func (l *Logger) outputln(level int, info string, args ...interface{}) {
	if l.level < level {
		return
	}
	cpargs := make([]interface{}, 0, len(args)+1)
	cpargs = append(cpargs, info)
	cpargs = append(cpargs, args...)
	l.Output(l.calldepth, fmt.Sprintln(cpargs...))
}

func (l *Logger) output(level int, info string, args ...interface{}) {
	if l.level < level {
		return
	}
	cpargs := make([]interface{}, 0, len(args)+1)
	cpargs = append(cpargs, info)
	cpargs = append(cpargs, args...)
	l.Output(l.calldepth, fmt.Sprint(cpargs...))
}

func (l *Logger) outputf(level int, info string, format string, args ...interface{}) {
	if l.level < level {
		return
	}
	l.Output(l.calldepth, fmt.Sprintf(info+format, args...))
}

func (l *Logger) Panicln(args ...interface{}) {
	l.outputln(LOG_LEVEL_PANIC, "[PANIC] ", args...)
	panic(fmt.Sprintln(args...))
}

func (l *Logger) Panic(args ...interface{}) {
	l.output(LOG_LEVEL_PANIC, "[PANIC] ", args...)
	panic(fmt.Sprint(args...))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.outputf(LOG_LEVEL_PANIC, "[PANIC] ", format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *Logger) Debugln(args ...interface{}) {
	l.outputln(LOG_LEVEL_DEBUG, "[DEBUG] ", args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.output(LOG_LEVEL_DEBUG, "[DEBUG] ", args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.outputf(LOG_LEVEL_DEBUG, "[DEBUG] ", format, args...)
}

func (l *Logger) Logln(args ...interface{}) {
	l.outputln(LOG_LEVEL_LOG, "[LOG] ", args...)
}

func (l *Logger) Log(args ...interface{}) {
	l.output(LOG_LEVEL_LOG, "[LOG] ", args...)
}

func (l *Logger) Logf(format string, args ...interface{}) {
	l.outputf(LOG_LEVEL_LOG, "[LOG] ", format, args...)
}

func (l *Logger) Noticeln(args ...interface{}) {
	l.outputln(LOG_LEVEL_NOTICE, "[NOTICE] ", args...)
}

func (l *Logger) Notice(args ...interface{}) {
	l.output(LOG_LEVEL_NOTICE, "[NOTICE] ", args...)
}

func (l *Logger) Noticef(format string, args ...interface{}) {
	l.outputf(LOG_LEVEL_NOTICE, "[NOTICE] ", format, args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.outputln(LOG_LEVEL_WARN, "[WARN] ", args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.output(LOG_LEVEL_WARN, "[WARN] ", args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.outputf(LOG_LEVEL_WARN, "[WARN] ", format, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.outputln(LOG_LEVEL_ERROR, "[ERROR] ", args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.output(LOG_LEVEL_ERROR, "[ERROR] ", args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.outputf(LOG_LEVEL_ERROR, "[ERROR] ", format, args...)
}
