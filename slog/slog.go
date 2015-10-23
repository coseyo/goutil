package slog

import (
	"fmt"
	"os"
	"time"
)

func LogToFile(filename string, v ...interface{}) error {

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	logString := fmt.Sprintf("%s : %v \n\n", time.Now().Format("2006-01-02 15:04:05"), v)
	if _, err := f.WriteString(logString); err != nil {
		return err
	}

	return nil
}

func SimpleLog(app string, v ...interface{}) error {

	logDir := fmt.Sprintf("/data/logs/apps/%s", app)
	exsit, err := exists(logDir)
	if err != nil {
		return err
	}
	if exsit == false {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return err
		}
	}

	filename := fmt.Sprintf("%s/%s_%s.txt", logDir, app, time.Now().Format("2006-01-02"))

	if err := LogToFile(filename, v); err != nil {
		return err
	}

	return nil
}

func exists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
