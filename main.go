package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	pages := tview.NewPages()
	pages.SetBorderPadding(1, 1, 1, 1)

	list := tview.NewList()

	loader := NewLocalLoader(NewMarkdownParser())
	loader.GetTopics("test")

	list.SetDoneFunc(func() {
		app.Stop()
	})

	changePage := func(topic Topic) func() {
		return func() {
			view := NewTopicView(&topic)
			pages.AddAndSwitchToPage("topic", view.View, true)

			view.View.SetDoneFunc(func(key tcell.Key) {
				view.Done()
				pages.RemovePage("topic")
			})
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

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
