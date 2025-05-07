// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure for table words.
type Words struct {
	Id                 string      `json:"id"                 orm:"id"                  description:""`                 //
	Uid                string      `json:"uid"                orm:"uid"                 description:"用户id，标记该单词所属用户"`   // 用户id，标记该单词所属用户
	Word               string      `json:"word"               orm:"word"                description:"单词"`               // 单词
	Definition         string      `json:"definition"         orm:"definition"          description:"单词定义"`             // 单词定义
	ExampleSentence    string      `json:"exampleSentence"    orm:"example_sentence"    description:"示例句子"`             // 示例句子
	ChineseTranslation string      `json:"chineseTranslation" orm:"chinese_translation" description:"单词的中文翻译"`          // 单词的中文翻译
	Pronunciation      string      `json:"pronunciation"      orm:"pronunciation"       description:"单词的发音"`            // 单词的发音
	ProficiencyLevel   uint        `json:"proficiencyLevel"   orm:"proficiency_level"   description:"单词掌握程度，1级最低，5级最高"` // 单词掌握程度，1级最低，5级最高
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"          description:""`                 //
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"          description:""`                 //
	DeletedAt          *gtime.Time `json:"deletedAt"          orm:"deleted_at"          description:""`                 //
}
