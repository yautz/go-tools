package st

import (
	"google.golang.org/grpc/codes"
)

// err -
type err struct {
	gst codes.Code // grpc Code
	st  *body      //自定義 Code
}

// body -
type body struct {
	Code int32  `json:"code" yaml:"code" `
	Msg  string `json:"msg" yaml:"msg"`
}
