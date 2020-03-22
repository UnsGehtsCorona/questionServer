package model

import "wirsindcorona/lib"

type Question struct {
	Quid     string
	Sort     int64
	Question string
	Answers  []QuestionAnswer
}

type QuestionAnswer struct {
	Auid   string
	Sort   int64
	Answer string
}

func (q *Question) GenerateQuid() {
	q.Quid = lib.GenerateUid(q.Question)
	for _, answer := range q.Answers {
		answer.GenerateAuid()
	}
}

func (a *QuestionAnswer) GenerateAuid() {
	a.Auid = lib.GenerateUid(a.Answer)
}
