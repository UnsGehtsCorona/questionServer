package model

type Question struct {
	Quid string
	Question string
	Answers []Answer
}

type Answer struct {
	Auid string
	Answer string
}