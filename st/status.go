package st

import (
	"google.golang.org/grpc/codes"
)

func NewError(code int32, msg string, gc codes.Code) *err {
	if gc == codes.OK {
		return nil
	}
	x := &body{
		Code: code,
		Msg:  msg,
	}

	err := &err{
		gst: gc,
		st:  x,
	}

	return err
}
