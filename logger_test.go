package logger

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestDebug ...
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
	logPath := "/Users/forisfang_mbp16/work/project/woa/logger/logs/"
	Init(
		logPath,
		"logger.log", ifDebug, ifSimpleLog, verbose, Params{RemovePathPrefix: "/Users/forisfang_mbp16/work/project/"})
	InitStateLogger(logPath, "logger_state.log", ifDebug, ifSimpleLog, verbose)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMaxAgeDay(100)
			Debug(tt.args.msg)
			Info(tt.args.msg)
			Errorf("dd: %+v", tt.args.msg)
			StatePrint(tt.args.msg)

			jsonBody := []byte(`{"client_message": "hello, server!"}`)
			bodyReader := bytes.NewReader(jsonBody)
			req, _ := http.NewRequest(http.MethodPost, "baidu.com", bodyReader)
			resp := &http.Response{
				Body: ioutil.NopCloser(bytes.NewBufferString("Hello World")),
			}
			DebugfHTTPReq("http: %+v", req, true)
			DebugfHTTPReq("http: %+q", req, false)
			DebugHTTPReq(req, true)
			InfofHTTPResp("resp: %+v", resp, true)

			ctx := context.Background()
			ctx = context.WithValue(ctx, struct{}{}, "testID")
			DebugfWithCtx(ctx, "%q", "12373")
		})
	}
}
