package resLib

type RspTest struct {
	Code uint32 `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

type Rsp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
