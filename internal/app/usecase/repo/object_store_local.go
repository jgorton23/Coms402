package repo

// objectStoreLocalBuilder creates a new objectStoreLocalBuilder
func ObjectStoreLocalBuilder() *objectStoreLocalBuilder {
	return &objectStoreLocalBuilder{}
}

type objectStoreLocalBuilder struct {
	basePath string
}

func (builder *objectStoreLocalBuilder) WithBasePath(path string) *objectStoreLocalBuilder {
	builder.basePath = path
	return builder
}

func (builder *objectStoreLocalBuilder) New() (*ObjectStoreLocal, error) {
	return &ObjectStoreLocal{
		basePath: builder.basePath,
	}, nil
}

type ObjectStoreLocal struct {
	basePath string
}
