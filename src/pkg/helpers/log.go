package helpers

import "github.com/sirupsen/logrus"

func NewModuleLogger(module string) *logrus.Entry { return logrus.WithField("module", module) }
