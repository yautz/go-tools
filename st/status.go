package st

import (
	"google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"
)

func NewError(code int32, msg string, gc codes.Code) error {
	if gc == codes.OK {
		return nil
	}
	// x := &body{
	// 	Code: code,
	// 	Msg:  msg,
	// }

	// str, _ := json.Marshal(x)
	// s := &err{
	// 	gst: gs.New(gc, string(str)),
	// 	st:  x,
	// }

	// err := s.gst.Err()

	x := &err{
		gst:  gs.New(gc, msg),
		Code: code,
		Msg:  msg,
	}
	err := x.gst.Err()
	return err
}

// // ConvertError -
// func ConvertError(errs error) Errors {
// 	if errs == nil {
// 		return &err{
// 			gst: gs.New(OK, "success"),
// 			st:  &body{Code: 0, Msg: "success"},
// 		}
// 	}

// 	gt, ok := gs.FromError(errs)

// 	temp := &err{
// 		gst: gt,
// 	}

// 	if ok {
// 		// 如果不能解析則給他錯誤代碼
// 		if err := temp.parseMsg(); err != nil {
// 			temp.st.Code = 99999
// 			temp.st.Msg = gt.Message()
// 		}
// 	} else {
// 		fmt.Printf("convert error failed with errr : %s", errs.Error())
// 	}

// 	return temp
// }
