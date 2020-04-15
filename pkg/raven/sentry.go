package raven

import (
	"github.com/getsentry/sentry-go"
	"time"
)

func ReportIfError(err error) {

	sentry.CaptureException(err)
	sentry.Flush(time.Second * 5)

}
