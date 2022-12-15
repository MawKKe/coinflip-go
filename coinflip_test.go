package main

import (
	"testing"
)

func TestChoose(t *testing.T) {
	always0 := func(int) (int, error) { return 0, nil }
	always1 := func(int) (int, error) { return 1, nil }
	data := []string{"A", "B"}
	t.Run("always get A", func(t *testing.T) {
		res, err := choose(data, always0)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if res != "A" {
			t.Fatalf("Expected %q, got %q", "A", res)
		}
	})
	t.Run("always get B", func(t *testing.T) {
		res, err := choose(data, always1)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if res != "B" {
			t.Fatalf("Expected %q, got %q", "B", res)
		}
	})
	t.Run("Not enough choices", func(t *testing.T) {
		_, err := choose([]string{}, always0)
		if err == nil {
			t.Fatalf("Expected error, got: %v", err)
		}
		_, err = choose([]string{"only one"}, always0)
		if err == nil {
			t.Fatalf("Expected error, got: %v", err)
		}
	})
}
