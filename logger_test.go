package gologger

import (
	"strconv"
	"testing"
)

const (
	file string = "./test.log"
)

func TestTrace(t *testing.T) {
	logger := Logger{Name: "TestTrace01", Level: LevelTrace}
	logger.Trace("This message should be shown.")

	logger = Logger{Name: "TestTrace02", Level: LevelDebug}
	logger.Trace("This message should not be shown.")

	logger = Logger{Name: "TestTrace03", Level: LevelTrace, File: file}
	logger.Trace("This message should be shown.")

	logger = Logger{Name: "TestTrace04", Level: LevelDebug, File: file}
	logger.Trace("This message should not be shown.")
}

func TestDebug(t *testing.T) {
	logger := Logger{Name: "TestDebug01", Level: LevelDebug}
	logger.Debug("This message should be shown.")

	logger = Logger{Name: "TestDebug02", Level: LevelInfo}
	logger.Debug("This message should not be shown.")

	logger = Logger{Name: "TestDebug03", Level: LevelDebug, File: file}
	logger.Debug("This message should be shown.")

	logger = Logger{Name: "TestDebug04", Level: LevelInfo, File: file}
	logger.Debug("This message should not be shown.")
}

func TestInfo(t *testing.T) {
	logger := Logger{Name: "TestInfo01", Level: LevelInfo}
	logger.Info("This message should be shown.")

	logger = Logger{Name: "TestInfo02", Level: LevelWarn}
	logger.Info("This message should not be shown.")

	logger = Logger{Name: "TestInfo03", Level: LevelInfo, File: file}
	logger.Info("This message should be shown.")

	logger = Logger{Name: "TestInfo04", Level: LevelWarn, File: file}
	logger.Info("This message should not be shown.")
}

func TestWarn(t *testing.T) {
	logger := Logger{Name: "TestWarn01", Level: LevelWarn}
	logger.Warn("This message should be shown.")

	logger = Logger{Name: "TestWarn02", Level: LevelError}
	logger.Warn("This message should not be shown.")

	logger = Logger{Name: "TestWarn03", Level: LevelWarn, File: file}
	logger.Warn("This message should be shown.")

	logger = Logger{Name: "TestWarn04", Level: LevelError, File: file}
	logger.Warn("This message should not be shown.")
}

func TestError(t *testing.T) {
	logger := Logger{Name: "TestError01", Level: LevelError}
	logger.Error("This message should be shown.")

	logger = Logger{Name: "TestError02", Level: LevelFatal}
	logger.Error("This message should not be shown.")

	logger = Logger{Name: "TestError03", Level: LevelError, File: file}
	logger.Error("This message should be shown.")

	logger = Logger{Name: "TestError04", Level: LevelFatal, File: file}
	logger.Error("This message should not be shown.")
}

func TestFatal(t *testing.T) {
	logger := Logger{Name: "TestFatal01", Level: LevelFatal}
	logger.Fatal("This message should be shown.", "This is an optional message.")

	logger = Logger{Name: "TestFatal02", Level: LevelFatal, File: file}
	logger.Fatal("This message should be shown.", "This is an optional message.")
}

func TestIsOutput(t *testing.T) {
	logger := Logger{Name: "TestIsOutput", Level: LevelInfo}

	if logger.isOutput(degreeDebug) {
		t.Error("Expected isOutput to be false.")
	}

	if !logger.isOutput(degreeInfo) {
		t.Error("Expected isOutput to be true.")
	}

	if !logger.isOutput(degreeWarn) {
		t.Error("Expected isOutput to be true.")
	}
}

func TestOutput(t *testing.T) {
	logger := Logger{Name: "TestOutput01", Level: LevelInfo}
	logger.output(LevelDebug, "This message should not be shown.")

	logger = Logger{Name: "TestOutput02", Level: LevelInfo}
	logger.output(LevelInfo, "This message should be shown on the stdout.")

	logger = Logger{Name: "TestOutput03", Level: LevelInfo, File: "/a/b/c/d/e/f/g/h"}
	logger.output(LevelInfo, "An error should occur and this message should not be shown.")

	logger = Logger{Name: "TestTrace04", Level: LevelInfo, File: file}
	logger.Info("This message should be shown on the file:", file)

	logger = Logger{Name: "TestLogger", Level: LevelTrace, File: file, OutputFileColored: true}
	logger.Trace("This branch was passed.")
	logger.Debug("The value of the parameter \"name\":", "Taro")
	logger.Info("The server is listening on port 8080.")
	logger.Warn("The method \"getName()\" is deprecated.")
	logger.Error("Could not connect to the database.")
	logger.Fatal("An application error occurred.")
}

func TestLevelProperties(t *testing.T) {
	checkDegreeColor(t, "", 0, colorNormal)
	checkDegreeColor(t, LevelTrace, degreeTrace, colorPurple)
	checkDegreeColor(t, LevelDebug, degreeDebug, colorAqua)
	checkDegreeColor(t, LevelInfo, degreeInfo, colorGreen)
	checkDegreeColor(t, LevelWarn, degreeWarn, colorYellow)
	checkDegreeColor(t, LevelError, degreeError, colorRed)
	checkDegreeColor(t, LevelFatal, degreeFatal, colorPink)
}

func checkDegreeColor(t *testing.T, level string, degreeExpected int, colorExpected string) {
	degree, color := levelProperties(level)
	if degree != degreeExpected || color != colorExpected {
		t.Error("degree should be " + strconv.Itoa(degreeExpected) + " and color should be " + colorExpected + ".")
	}
}
