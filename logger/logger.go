package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	logrus "github.com/sirupsen/logrus"
)

// Log - logrus logging.
//
//revive:disable:import-shadowing
var Log *logrus.Logger

// SetupLog - return logrus.
func init() {
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{
		TimestampFormat: "Jan 02 03:04:05.000 PM",
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	})

	Log = log
}
