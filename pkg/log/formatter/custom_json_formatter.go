package formatter

import (
	"github.com/sirupsen/logrus"
	"time"
)

// CustomJSONFormatter is a custom JSON formatter with timestamp and extra newline.
type CustomJSONFormatter struct {
	logrus.JSONFormatter
}

// Format formats the log entry to JSON with timestamp and extra newline.
func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Time = entry.Time.In(time.FixedZone("Europe/Moscow", 3*60*60)) // 3*60*60 - это смещение в секундах

	data, err := f.JSONFormatter.Format(entry)
	if err != nil {
		return nil, err
	}

	// Добавление метки времени в начало строки и дополнительной пустой строки
	formattedLog := append([]byte(entry.Time.Format("[2006-01-02 15:04:05]")+" "), data...)
	formattedLog = append(formattedLog, '\n')

	return formattedLog, nil
}
