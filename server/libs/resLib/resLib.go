package resLib

type Rsp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}
