package cron

import (
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type logger struct {
	logInfo bool
}

// Info logs routine messages about cron's operation.
func (l logger) Info(msg string, keysAndValues ...interface{}) {
	if l.logInfo {
		keysAndValues = formatTimes(keysAndValues)
		logx.Infof(formatString(len(keysAndValues)),
			append([]interface{}{msg}, keysAndValues...)...)
	}
}

// Error logs an error condition.
func (l logger) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = formatTimes(keysAndValues)
	logx.Infof(
		formatString(len(keysAndValues)+2),
		append([]interface{}{msg, "error", err}, keysAndValues...)...)
}

// formatString returns a logfmt-like format string for the number of
// key/values.
func formatString(numKeysAndValues int) string {
	var sb strings.Builder
	sb.WriteString("%s")
	if numKeysAndValues > 0 {
		sb.WriteString(", ")
	}
	for i := 0; i < numKeysAndValues/2; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("%v=%v")
	}
	return sb.String()
}

// formatTimes formats any time.Time values as RFC3339.
func formatTimes(keysAndValues []interface{}) []interface{} {
	var formattedArgs []interface{}
	for _, arg := range keysAndValues {
		if t, ok := arg.(time.Time); ok {
			arg = t.Format(time.RFC3339)
		}
		formattedArgs = append(formattedArgs, arg)
	}
	return formattedArgs
}
