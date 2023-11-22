package observer

import (
	"context"
	"fmt"

	"github.com/txsvc/cloudlib"
)

const (
	DefaultLogId = "default"
	MetricsLogId = "metric"
	ValuesLogId  = "values"

	LevelDebug Severity = iota
	LevelInfo
	LevelNotice
	LevelWarn
	LevelError
	LevelAlert

	TypeLogger        cloudlib.ProviderType = 10
	TypeErrorReporter cloudlib.ProviderType = 11
	TypeMetrics       cloudlib.ProviderType = 12
)

type (
	Severity int

	ErrorReportingProvider interface {
		ReportError(error) error
	}

	MetricsProvider interface {
		Meter(ctx context.Context, metric string, vals ...string)
	}

	LoggingProvider interface {
		Log(string, ...string)
		LogWithLevel(Severity, string, ...string)

		EnableLogging()
		DisableLogging()
	}
)

var (
	observerProvider *cloudlib.Provider
)

func Instance() *cloudlib.Provider {
	return observerProvider
}

func NewConfig(opts ...cloudlib.ProviderConfig) (*cloudlib.Provider, error) {
	if pc := validateProviders(opts...); pc != nil {
		return nil, fmt.Errorf(cloudlib.MsgUnsupportedProviderType, pc.Type)
	}

	o, err := cloudlib.New(opts...)
	if err != nil {
		return nil, err
	}
	observerProvider = o

	return o, nil
}

func UpdateConfig(opts ...cloudlib.ProviderConfig) (*cloudlib.Provider, error) {
	if pc := validateProviders(opts...); pc != nil {
		return nil, fmt.Errorf(cloudlib.MsgUnsupportedProviderType, pc.Type)
	}

	return observerProvider, Instance().RegisterProviders(true, opts...)
}

func validateProviders(opts ...cloudlib.ProviderConfig) *cloudlib.ProviderConfig {
	for _, pc := range opts {
		if pc.Type != TypeLogger && pc.Type != TypeErrorReporter && pc.Type != TypeMetrics {
			return &pc // this is not one of the above i.e. not supported
		}
	}
	return nil
}

func Log(msg string, keyValuePairs ...string) {
	imp, found := Instance().Find(TypeLogger)
	if !found {
		return
	}
	imp.(LoggingProvider).Log(msg, keyValuePairs...)
}

func LogWithLevel(lvl Severity, msg string, keyValuePairs ...string) {
	imp, found := Instance().Find(TypeLogger)
	if !found {
		return
	}
	imp.(LoggingProvider).LogWithLevel(lvl, msg, keyValuePairs...)
}

func EnableLogging() {
	imp, found := Instance().Find(TypeLogger)
	if !found {
		return
	}
	imp.(LoggingProvider).EnableLogging()
}

func DisableLogging() {
	imp, found := Instance().Find(TypeLogger)
	if !found {
		return
	}
	imp.(LoggingProvider).DisableLogging()
}

func Meter(ctx context.Context, metric string, vals ...string) {
	imp, found := Instance().Find(TypeMetrics)
	if !found {
		return
	}
	imp.(MetricsProvider).Meter(ctx, metric, vals...)
}

func ReportError(e error) error {
	imp, found := Instance().Find(TypeErrorReporter)
	if !found {
		return nil
	}
	return imp.(ErrorReportingProvider).ReportError(e)
}
