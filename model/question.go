package model

import "wirsindcorona/lib"

type Question struct {
	Quid     string
	Sort     int64
	Question string
	Answers  []Answer
}

type Answer struct {
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

func (a *Answer) GenerateAuid() {
	a.Auid = lib.GenerateUid(a.Answer)
}
