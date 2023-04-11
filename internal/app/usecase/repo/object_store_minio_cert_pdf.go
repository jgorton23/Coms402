package repo

import (
	"context"
	"io"
	"log"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

// Pattern to verify databaseEntImplemUser conforms to the required interfaces
var (
	_assertObjectStoreMinioImplemCertPDFStore                              = &objectStoreMinioImplemCertPDFStore{}
	_                                         usecase.CertificationPDFRepo = _assertObjectStoreMinioImplemCertPDFStore
)

func (implem *ObjectStoreMinio) PDFRepo(ctx context.Context) usecase.CertificationPDFRepo {
	bucketName := "pdf-repo"

	err := implem.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: implem.bucketLocation})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := implem.client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	return &objectStoreMinioImplemCertPDFStore{
		client:     implem.client,
		bucketName: bucketName,
	}
}

type objectStoreMinioImplemCertPDFStore struct {
	client     *minio.Client
	bucketName string
}

func (implem *objectStoreMinioImplemCertPDFStore) Create(ctx context.Context, reader io.Reader, objectSize int) (uuid.UUID, error) {
	// Upload the zip file
	objectName := uuid.New()

	contentType := "application/pdf"

	_, err := implem.client.PutObject(ctx, implem.bucketName, objectName.String(), reader, int64(objectSize), minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		return uuid.Nil, err
	}

	return objectName, nil

}

func (implem objectStoreMinioImplemCertPDFStore) Get(ctx context.Context, uuid uuid.UUID) (io.Reader, error) {
	reader, err := implem.client.GetObject(context.Background(), implem.bucketName, uuid.String(), minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return reader, nil
}
