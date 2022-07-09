package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func PrevQuestion(view *tview.TextView, topic *Topic) {
	if question := topic.PrevQuestion(); question != nil {
		view.SetText(fmt.Sprintf("(%d/%d) %s", topic.Index(), len(topic.Questions), question.Statement))
	}
}

func NextQuestion(view *tview.TextView, topic *Topic) {
	if question := topic.NextQuestion(); question != nil {
		view.SetText(fmt.Sprintf("(%d/%d) %s", topic.Index(), len(topic.Questions), question.Statement))
	}
}

func main() {
	var selectedTopic *Topic
	app := tview.NewApplication()
	pages := tview.NewPages()
	pages.SetBorderPadding(1, 1, 1, 1)

	topicView := tview.NewTextView()
	topicView.SetDoneFunc(func(key tcell.Key) {
		selectedTopic.Reset()
		pages.SwitchToPage("topics")
	})

	topicView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight || event.Rune() == 'l' {
			NextQuestion(topicView, selectedTopic)
		}
		if event.Key() == tcell.KeyLeft || event.Rune() == 'h' {
			PrevQuestion(topicView, selectedTopic)
		}
		return event
	})
	pages.AddPage("topic", topicView, true, false)

	list := tview.NewList()
	loader := NewLocalLoader(NewMarkdownParser())
	loader.GetTopics("test")
	list.SetDoneFunc(func() {
		app.Stop()
	})

	changePage := func(topic Topic) func() {
		return func() {
			selectedTopic = &topic
			pages.SwitchToPage("topic")
		}
	}

	for _, topic := range loader.GetTopics("test") {
		list.AddItem(fmt.Sprintf("%s (%d)", topic.Title, len(topic.Questions)), "", 0, changePage(topic))
	}

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'j' || event.Key() == tcell.KeyDown {
			list.SetCurrentItem(list.GetCurrentItem() + 1)
		}

		if event.Rune() == 'k' || event.Key() == tcell.KeyUp {
			list.SetCurrentItem(list.GetCurrentItem() - 1)
		}

		return event
	})

	pages.AddPage("topics", list, true, true)

	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		}
		return event
	})

	pages.SetChangedFunc(func() {
		name, view := pages.GetFrontPage()
		switch name {
		case "topic":
			NextQuestion(view.(*tview.TextView), selectedTopic)
		}
	})

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
