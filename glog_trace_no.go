package glog

import (
	"context"
	"fmt"
)

type TraceNo string

var emptyTraceNo TraceNo

const (
	traceNo TraceNo = "_traceNo"
)

func WithTraceNo(c context.Context) context.Context {
	return context.WithValue(c, traceNo, TraceNo(RandomNumeric(6)))
}

func TContext(ctx context.Context) TraceNo {
	if v := ctx.Value(traceNo); v != nil {
		return v.(TraceNo)
	}
	return emptyTraceNo
}

func T(i interface{}) TraceNo {
	return TraceNo(fmt.Sprint(i))
}

func (t TraceNo) Info(args ...interface{}) {
	logging.print(infoLog, t, args...)
}

func (t TraceNo) Infoln(args ...interface{}) {
	logging.println(infoLog, t, args...)
}

func (t TraceNo) Infof(format string, args ...interface{}) {
	logging.printf(infoLog, t, format, args...)
}

func (t TraceNo) Warning(args ...interface{}) {
	logging.print(warningLog, t, args...)
}

func (t TraceNo) Warningln(args ...interface{}) {
	logging.println(warningLog, t, args...)
}

func (t TraceNo) Warningf(format string, args ...interface{}) {
	logging.printf(warningLog, t, format, args...)
}

func (t TraceNo) Error(args ...interface{}) {
	logging.print(errorLog, t, args...)
}

func (t TraceNo) Errorln(args ...interface{}) {
	logging.println(errorLog, t, args...)
}

func (t TraceNo) Errorf(format string, args ...interface{}) {
	logging.printf(errorLog, t, format, args...)
}

func (t TraceNo) Fatal(args ...interface{}) {
	logging.print(fatalLog, t, args...)
}

func (t TraceNo) Fatalln(args ...interface{}) {
	logging.println(fatalLog, t, args...)
}

func (t TraceNo) Fatalf(format string, args ...interface{}) {
	logging.printf(fatalLog, t, format, args...)
}
