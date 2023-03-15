package slack

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-logr/logr"
	"github.com/hashicorp/go-retryablehttp"
)

var (
	retryMaxStr = os.Getenv("SLACK_API_RETRY_MAX")
	waitMaxStr  = os.Getenv("SLACK_API_RETRY_WAIT_MAX")
	waitMinStr  = os.Getenv("SLACK_API_RETRY_WAIT_MIN")
	retryMax    int
	waitMax     time.Duration
	waitMin     time.Duration
)

func init() {
	if retryMaxStr == "" {
		retryMax = 1000
	} else {
		var err error
		retryMax, err = strconv.Atoi(retryMaxStr)
		if err != nil {
			panic(err)
		}
	}
	if waitMaxStr == "" {
		waitMax = 1 * time.Minute
	} else {
		var err error
		waitMax, err = time.ParseDuration(waitMaxStr)
		if err != nil {
			panic(err)
		}
	}
	if waitMinStr == "" {
		waitMin = 1 * time.Second
	} else {
		var err error
		waitMin, err = time.ParseDuration(waitMinStr)
		if err != nil {
			panic(err)
		}
	}
}

type TransportLogger struct {
	logger logr.Logger
}

func (t TransportLogger) Error(msg string, keysAndValues ...interface{}) {
	t.logger.Error(nil, msg, keysAndValues...)
}

func (t TransportLogger) Info(msg string, keysAndValues ...interface{}) {
	t.logger.Info(msg, keysAndValues...)
}

func (t TransportLogger) Debug(msg string, keysAndValues ...interface{}) {
	t.logger.V(1).Info(msg, keysAndValues...)
}

func (t TransportLogger) Warn(msg string, keysAndValues ...interface{}) {
	t.logger.V(1).Info(msg, keysAndValues...)
}

// NewThrottledTransport wraps transportWrap with a rate limitter
func NewThrottledTransport(logger logr.Logger) *http.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = retryMax
	retryClient.Logger = TransportLogger{logger: logger}
	retryClient.RetryWaitMax = waitMax
	retryClient.RetryWaitMin = waitMin
	// Default retry policy handles 429 and more
	retryClient.CheckRetry = retryablehttp.DefaultRetryPolicy
	// Default backoff uses `Retry-After` header if present
	retryClient.Backoff = retryablehttp.DefaultBackoff

	return retryClient.StandardClient()
}
