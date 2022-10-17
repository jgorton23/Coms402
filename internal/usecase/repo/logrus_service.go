package repo

import (
	"os"

	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	logrusadapter "logur.dev/adapter/logrus"
)

func NewLogrusService(config entity.LOG) func(i *do.Injector) (logrusadapter.Logger, error) {
	return func(i *do.Injector) (logrusadapter.Logger, error) {
		l := logrus.New()

		l.SetOutput(os.Stdout)

		switch config.Format {
		case "json":
			l.SetFormatter(&logrus.JSONFormatter{})
		case "logfmt":
			l.SetFormatter(&logrus.TextFormatter{
				DisableColors:             config.NoColor,
				EnvironmentOverrideColors: true,
			})
		}
		level, err := logrus.ParseLevel(config.Level)

		if err != nil {
			return logrusadapter.Logger{}, err
		}

		l.SetLevel(level)

		return *logrusadapter.New(l), nil
	}
}
