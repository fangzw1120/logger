package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				msg: "print msg ok",
			},
		},
	}
	ifDebug := true
	ifSimpleLog := false
	verbose := true
	logPath := "/Users/forisfang/project/IOAProject/logger/logs/"
	Init(logPath, "logger.log", ifDebug, ifSimpleLog, verbose)
	InitStateLogger(logPath, "logger_state.log", ifDebug, ifSimpleLog, verbose)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.msg)
			Info(tt.args.msg)
			Errorf("dd: %+v", tt.args.msg)
			StatePrint(tt.args.msg)
		})
	}
}
