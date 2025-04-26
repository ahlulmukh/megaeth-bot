package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

var (
	infoColor    = color.New(color.FgCyan).SprintFunc()
	successColor = color.New(color.FgGreen).SprintFunc()
	errorColor   = color.New(color.FgRed).SprintFunc()
	warningColor = color.New(color.FgYellow).SprintFunc()
	debugColor   = color.New(color.FgHiBlue).SprintFunc()
	processColor = color.New(color.FgHiMagenta).SprintFunc()
)

func (l *Logger) log(_, emoji string, message string, c func(a ...interface{}) string, current, total *int) {
	now := time.Now()
	timestamp := now.Format("02-01-2006 15:04:05")

	var status string
	if current != nil && total != nil {
		status = fmt.Sprintf("[%d/%d] ", *current, *total)
	}

	logText := c(fmt.Sprintf("%s %s%s", emoji, status, message))
	fmt.Printf("[%s] %s\n", timestamp, logText)
}

func (l *Logger) Info(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	l.log("info", "[i]", message, infoColor, nil, nil)
}

func (l *Logger) Success(message string) {
	l.log("success", "[✓]", message, successColor, nil, nil)
}

func (l *Logger) Error(message string) {
	l.log("error", "[-]", message, errorColor, nil, nil)
}

func (l *Logger) Warning(message string) {
	l.log("warning", "[!]", message, warningColor, nil, nil)
}

func (l *Logger) Debug(message string) {
	l.log("debug", "[*]", message, debugColor, nil, nil)
}

func (l *Logger) Process(current, total int, message string) {
	l.log("process", "[>]", message, processColor, &current, &total)
}
