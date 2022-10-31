package main

import (
	"reflect"
	"testing"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func TestCommand(t *testing.T) {
	const numberOfCommands = 18
	if len(CommandTypes) != numberOfCommands {
		t.Errorf("Expected %d commands, got %d", numberOfCommands, len(CommandTypes))
	}

	if len(CommandFuncs) != numberOfCommands {
		t.Errorf("Expected %d commands, got %d", numberOfCommands, len(CommandTypes))
	}
}

func TestExecuteSetTheme(t *testing.T) {
	browser := rod.New().MustConnect()
	t.Cleanup(func() {
		browser.MustClose()
	})
	p, err := browser.Page(proto.TargetCreateTarget{})
	if err != nil {
		t.Skip("could not start a browser:", err)
	}

	t.Run("empty", func(t *testing.T) {
		v := &VHS{
			Options: &Options{},
			Page:    p,
		}
		ExecuteSetTheme(Command{
			Type: "Theme",
			Args: "  ",
		}, v)
		if !reflect.DeepEqual(DefaultTheme, v.Options.Theme) {
			t.Errorf("expected theme to be the default theme, got something else: %+v", v.Options.Theme)
		}
	})
	t.Run("named", func(t *testing.T) {
		v := &VHS{
			Options: &Options{},
			Page:    p,
		}
		theme := "Andromeda"
		ExecuteSetTheme(Command{
			Type: "Theme",
			Args: theme,
		}, v)
		expect, _ := findTheme(theme)
		if !reflect.DeepEqual(expect, v.Options.Theme) {
			t.Errorf("expected theme to be %s, got something else: %+v", theme, v.Options.Theme)
		}
	})
	t.Run("json", func(t *testing.T) {
		v := &VHS{
			Options: &Options{},
			Page:    p,
		}
		theme := `{"background": "#29283b"}`
		ExecuteSetTheme(Command{
			Type: "Theme",
			Args: theme,
		}, v)
		if !reflect.DeepEqual("#29283b", v.Options.Theme.Background) {
			t.Errorf("expected theme to be %s, got something else: %+v", theme, v.Options.Theme)
		}
	})
	t.Run("invalid json", func(t *testing.T) {
		v := &VHS{
			Options: &Options{},
			Page:    p,
		}
		theme := `{"backgroun`
		ExecuteSetTheme(Command{
			Type: "Theme",
			Args: theme,
		}, v)
		if !reflect.DeepEqual(DefaultTheme, v.Options.Theme) {
			t.Errorf("expected theme to be %s, got something else: %+v", DefaultTheme, v.Options.Theme)
		}
	})
	t.Run("unknown theme", func(t *testing.T) {
		v := &VHS{
			Options: &Options{},
			Page:    p,
		}
		theme := `foobar`
		ExecuteSetTheme(Command{
			Type: "Theme",
			Args: theme,
		}, v)
		if !reflect.DeepEqual(DefaultTheme, v.Options.Theme) {
			t.Errorf("expected theme to be %s, got something else: %+v", DefaultTheme, v.Options.Theme)
		}
	})
}
