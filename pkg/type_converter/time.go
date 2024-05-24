package type_converter

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func protoTimeToSqlTime(sqlTime sql.NullTime) *timestamppb.Timestamp {
	if sqlTime.Valid {
		return timestamppb.New(sqlTime.Time)
	}

	return nil
}

func sqlTimeToProtoTime(protoTimestamp *timestamppb.Timestamp) sql.NullTime {
	if protoTimestamp != nil {
		return sql.NullTime{Time: protoTimestamp.AsTime(), Valid: true}
	}

	return sql.NullTime{Valid: false}
}
