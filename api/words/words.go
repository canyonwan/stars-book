// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package words

import (
	"context"

	"gf-playground-1/api/words/v1"
)

type IWordsV1 interface {
	CreateWords(ctx context.Context, req *v1.CreateWordsReq) (res *v1.CreateWordsRes, err error)
}
