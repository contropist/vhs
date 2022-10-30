package main

import "testing"

func TestThemesTest(t *testing.T) {
	for i := 0; i < 5; i++ {
		themes := Themes()
		expect := 264
		if l := len(themes); l != expect {
			t.Errorf("expected to load %d themes, got %d", expect, l)
		}
	}
}
