package models

import (
	"math/rand"
	"reflect"
)

type Question struct {
	Question         string        `json:"question"`
	CorrectAnswer    string        `json:"correct_answer"`
	IncorrectAnswers []interface{} `json:"incorrect_answers"`
	Difficulty       string        `json:"difficulty"`
}

func (q Question) GetQuestion() string {
	return q.Question
}

func (q Question) GetCorrectAnswer() string {
	return q.CorrectAnswer
}

func (q Question) GetAllOptions() []interface{} {
	options := append(q.IncorrectAnswers, q.CorrectAnswer)
	Shuffle(options)
	return options
}

func Shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}
