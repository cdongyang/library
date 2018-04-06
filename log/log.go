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

func (l *Logger) outputln(level int, args ...interface{}) {
	if l.level < level {
		return
	}
	l.Output(l.calldepth, fmt.Sprintln(args...))
}

func (l *Logger) output(level int, args ...interface{}) {
	if l.level < level {
		return
	}
	l.Output(l.calldepth, fmt.Sprint(args...))
}

func (l *Logger) outputf(level int, format string, args ...interface{}) {
	if l.level < level {
		return
	}
	l.Output(l.calldepth, fmt.Sprintf(format, args...))
}

func (l *Logger) Panicln(args ...interface{}) {
	args = append(args, "[PANIC] ")
	l.outputln(LOG_LEVEL_PANIC, args...)
	panic(fmt.Sprintln(args...))
}

func (l *Logger) Panic(args ...interface{}) {
	args = append(args, "[PANIC] ")
	l.output(LOG_LEVEL_PANIC, args...)
	panic(fmt.Sprint(args...))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	args = append(args, "[PANIC] ")
	l.outputf(LOG_LEVEL_PANIC, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *Logger) Debugln(args ...interface{}) {
	args = append(args, "[DEBUG] ")
	l.outputln(LOG_LEVEL_DEBUG, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	args = append(args, "[DEBUG] ")
	l.output(LOG_LEVEL_DEBUG, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	args = append(args, "[DEBUG] ")
	l.outputf(LOG_LEVEL_DEBUG, format, args...)
}

func (l *Logger) Logln(args ...interface{}) {
	args = append(args, "[LOG] ")
	l.outputln(LOG_LEVEL_LOG, args...)
}

func (l *Logger) Log(args ...interface{}) {
	args = append(args, "[LOG] ")
	l.output(LOG_LEVEL_LOG, args...)
}

func (l *Logger) Logf(format string, args ...interface{}) {
	args = append(args, "[LOG] ")
	l.outputf(LOG_LEVEL_LOG, format, args...)
}

func (l *Logger) Noticeln(args ...interface{}) {
	args = append(args, "[NOTICE] ")
	l.outputln(LOG_LEVEL_NOTICE, args...)
}

func (l *Logger) Notice(args ...interface{}) {
	args = append(args, "[NOTICE] ")
	l.output(LOG_LEVEL_NOTICE, args...)
}

func (l *Logger) Noticef(format string, args ...interface{}) {
	args = append(args, "[NOTICE] ")
	l.outputf(LOG_LEVEL_NOTICE, format, args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	args = append(args, "[WARN] ")
	l.outputln(LOG_LEVEL_WARN, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	args = append(args, "[WARN] ")
	l.output(LOG_LEVEL_WARN, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	args = append(args, "[WARN] ")
	l.outputf(LOG_LEVEL_WARN, format, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	args = append(args, "[ERROR] ")
	l.outputln(LOG_LEVEL_ERROR, args...)
}

func (l *Logger) Error(args ...interface{}) {
	args = append(args, "[ERROR] ")
	l.output(LOG_LEVEL_ERROR, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	args = append(args, "[ERROR] ")
	l.outputf(LOG_LEVEL_ERROR, format, args...)
}
