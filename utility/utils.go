package utility

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/crypto/gmd5"
)

// GenerateSnowId 生成雪花id
func GenerateSnowId() snowflake.ID {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	id := node.Generate()
	return id
}
func EncryptPassword(password string) string {
	return gmd5.MustEncryptString(password)
}
