package logger

import (
	"net/http"
	"net/http/httputil"
)

// DebugfHTTPReq ...
//  @Description:
//  @param format
//  @param req
//  @param ifPrintBody
//
func DebugfHTTPReq(format string, req *http.Request, ifPrintBody bool) {
	format = FILE() + ":" + LINE() + " " + format
	reqDump, _ := httputil.DumpRequest(req, ifPrintBody)
	logger.Debugf(format, string(reqDump))
}

// DebugHTTPReq ...
//  @Description:
//  @param req
//  @param ifPrintBody
//
func DebugHTTPReq(req *http.Request, ifPrintBody bool) {
	reqDump, _ := httputil.DumpRequest(req, ifPrintBody)
	msg := FILE() + ":" + LINE() + " " + string(reqDump)
	logger.Debug(msg)
}

// InfofHTTPReq ...
//  @Description:
//  @param format
//  @param req
//  @param ifPrintBody
//
func InfofHTTPReq(format string, req *http.Request, ifPrintBody bool) {
	format = FILE() + ":" + LINE() + " " + format
	reqDump, _ := httputil.DumpRequest(req, ifPrintBody)
	logger.Infof(format, string(reqDump))
}

// InfoHTTPReq ...
//  @Description:
//  @param req
//  @param ifPrintBody
//
func InfoHTTPReq(req *http.Request, ifPrintBody bool) {
	reqDump, _ := httputil.DumpRequest(req, ifPrintBody)
	msg := FILE() + ":" + LINE() + " " + string(reqDump)
	logger.Info(msg)
}

// WarnfHTTPReq ...
//  @Description:
//  @param format
//  @param req
//  @param ifPrintBody
//
func WarnfHTTPReq(format string, req *http.Request, ifPrintBody bool) {
	format = FILE() + ":" + LINE() + " " + format
	reqDump, _ := httputil.DumpRequest(req, ifPrintBody)
	logger.Warnf(format, string(reqDump))
}

// WarnHTTPReq ...
//  @Description:
//  @param req
//  @param ifPrintBody
//
func WarnHTTPReq(req *http.Request, ifPrintBody bool) {
	reqDump, _ := httputil.DumpRequest(req, ifPrintBody)
	msg := FILE() + ":" + LINE() + " " + string(reqDump)
	logger.Warn(msg)
}

// ErrorfHTTPReq ...
//  @Description:
//  @param format
//  @param req
//  @param ifPrintBody
//
func ErrorfHTTPReq(format string, req *http.Request, ifPrintBody bool) {
	format = FILE() + ":" + LINE() + " " + format
	reqDump, _ := httputil.DumpRequest(req, ifPrintBody)
	logger.Errorf(format, string(reqDump))
}

// ErrorHTTPReq ...
//  @Description:
//  @param req
//  @param ifPrintBody
//
func ErrorHTTPReq(req *http.Request, ifPrintBody bool) {
	reqDump, _ := httputil.DumpRequest(req, ifPrintBody)
	msg := FILE() + ":" + LINE() + " " + string(reqDump)
	logger.Error(msg)
}

// FatalfHTTPReq ...
//  @Description:
//  @param format
//  @param req
//  @param ifPrintBody
//
func FatalfHTTPReq(format string, req *http.Request, ifPrintBody bool) {
	format = FILE() + ":" + LINE() + " " + format
	reqDump, _ := httputil.DumpRequest(req, ifPrintBody)
	logger.Fatalf(format, string(reqDump))
}

// FatalHTTPReq ...
//  @Description:
//  @param req
//  @param ifPrintBody
//
func FatalHTTPReq(req *http.Request, ifPrintBody bool) {
	reqDump, _ := httputil.DumpRequest(req, ifPrintBody)
	msg := FILE() + ":" + LINE() + " " + string(reqDump)
	logger.Fatal(msg)
}

// DebugfHTTPResp ...
//  @Description:
//  @param format
//  @param resp
//  @param ifPrintBody
//
func DebugfHTTPResp(format string, resp *http.Response, ifPrintBody bool) {
	format = FILE() + ":" + LINE() + " " + format
	reqDump, _ := httputil.DumpResponse(resp, ifPrintBody)
	logger.Debugf(format, string(reqDump))
}

// DebugHTTPResp ...
//  @Description:
//  @param resp
//  @param ifPrintBody
//
func DebugHTTPResp(resp *http.Response, ifPrintBody bool) {
	reqDump, _ := httputil.DumpResponse(resp, ifPrintBody)
	msg := FILE() + ":" + LINE() + " " + string(reqDump)
	logger.Debug(msg)
}

// InfofHTTPResp ...
//  @Description:
//  @param format
//  @param resp
//  @param ifPrintBody
//
func InfofHTTPResp(format string, resp *http.Response, ifPrintBody bool) {
	format = FILE() + ":" + LINE() + " " + format
	reqDump, _ := httputil.DumpResponse(resp, ifPrintBody)
	logger.Infof(format, string(reqDump))
}

// InfoHTTPResp ...
//  @Description:
//  @param resp
//  @param ifPrintBody
//
func InfoHTTPResp(resp *http.Response, ifPrintBody bool) {
	reqDump, _ := httputil.DumpResponse(resp, ifPrintBody)
	msg := FILE() + ":" + LINE() + " " + string(reqDump)
	logger.Info(msg)
}

// WarnfHTTPResp ...
//  @Description:
//  @param format
//  @param resp
//  @param ifPrintBody
//
func WarnfHTTPResp(format string, resp *http.Response, ifPrintBody bool) {
	format = FILE() + ":" + LINE() + " " + format
	reqDump, _ := httputil.DumpResponse(resp, ifPrintBody)
	logger.Warnf(format, string(reqDump))
}

// WarnHTTPResp ...
//  @Description:
//  @param resp
//  @param ifPrintBody
//
func WarnHTTPResp(resp *http.Response, ifPrintBody bool) {
	reqDump, _ := httputil.DumpResponse(resp, ifPrintBody)
	msg := FILE() + ":" + LINE() + " " + string(reqDump)
	logger.Warn(msg)
}

// ErrorfHTTPResp ...
//  @Description:
//  @param format
//  @param resp
//  @param ifPrintBody
//
func ErrorfHTTPResp(format string, resp *http.Response, ifPrintBody bool) {
	format = FILE() + ":" + LINE() + " " + format
	reqDump, _ := httputil.DumpResponse(resp, ifPrintBody)
	logger.Errorf(format, string(reqDump))
}

// ErrorHTTPResp ...
//  @Description:
//  @param resp
//  @param ifPrintBody
//
func ErrorHTTPResp(resp *http.Response, ifPrintBody bool) {
	reqDump, _ := httputil.DumpResponse(resp, ifPrintBody)
	msg := FILE() + ":" + LINE() + " " + string(reqDump)
	logger.Error(msg)
}

// FatalfHTTPResp ...
//  @Description:
//  @param format
//  @param resp
//  @param ifPrintBody
//
func FatalfHTTPResp(format string, resp *http.Response, ifPrintBody bool) {
	format = FILE() + ":" + LINE() + " " + format
	reqDump, _ := httputil.DumpResponse(resp, ifPrintBody)
	logger.Fatalf(format, string(reqDump))
}

// FatalHTTPResp ...
//  @Description:
//  @param resp
//  @param ifPrintBody
//
func FatalHTTPResp(resp *http.Response, ifPrintBody bool) {
	reqDump, _ := httputil.DumpResponse(resp, ifPrintBody)
	msg := FILE() + ":" + LINE() + " " + string(reqDump)
	logger.Fatal(msg)
}
