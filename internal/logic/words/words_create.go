package words

import (
	"context"

	v1 "gf-playground-1/api/words/v1"
	"gf-playground-1/internal/dao"
	"gf-playground-1/internal/model/do"
	"gf-playground-1/internal/model/entity"
	"gf-playground-1/utility"

	"github.com/gogf/gf/v2/errors/gerror"
)

type CreateWordInput struct {
	Uid                string
	Word               string
	Definition         string
	ExampleSentence    string
	ChineseTranslation string
	Pronunciation      string
	ProficiencyLevel   v1.ProficiencyLevel
}

func (w *Words) Create(ctx context.Context, in entity.Words) error {
	var cls = dao.Words.Columns()
	model := dao.Words.Ctx(ctx)

	count, err := model.
		Where(cls.Uid, in.Uid).
		Where(cls.Word, in.Word).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}

	uid := ctx.Value("userId").(string)

	_, err = model.Data(do.Words{
		Id:                 utility.GenerateSnowId(),
		Uid:                uid,
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}
