package log

import (
	"context"
)

const (
	INDEX_LOG_ERROR    = "log_error"
	INDEX_LOG_ACTIVITY = "log_activity"
	INDEX_LOG_LOGIN    = "log_login"
)

func InsertErrorLog(ctx context.Context, log *LogError) error {
	return nil
	// return elasticsearch.Insert(ctx, INDEX_LOG_ERROR, log)
}

func InsertActivityLog(ctx context.Context, log *LogError) error {
	return nil
	// return elasticsearch.Insert(ctx, INDEX_LOG_ACTIVITY, log)
}

func InsertLoginLog(ctx context.Context, log *LogError) error {
	return nil
	// return elasticsearch.Insert(ctx, INDEX_LOG_LOGIN, log)
}
