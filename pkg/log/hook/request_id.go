package hook

import "github.com/sirupsen/logrus"

const RequestIdFieldName = "request-id"

type RequestIdHook struct{}

func (hook *RequestIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *RequestIdHook) Fire(entry *logrus.Entry) error {
	if entry.Context != nil {
		requestID, ok := entry.Context.Value(RequestIdFieldName).(string)
		if ok {
			entry.Data[RequestIdFieldName] = requestID
		}
	}

	return nil
}
