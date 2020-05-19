package st

import (
	"google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"
)

func NewErr(code int32, msg string, gc codes.Code) error {
	if gc == codes.OK {
		return nil
	}
	x := &body{
		Code: code,
		Msg:  msg,
	}

	str, _ := json.Marshal(x)
	s := &err{
		gst: gs.New(gc, string(str)),
		st:  x,
	}

	err := s.gst.Err()

	return err
}
