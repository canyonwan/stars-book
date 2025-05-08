package v1

import "github.com/gogf/gf/v2/frame/g"

type ProficiencyLevel uint

const (
	ProficiencyLevel1 ProficiencyLevel = iota + 1
	ProficiencyLevel2
	ProficiencyLevel3
	ProficiencyLevel4
	ProficiencyLevel5
)

type CreateWordsReq struct {
	g.Meta             `path:"/words" method:"post" tags:"单词管理" summary:"新增单词"`
	Word               string           `json:"word" v:"required#单词不能为空|length:1,100#单词长度不能超过100" dc:"单词"`
	Definition         string           `json:"definition" v:"required#单词定义不能为空|length:1,1000#单词定义长度不能超过1000" dc:"单词定义"`
	ExampleSentence    string           `json:"exampleSentence" v:"length:1,1000#单词例句长度不能超过1000" dc:"单词例句"`
	ChineseTranslation string           `json:"chineseTranslation" v:"length:1,1000#中文翻译长度不能超过1000" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation" v:"length:1,100#发音长度不能超过100" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiencyLevel" v:"between:1,5#熟练度必须在1-5之间" dc:"熟练度，1最低，5最高"`
}
type CreateWordsRes struct{}
