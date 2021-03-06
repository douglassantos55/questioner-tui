package main

import (
	"container/list"
	"math/rand"
)

type Topic struct {
	Title     string
	Questions []Question

	curIndex int
	visited  *list.List
	current  *list.Element
	selected map[int]*Question
}

type Question struct {
	Statement string
	Answer    string
}

func NewTopic() Topic {
	return Topic{
		visited:  list.New(),
		selected: make(map[int]*Question),
	}
}

func (t *Topic) AddQuestion(question Question) {
	t.Questions = append(t.Questions, question)
}

func (t *Topic) Reset() {
	t.curIndex = 0
	t.visited = list.New()
	t.selected = make(map[int]*Question)
}

func (t *Topic) NextQuestion() *Question {
	if t.current == nil || t.current.Next() == nil {
		next := t.GetRandomQuestion()
		if next != nil {
			t.curIndex++
			t.current = t.visited.PushBack(next)
		}
	} else if t.current.Next() != nil {
		t.curIndex++
		t.current = t.current.Next()
	}
	return t.current.Value.(*Question)
}

func (t *Topic) PrevQuestion() *Question {
	if t.current == nil || t.current.Prev() == nil {
		return nil
	}
	t.curIndex--
	t.current = t.current.Prev()
	return t.current.Value.(*Question)
}

func (t *Topic) GetRandomQuestion() *Question {
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

func (t Topic) Index() int {
	return t.curIndex
}
