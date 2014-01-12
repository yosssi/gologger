# gologger - Logger in Go

![Console](./images/console.png)

gologger brings a logging function to the Go language.

## Installation

	$ go get github.com/yosssi/gologger.git

## Examples

```Go
package main

import (
	"github.com/yosssi/gologger"
)

func main() {
	// This logger outputs the logs to the stdout.
	logger := gologger.Logger{Name: "LoggerTest", Level: gologger.LevelInfo}
	logger.Info("This is a logger test.")

	// This logger outputs the logs to "./test.log".
	loggerFile := gologger.Logger{Name: "LoggerTest", Level: gologger.LevelInfo, File: "./test.log"}
	loggerFile.Info("This is a logger test. This message will be shown on the log file.")
}
```
