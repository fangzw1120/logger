package logger

import "context"

func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, struct{}{}, traceID)
}

func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value(struct{}{}).(string)
	if !ok {
		return ""
	}
	return traceID
}

// DebugfWithCtx ...
func DebugfWithCtx(ctx context.Context, format string, v ...interface{}) {
	traceID := GetTraceID(ctx)
	format = FILE() + ":" + LINE() + ":[" + traceID + "]" + " " + format
	logger.Debugf(format, v...)
}

// InfofWithCtx ...
func InfofWithCtx(ctx context.Context, format string, v ...interface{}) {
	traceID := GetTraceID(ctx)
	format = FILE() + ":" + LINE() + ":[" + traceID + "]" + " " + format
	logger.Infof(format, v...)
}

// WarnfWithCtx ...
func WarnfWithCtx(ctx context.Context, format string, v ...interface{}) {
	traceID := GetTraceID(ctx)
	format = FILE() + ":" + LINE() + ":[" + traceID + "]" + " " + format
	logger.Warnf(format, v...)
}

// ErrorfWithCtx ...
func ErrorfWithCtx(ctx context.Context, format string, v ...interface{}) {
	traceID := GetTraceID(ctx)
	format = FILE() + ":" + LINE() + ":[" + traceID + "]" + " " + format
	logger.Errorf(format, v...)
}

// FatalfWithCtx ...
func FatalfWithCtx(ctx context.Context, format string, v ...interface{}) {
	traceID := GetTraceID(ctx)
	format = FILE() + ":" + LINE() + ":[" + traceID + "]" + " " + format
	logger.Fatalf(format, v...)
}
