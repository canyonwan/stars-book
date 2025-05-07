// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure of table words for DAO operations like Where/Data.
type Words struct {
	g.Meta             `orm:"table:words, do:true"`
	Id                 interface{} //
	Uid                interface{} // 用户id，标记该单词所属用户
	Word               interface{} // 单词
	Definition         interface{} // 单词定义
	ExampleSentence    interface{} // 示例句子
	ChineseTranslation interface{} // 单词的中文翻译
	Pronunciation      interface{} // 单词的发音
	ProficiencyLevel   interface{} // 单词掌握程度，1级最低，5级最高
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
	DeletedAt          *gtime.Time //
}
