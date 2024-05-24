package log

import (
	log_formatter "github.com/Freeline95/go-common/pkg/log/formatter"
	log_hook "github.com/Freeline95/go-common/pkg/log/hook"
	"github.com/sirupsen/logrus"
)

func InitLogger() {
	logrus.SetFormatter(&log_formatter.CustomJSONFormatter{})
	logrus.AddHook(&log_hook.RequestIdHook{})
}
