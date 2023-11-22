package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/txsvc/cloudlib/provider"
)

const (
	TypeStorage provider.ProviderType = 20
)

type (
	StorageProvider interface {
		Bucket(string) BucketHandle
	}

	BucketHandle interface {
		Object(string) ObjectHandle
	}

	ObjectHandle interface {
		Close() error
		Delete() error
		NewReader(context.Context) (io.Reader, error)
		NewWriter(context.Context) (io.Writer, error)
	}
)

var (
	storageProvider *provider.Provider
)

func NewConfig(opts provider.ProviderConfig) (*provider.Provider, error) {
	if opts.Type != TypeStorage {
		return nil, fmt.Errorf(provider.MsgUnsupportedProviderType, opts.Type)
	}

	o, err := provider.New(opts)
	if err != nil {
		return nil, err
	}
	storageProvider = o

	return o, nil
}

func UpdateConfig(opts provider.ProviderConfig) (*provider.Provider, error) {
	if opts.Type != TypeStorage {
		return nil, fmt.Errorf(provider.MsgUnsupportedProviderType, opts.Type)
	}

	return storageProvider, storageProvider.RegisterProviders(true, opts)
}

func Bucket(name string) BucketHandle {
	imp, found := storageProvider.Find(TypeStorage)
	if !found {
		return nil
	}
	return imp.(StorageProvider).Bucket(name)
}
