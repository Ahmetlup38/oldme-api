package other

import (
	"context"

	"github.com/oldme-git/oldme-api/api/other/app"
	"github.com/oldme-git/oldme-api/internal/logic/sentence"
	"github.com/oldme-git/oldme-api/internal/logic/word"
	"github.com/oldme-git/oldme-api/utility/uinit"
)

func (c *ControllerApp) Day(ctx context.Context, req *app.DayReq) (res *app.DayRes, err error) {
	poem, err := sentence.Text(ctx, uinit.PoemTagId)
	if err != nil {
		return nil, err
	}

	wordListRaw, err := word.ReadDay(ctx, 5)
	if err != nil {
		return nil, err
	}

	wordList := make([]app.WordList, 0, len(wordListRaw))
	for _, v := range wordListRaw {
		wordList = append(wordList, app.WordList{
			Word:    v.Word,
			Explain: v.Explain,
		})
	}

	return &app.DayRes{
		Poem:     poem.Sentence,
		WordList: wordList,
	}, nil
}
