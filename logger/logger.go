package logger

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"time"
)

func Info(obj ...interface{}) {
	// create the logger
	logger := logrus.New()
	logger.SetReportCaller(true)
	// with Json Formatter
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(os.Stdout)

	currentTime := time.Now()
	logFileName := "storage/logs/jumia_" + currentTime.Format("2006-01-02") + ".txt"
	file, file_err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if file_err != nil {
		logger.Fatal(file_err)
	}

	logger.SetOutput(file)
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	pc, fn, line, _ := runtime.Caller(1)
	logger.Info(fmt.Sprintf(" %s[%s:%d] %s", runtime.FuncForPC(pc).Name(), fn, line, jsonObj))
	defer file.Close()
}
