package usecase

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/samber/do"
)

var (
	ErrCertificationPDFFound    = errors.New("Certification pdf found")
	ErrCertificationPDFNotFound = errors.New("Certification pdf not found")
)

func NewCertificationPDF(i *do.Injector) (*CertificationPDF, error) {
	c := &CertificationPDF{
		certificationPDFRepo: do.MustInvoke[CertificationPDFRepo](i),
		logger:               do.MustInvoke[*Logger](i),
	}

	return c, nil
}

type CertificationPDF struct {
	certificationPDFRepo CertificationPDFRepo
	logger               *Logger
}

// Create
// creates a new cert
func (u CertificationPDF) Create(ctx context.Context, buffer *bytes.Buffer) (uuid.UUID, error) {
	return u.certificationPDFRepo.Create(ctx, buffer, buffer.Len())
}

// GetByUUID
// returns the cert with the given UUID
func (u CertificationPDF) GetByUUID(ctx context.Context, uuid uuid.UUID) (io.Reader, error) {
	return u.certificationPDFRepo.Get(ctx, uuid)
}
