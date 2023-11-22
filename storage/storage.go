package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/txsvc/cloudlib"
)

const (
	TypeStorage cloudlib.ProviderType = 20
)

type (
	StorageProvider interface {
		Bucket(string) BucketHandle
	}

	BucketHandle interface {
		Object(string) ObjectHandle
	}

	ObjectHandle interface {
		Delete() error
		NewReader(context.Context) (io.Reader, error)
		NewWriter(context.Context) (io.Writer, error)
		Close() error
	}
)

var (
	storageProvider *cloudlib.Provider
)

func NewConfig(opts cloudlib.ProviderConfig) (*cloudlib.Provider, error) {
	if opts.Type != TypeStorage {
		return nil, fmt.Errorf(cloudlib.MsgUnsupportedProviderType, opts.Type)
	}

	o, err := cloudlib.New(opts)
	if err != nil {
		return nil, err
	}
	storageProvider = o

	return o, nil
}

func UpdateConfig(opts cloudlib.ProviderConfig) (*cloudlib.Provider, error) {
	if opts.Type != TypeStorage {
		return nil, fmt.Errorf(cloudlib.MsgUnsupportedProviderType, opts.Type)
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
