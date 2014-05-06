package gologger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	LevelTrace = "TRACE"
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
	LevelFatal = "FATAL"

	degreeTrace = 1
	degreeDebug = 2
	degreeInfo  = 3
	degreeWarn  = 4
	degreeError = 5
	degreeFatal = 6

	colorNormal = "\x1b[0m"
	colorRed    = "\x1b[31;3m"
	colorGreen  = "\x1b[32;3m"
	colorYellow = "\x1b[33;3m"
	colorPurple = "\x1b[34;3m"
	colorPink   = "\x1b[35;3m"
	colorAqua   = "\x1b[36;3m"

	timeFormatLayout string = "2006-01-02 15:04:05.000"

	logFormat string = "%s%s %-5s %s %s- "
)

type Logger struct {
	Name              string
	Level             string
	File              string
	OutputFileColored bool
}

func NewLogger(name, level, file string, outputFileColored bool) *Logger {
	return &Logger{Name: name, Level: level, File: file, OutputFileColored: outputFileColored}
}

func GetLogger(m map[string]string) Logger {
	return Logger{Name: m["Name"], Level: m["Level"], File: m["File"], OutputFileColored: m["OutputFileColored"] == "true"}
}

func (logger *Logger) Trace(messages ...interface{}) {
	logger.output(LevelTrace, messages...)
}

func (logger *Logger) Tracef(format string, messages ...interface{}) {
	logger.output(LevelTrace, fmt.Sprintf(format, messages...))
}

func (logger *Logger) Debug(messages ...interface{}) {
	logger.output(LevelDebug, messages...)
}

func (logger *Logger) Debugf(format string, messages ...interface{}) {
	logger.output(LevelDebug, fmt.Sprintf(format, messages...))
}

func (logger *Logger) Info(messages ...interface{}) {
	logger.output(LevelInfo, messages...)
}

func (logger *Logger) Infof(format string, messages ...interface{}) {
	logger.output(LevelInfo, fmt.Sprintf(format, messages...))
}

func (logger *Logger) Warn(messages ...interface{}) {
	logger.output(LevelWarn, messages...)
}

func (logger *Logger) Warnf(format string, messages ...interface{}) {
	logger.output(LevelWarn, fmt.Sprintf(format, messages...))
}

func (logger *Logger) Error(messages ...interface{}) {
	logger.output(LevelError, messages...)
}

func (logger *Logger) Errorf(format string, messages ...interface{}) {
	logger.output(LevelError, fmt.Sprintf(format, messages...))
}

func (logger *Logger) Fatal(messages ...interface{}) {
	logger.output(LevelFatal, messages...)
}

func (logger *Logger) Fatalf(format string, messages ...interface{}) {
	logger.output(LevelFatal, fmt.Sprintf(format, messages...))
}

func (logger *Logger) isOutput(degree int) bool {
	loggerDegree, _ := levelProperties(logger.Level)
	return loggerDegree <= degree
}

func (logger *Logger) output(level string, messages ...interface{}) {
	degree, color := levelProperties(level)
	if logger.isOutput(degree) {
		if logger.File != "" {
			file, err := os.OpenFile(logger.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Printf("Can't open the file: %s\n", logger.File)
				return
			}
			writer := bufio.NewWriter(file)
			fmt.Fprintf(writer, logFormat, "", time.Now().Format(timeFormatLayout), level, logger.Name, "")
			fmt.Fprintln(writer, messages...)
			writer.Flush()
			if logger.OutputFileColored {
				fileColoredPath := logger.File + ".colored"
				file, err = os.OpenFile(fileColoredPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
				if err != nil {
					fmt.Printf("Can't open the file: %s\n", fileColoredPath)
					return
				}
				writer := bufio.NewWriter(file)
				fmt.Fprintf(writer, logFormat, color, time.Now().Format(timeFormatLayout), level, logger.Name, colorNormal)
				fmt.Fprintln(writer, messages...)
				writer.Flush()
			}
		} else {
			fmt.Printf(logFormat, color, time.Now().Format(timeFormatLayout), level, logger.Name, colorNormal)
			fmt.Println(messages...)
		}
	}
}

func levelProperties(level string) (int, string) {
	degree, color := 0, colorNormal
	switch level {
	case LevelTrace:
		degree, color = degreeTrace, colorPurple
	case LevelDebug:
		degree, color = degreeDebug, colorAqua
	case LevelInfo:
		degree, color = degreeInfo, colorGreen
	case LevelWarn:
		degree, color = degreeWarn, colorYellow
	case LevelError:
		degree, color = degreeError, colorRed
	case LevelFatal:
		degree, color = degreeFatal, colorPink
	}
	return degree, color
}
