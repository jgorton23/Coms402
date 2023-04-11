package repo

import (
	"context"
	"io"
	"os"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

// Pattern to verify databaseEntImplemUser conforms to the required interfaces
var (
	_assertObjectStoreLocalImplemCertPDFStore                              = &objectStoreLocalImplemCertPDFStore{}
	_                                         usecase.CertificationPDFRepo = _assertObjectStoreLocalImplemCertPDFStore
)

func (implem *ObjectStoreLocal) PDFRepo() usecase.CertificationPDFRepo {
	bucketName := "pdf-repo"

	if _, err := os.Stat(implem.basePath + "/" + bucketName); os.IsNotExist(err) {
		os.MkdirAll(implem.basePath+"/"+bucketName, 0700)
	}

	return &objectStoreLocalImplemCertPDFStore{
		bucketName: bucketName,
		basePath:   implem.basePath,
	}
}

type objectStoreLocalImplemCertPDFStore struct {
	bucketName string
	basePath   string
}

func (implem *objectStoreLocalImplemCertPDFStore) Create(ctx context.Context, reader io.Reader, objectSize int) (uuid.UUID, error) {
	// Upload the zip file
	objectName := uuid.New()

	f, err := os.Create(implem.basePath + "/" + implem.bucketName + "/" + objectName.String() + ".pdf")

	if err != nil {
		return uuid.Nil, err
	}

	_, err = io.Copy(f, reader)

	f.Sync()

	if err != nil {
		return uuid.Nil, err
	}

	f.Close()

	return objectName, nil
}

func (implem objectStoreLocalImplemCertPDFStore) Get(ctx context.Context, uuid uuid.UUID) (io.Reader, error) {

	f, err := os.Open(implem.basePath + "/" + implem.bucketName + "/" + uuid.String() + ".pdf")

	if err != nil {
		return nil, err
	}

	return f, nil
}
