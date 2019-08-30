package raven

import (
	"github.com/getsentry/sentry-go"
	"time"
)

func ReportIfError(err error) {
	if err != nil {
		sentry.CaptureException(err)
		sentry.Flush(time.Second * 5)
	}
}
