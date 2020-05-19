package st

import (
	gs "google.golang.org/grpc/status"
)

// err -
type err struct {
	gst *gs.Status
	// st  *body
	Code int32  `json:"code" yaml:"code" `
	Msg  string `json:"msg" yaml:"msg"`
}

// // body -
// type body struct {
	
// }

