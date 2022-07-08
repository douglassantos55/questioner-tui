package main

import (
	"container/list"
	"math/rand"
)

type Topic struct {
	Title     string
	Questions []Question

	visited *list.List
	current *list.Element
}

type Question struct {
	Statement string
	Answer    string
}

func NewTopic(title string, questions []Question) Topic {
	return Topic{
		Title:     title,
		Questions: questions,

		visited: list.New(),
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

func (t Topic) GetRandomQuestion() *Question {
	count := 0
	var selected *Question
	for count < len(t.Questions) && selected == nil {
		index := rand.Intn(len(t.Questions))
		selected = &t.Questions[index]
		for cur := t.visited.Front(); cur != nil; cur = cur.Next() {
			if selected == cur.Value.(*Question) {
				selected = nil
			}
		}
		count++
	}
	return selected
}
