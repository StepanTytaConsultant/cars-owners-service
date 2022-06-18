package ctx

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/cars-owners-service/internal/data/store"
)

const (
	ctxLog      = "ctxLog"
	ctxProvider = "ctxProvider"
)

func CtxLog(log *logrus.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxLog, log)
	}
}

func Log(r *http.Request) *logrus.Entry {
	return r.Context().Value(ctxLog).(*logrus.Entry)
}

func CtxProvider(provider store.DataProvider) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxProvider, provider)
	}
}

func Provider(r *http.Request) store.DataProvider {
	return r.Context().Value(ctxProvider).(store.DataProvider)
}
