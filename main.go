package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	var selectedTopic *Topic
	app := tview.NewApplication()
	pages := tview.NewPages()
	pages.SetBorder(true)

	topicView := tview.NewTextView()
	topicView.SetDoneFunc(func(key tcell.Key) {
		selectedTopic.Reset()
		pages.SwitchToPage("topics")
	})

	topicView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight || event.Rune() == 'l' {
			question := selectedTopic.NextQuestion()
			topicView.SetText(question.Statement)
		}
		if event.Key() == tcell.KeyLeft || event.Rune() == 'h' {
			if question := selectedTopic.PrevQuestion(); question != nil {
				topicView.SetText(question.Statement)
			}
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
		list.AddItem(topic.Title, "", 0, changePage(topic))
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
			if question := selectedTopic.NextQuestion(); question != nil {
				view.(*tview.TextView).SetText(question.Statement)
			}
		}
	})

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
