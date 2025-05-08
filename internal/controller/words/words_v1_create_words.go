package words

import (
	"context"

	"gf-playground-1/api/words/v1"
	"gf-playground-1/internal/model/entity"
)

func (c *ControllerV1) CreateWords(ctx context.Context, req *v1.CreateWordsReq) (res *v1.CreateWordsRes, err error) {
	err = c.words.Create(ctx, entity.Words{
		Word:               req.Word,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ChineseTranslation: req.ChineseTranslation,
		Pronunciation:      req.Pronunciation,
		ProficiencyLevel:   uint(req.ProficiencyLevel),
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
