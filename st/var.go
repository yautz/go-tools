package st

// 0	    success
// 11000	warning
// 12000	other warning(by project)
// 21000	error
// 22000	other error(by project)

var (
	NoError = NewErr(0, "success", OK)

	//
	WarningLoginFail = NewErr(11000, "loging failed", NotFound)
)
