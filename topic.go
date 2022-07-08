package main

import (
	"container/list"
	"math/rand"
)

type Topic struct {
	Title     string
	Questions []Question

	visited  *list.List
	current  *list.Element
	selected map[int]*Question
}

type Question struct {
	Statement string
	Answer    string
}

func NewTopic(title string, questions []Question) Topic {
	return Topic{
		Title:     title,
		Questions: questions,

		visited:  list.New(),
		selected: make(map[int]*Question),
	}
}

func (t *Topic) AddQuestion(question Question) {
	t.Questions = append(t.Questions, question)
}

func (t *Topic) NextQuestion() *Question {
	if t.current == nil || t.current.Next() == nil {
		next := t.GetRandomQuestion()
		if next != nil {
			t.current = t.visited.PushBack(next)
		}
	}
	return t.current.Value.(*Question)
}

func (t *Topic) PrevQuestion() *Question {
	if t.current == nil || t.current.Prev() == nil {
		return nil
	}
	t.current = t.current.Prev()
	return t.current.Value.(*Question)
}

func (t Topic) GetRandomQuestion() *Question {
	if t.visited.Len() == len(t.Questions) {
		return nil
	}

	var index int
	for t.selected[index] != nil {
		index = rand.Intn(len(t.Questions))
	}

	question := &t.Questions[index]
	t.selected[index] = question
	return question
}
