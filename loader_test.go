package main

import (
	"reflect"
	"testing"
)

func TestLoader(t *testing.T) {
	t.Run("GetFiles", func(t *testing.T) {
		t.Run("recurses", func(t *testing.T) {
			loader := NewLocalLoader(NewMarkdownParser())

			got := loader.GetFiles("test")

			expected := []string{
				"test/clean-code.md",
				"test/laravel/queue.md",
				"test/laravel/service-container.md",
				"test/magento/plugins.md",
				"test/magento/themes.md",
			}

			if !reflect.DeepEqual(expected, got) {
				t.Errorf("Expected %v, got %v", expected, got)
			}
		})

		t.Run("no subdirectories", func(t *testing.T) {
			loader := NewLocalLoader(NewMarkdownParser())
			got := loader.GetFiles("test/magento")
			expected := []string{
				"test/magento/plugins.md",
				"test/magento/themes.md",
			}

			if !reflect.DeepEqual(got, expected) {
				t.Errorf("Expected %v, got %v", expected, got)
			}
		})
	})

	t.Run("GetTopics", func(t *testing.T) {
		loader := NewLocalLoader(NewMarkdownParser())
		got := loader.GetTopics("test")

		expected := []Topic{
			{
				Title: "Clean Code",
				Questions: []Question{
					{Statement: "What does SOLID mean?", Answer: "SOLID"},
					{Statement: "Explain dependency inversion", Answer: "Interfaces rock!"},
				},
			},
			{
				Title: "Laravel Queues",
				Questions: []Question{
					{
						Statement: "How to register things to queue?",
						Answer:    "Implement the ShouldQueue interface",
					},
				},
			},
			{
				Title: "Laravel - Service Container",
				Questions: []Question{
					{
						Statement: "How to bind concrete implementations to interfaces?",
						Answer:    "$app->bind(Abstract::class, Concrete::class)",
					},
				},
			},
			{
				Title: "Magento Plugins",
				Questions: []Question{
					{
						Statement: "Which are the required files to create a plugin?",
						Answer:    "Some xml and composer.json",
					},
				},
			},
			{
				Title: "Magento Themes",
				Questions: []Question{
					{
						Statement: "Which are the required files to create a theme?",
						Answer:    "theme.xml and more",
					},
				},
			},
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
