package requests

import "math/big"

type EmailCode struct {
	Code *big.Int `json:"code" form:"code"`
}
