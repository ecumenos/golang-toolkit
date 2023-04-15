package customerror

import "go.uber.org/zap"

// Log log error according to error class.
func Log(log *zap.SugaredLogger, err error) {
	if err == nil {
		return
	}
	ce := CastStrict(err)
	if ce == nil {
		log.Error("something went wrong", zap.Error(err))
		return
	}

	switch ce.Class {
	case ErrorClass:
		logError(log, ce)
	case FailureClass:
		logFailure(log, ce)
	default:
		log.Error("unknown class error", zap.Error(err))
	}
}

func logError(log *zap.SugaredLogger, ce *CustomError) {
	if ce == nil {
		return
	}
	log.Error(ce.Err.Message, zap.Uint32("error_code", ce.Err.Code.UInt32()))
}

func logFailure(log *zap.SugaredLogger, ce *CustomError) {
	if ce == nil {
		return
	}
	fields := []interface{}{
		ce.Fail.Message,
		zap.Uint32("failure_code", ce.Fail.Code.UInt32()),
		zap.String("description", ce.Fail.Description),
		zap.Int("status_code", ce.Fail.StatusCode),
	}
	if ce.Fail.Data != nil {
		fields = append(fields, zap.Any("data", ce.Fail.Data))
	}
	log.Error(fields...)
}
