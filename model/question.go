package model

type Question struct {
	Quid     string
	Sort     int64
	Question string
	Answers  []Answer
}

type Answer struct {
	Auid   string
	Answer string
}
