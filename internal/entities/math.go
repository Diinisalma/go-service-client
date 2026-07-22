package entities

type NumbersReq struct {
	A int32 `json:"a"`
	B int32	`json:"b"`
}

type NumbersResp struct {
	Error string `json:"error,omitempty"`
	Result int32 `json:"result,omitempty"`
}