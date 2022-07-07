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

}
