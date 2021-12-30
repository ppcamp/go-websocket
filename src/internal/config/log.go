package config

import "github.com/sirupsen/logrus"

func SetupLoggers() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: App.LogPrettyPrint})
	level, err := logrus.ParseLevel(App.LogLevel)

	if err != nil {
		logrus.WithError(err).Fatal("parsing log level")
	}
	logrus.SetLevel(level)
	logrus.WithField("AppConfig", App).Info("Environment variables")
}
