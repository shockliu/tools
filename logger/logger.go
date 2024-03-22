package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type (
	Level int
)

const (
	LevelFatal = iota
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

var (
	errColor   = color.New(color.BgRed, color.FgWhite).SprintfFunc()
	fatalColor = color.New(color.BgHiRed, color.FgWhite).SprintfFunc() //"\033[1;37;41m"
	// errColor     = "\033[1;37;45m"
	warnColor  = color.New(color.BgYellow, color.FgWhite).SprintfFunc() //"\033[1;37;43m"
	infoColor  = color.New(color.BgGreen, color.FgWhite).SprintfFunc()  //"\033[1;37;42m"
	debugColor = color.New(color.FgBlack, color.BgWhite).SprintfFunc()  //"\033[1;36m"
	entColor   = color.New(color.BgHiBlue, color.FgWhite).SprintfFunc()
	modelColor = color.New(color.BgCyan, color.FgWhite).SprintfFunc() //"\033[1;37;43m"
)

var logger = New()

func Fatal(v ...interface{}) {
	logger.Output(LevelFatal, fmt.Sprint(v...))
	//p(LevelFatal, v...)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	logger.Output(LevelFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Error(v ...interface{}) {
	logger.Output(LevelError, fmt.Sprint(v...))
	//	p(LevelError, v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Output(LevelError, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	logger.Output(LevelWarning, fmt.Sprint(v...))
	//	p(LevelWarning, v...)
}

func Warnf(format string, v ...interface{}) {
	logger.Output(LevelWarning, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	logger.Output(LevelInfo, fmt.Sprint(v...))
	//	p(LevelInfo, v...)
}

func Infof(format string, v ...interface{}) {
	logger.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	logger.Output(LevelDebug, fmt.Sprint(v...))
	//	p(LevelDebug, v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Output(LevelDebug, fmt.Sprintf(format, v...))
}

func Check(s string, err error) {
	if err == nil {
		logger.Output(LevelInfo, s+" ok")
	} else {
		logger.Output(LevelFatal, fmt.Sprintf("%s %+v", s, err))
		os.Exit(1)
	}
}

func SetLogLevel(level Level) {
	logger.SetLogLevel(level)
}

type logManager struct {
	_log *log.Logger
	//小于等于该级别的level才会被记录
	logLevel Level
	_model   string
}

// NewLogger 实例化，供自定义
func NewLogger() *logManager {
	return &logManager{_log: log.New(os.Stderr, fmt.Sprintf("[%s] ", entColor("DQK")), log.Lshortfile|log.LstdFlags), logLevel: LevelDebug}
}

// New 实例化，供外部直接调用 log.XXXX
func New() *logManager {
	//return &logManager{_log: log.New(os.Stderr, "", log.Lshortfile|log.LstdFlags), logLevel: LevelDebug}
	return &logManager{_log: log.New(os.Stdout, fmt.Sprintf("[%s] ", entColor("DQK")), log.Lshortfile|log.LstdFlags), logLevel: LevelDebug}
}

func (l *logManager) Output(level Level, s string) error {
	if l.logLevel < level {
		return nil
	}
	if len(l._model) > 0 {
		s = fmt.Sprintf("|%s| %s", modelColor(l._model), s)
	}
	switch level {
	case LevelFatal:
		s = fmt.Sprintf("[%s] %s", fatalColor("FATAL"), s)
	case LevelError:
		s = fmt.Sprintf("[%s] %s", errColor("ERROR"), s)
	case LevelWarning:
		s = fmt.Sprintf("[%s] %s", warnColor("WARN"), s)
	case LevelInfo:
		s = fmt.Sprintf("[%s] %s", infoColor("INFO"), s)
	case LevelDebug:
		s = fmt.Sprintf("[%s] %s", debugColor("DEBUG"), s)
	default:
		s = fmt.Sprintf("[%s] %s", infoColor("UNKNOW"), s)
	}
	return l._log.Output(3, s)
}

func (l *logManager) Fatal(s string) {
	l.Output(LevelFatal, s)
	os.Exit(1)
}

func (l *logManager) Fatalf(format string, v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *logManager) Error(s string) {
	l.Output(LevelError, s)
}

func (l *logManager) Errorf(format string, v ...interface{}) {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *logManager) Warn(s string) {
	l.Output(LevelWarning, s)
}

func (l *logManager) Warnf(format string, v ...interface{}) {
	l.Output(LevelWarning, fmt.Sprintf(format, v...))
}

func (l *logManager) Info(s string) {
	l.Output(LevelInfo, s)
}

func (l *logManager) Infof(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *logManager) Debug(s string) {
	l.Output(LevelDebug, s)
}

func (l *logManager) Debugf(format string, v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *logManager) Check(s string, err error) {
	if err != nil {
		l.Output(LevelInfo, s+"ok")
	} else {
		l.Output(LevelWarning, s+fmt.Sprintf("%+v", err))
	}
}

func (l *logManager) SetLogLevel(level Level) *logManager {
	l.logLevel = level
	return l
}

func (l *logManager) SetLogModel(model string) *logManager {
	l._model = model
	return l
}

func SetLogModel(model string) {
	logger._model = model
}
