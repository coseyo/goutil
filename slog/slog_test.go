package slog

import "testing"

func Test_SimpleLog(t *testing.T) {
	err := SimpleLog("filename", "logtext1", "logtext2")
	if err != nil {
		t.Error(err)
	}
}
