package err

import (
	"asyncflow/pkg/constant"
	"fmt"
	"github.com/pkg/errors"
)

type BizCode = int32

const (
	BizCodeInvalidInputErr        BizCode = 9001
	BizCodeTaskCfgNotFoundErr     BizCode = 9002
	BizCodeSchedulePosNotFoundErr BizCode = 9003
	BizCodeCreateTaskErr          BizCode = 9004
	BizCodeSnowflakeErr           BizCode = 9101
	BizCodeInternalErr            BizCode = 9999
)

func MapToRspCode(code BizCode) constant.RspCode {
	switch code {
	case BizCodeInvalidInputErr:
		return constant.RspCodeInvalidInputErr
	case BizCodeTaskCfgNotFoundErr:
		return constant.RspCodeTaskCfgNotFoundErr
	case BizCodeSchedulePosNotFoundErr:
		return constant.RspCodeSchedulePosNotFoundErr
	case BizCodeCreateTaskErr:
		return constant.RspCodeCreateTaskErr
	case BizCodeSnowflakeErr:
		return constant.RspCodeSnowflakeErr
	case BizCodeInternalErr:
		return constant.RspCodeInternalErr
	default:
		panic("unknown BizCode: " + string(code))
	}
}

type BizError struct {
	Code    BizCode
	Message string
	cause   error
}

func (e *BizError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.cause != nil {
		return e.cause.Error()
	}
	return fmt.Sprintf("BizError Code: %d", e.Code)
}

func (e *BizError) Unwrap() error {
	return e.cause
}

func NewBizError(code BizCode, msg string) *BizError {
	return &BizError{
		Code:    code,
		Message: msg,
		cause:   errors.WithStack(errors.New("BizError: " + string(code))),
	}
}

func NewBizErrorWithCause(code BizCode, msg string, cause error) *BizError {
	return &BizError{
		Code:    code,
		Message: msg,
		cause:   errors.WithStack(cause),
	}
}

func WrapBizError(err error, msg string) *BizError {
	if err == nil {
		return nil
	}
	var bizErr BizError
	if ok := errors.As(err, &bizErr); ok {
		m := msg
		if m == "" {
			m = bizErr.Message
		}
		return &BizError{
			Code:    bizErr.Code,
			Message: m,
			cause:   errors.WithStack(&bizErr),
		}
	}
	panic(fmt.Sprintf("WrapBizError: err is not BizError, got %T", err))
}
