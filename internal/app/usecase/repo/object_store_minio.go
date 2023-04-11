package repo

import (
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/samber/do"
)

// Pattern to verify objectStoreMinioImplem conforms to the required interfaces
var (
	_assertObjectStoreMinioImplem                    = &ObjectStoreMinio{}
	_                             do.Healthcheckable = _assertObjectStoreMinioImplem
)

// objectStoreMinioBuilder creates a new objectStoreMinioBuilder
func ObjectStoreMinioBuilder() *objectStoreMinioBuilder {
	return &objectStoreMinioBuilder{}
}

type objectStoreMinioBuilder struct {
	creds    *credentials.Credentials
	endpoint string
	useSSL   bool
}

func (builder *objectStoreMinioBuilder) WithEndpoint(endpoint string) *objectStoreMinioBuilder {
	builder.endpoint = endpoint
	return builder
}

func (builder *objectStoreMinioBuilder) WithCreds(creds *credentials.Credentials) *objectStoreMinioBuilder {
	builder.creds = creds
	return builder
}

func (builder *objectStoreMinioBuilder) WithSSL(ssl bool) *objectStoreMinioBuilder {
	builder.useSSL = ssl
	return builder
}

func (builder *objectStoreMinioBuilder) New() (*ObjectStoreMinio, error) {
	// Initialize minio client object.
	minioClient, err := minio.New(builder.endpoint, &minio.Options{
		Creds:  builder.creds,
		Secure: builder.useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &ObjectStoreMinio{
		client:         minioClient,
		bucketLocation: "us-east-1",
	}, nil
}

type ObjectStoreMinio struct {
	client         *minio.Client
	bucketLocation string
}

func (osm *ObjectStoreMinio) HealthCheck() error {
	_, err := osm.client.HealthCheck(10 * time.Second)

	return err
}
