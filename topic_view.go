package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TopicView struct {
	topic    *Topic
	question *Question
	View     *tview.TextView

	showAnswer bool
}

func NewTopicView(topic *Topic) *TopicView {
	view := &TopicView{
		topic: topic,
		View:  tview.NewTextView(),
	}

	view.HandleInput()
	view.NextQuestion()
	view.Render()

	return view
}

func (t *TopicView) Render() {
	if t.showAnswer {
		t.View.SetText(fmt.Sprintf("(%d/%d) %s\n\n%s", t.topic.Index(), len(t.topic.Questions), t.question.Statement, t.question.Answer))
	} else {
		t.View.SetText(fmt.Sprintf("(%d/%d) %s", t.topic.Index(), len(t.topic.Questions), t.question.Statement))
	}
}

func (t *TopicView) HandleInput() {
	t.View.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight || event.Rune() == 'l' {
			t.NextQuestion()
		}
		if event.Key() == tcell.KeyLeft || event.Rune() == 'h' {
			t.PrevQuestion()
		}
		if event.Rune() == 'a' {
			t.showAnswer = !t.showAnswer
		}
		t.Render()
		return event
	})
}

func (t *TopicView) Done() {
	t.topic.Reset()
}

func (t *TopicView) PrevQuestion() {
	if question := t.topic.PrevQuestion(); question != nil {
		t.showAnswer = false
		t.question = question
	}
}

func (t *TopicView) NextQuestion() {
	if question := t.topic.NextQuestion(); question != nil {
		t.showAnswer = false
		t.question = question
	}
}
