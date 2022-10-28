// Package vhs theme.go contains the information about a terminal theme.
// It stores the 16 base colors as well as the background and foreground colors
// of the terminal theme.
//
// It can be changed through the Set command.
//
// Set Theme {"background": "#171717"}
// Set Theme TokyoNight
package main

import (
	"encoding/json"

	"github.com/charmbracelet/glamour/ansi"
)

// Themes return the list of themes.
func Themes() map[string]Theme {
	return map[string]Theme{
		"TokyoNight": {
			Black:         "#1d202f",
			Red:           "#f7768e",
			Green:         "#9ece6a",
			Yellow:        "#e0af68",
			Blue:          "#7aa2f7",
			Magenta:       "#bb9af7",
			Cyan:          "#7dcfff",
			White:         "#a9b1d6",
			BrightBlack:   "#414868",
			BrightRed:     "#f7768e",
			BrightGreen:   "#a6e3a1",
			BrightYellow:  "#e0af68",
			BrightBlue:    "#7aa2f7",
			BrightMagenta: "#bb9af7",
			BrightCyan:    "#7dcfff",
			BrightWhite:   "#c0caf5",
			Background:    "#24283b",
			Foreground:    "#c0caf5",
			Selection:     "#364a82",
			Cursor:        "#c0caf5",
		},
		"Catppuccin-Latte": {
			Black:         "#5c5f77",
			Red:           "#d20f39",
			Green:         "#40a02b",
			Yellow:        "#df8e1d",
			Blue:          "#1e66f5",
			Magenta:       "#ea76cb",
			Cyan:          "#179299",
			White:         "#acb0be",
			BrightBlack:   "#6c6f85",
			BrightRed:     "#d20f39",
			BrightGreen:   "#40a02b",
			BrightYellow:  "#df8e1d",
			BrightBlue:    "#1e66f5",
			BrightMagenta: "#ea76cb",
			BrightCyan:    "#179299",
			BrightWhite:   "#bcc0cc",
			Background:    "#eff1f5",
			Foreground:    "#4c4f69",
			Selection:     "#acb0be",
			Cursor:        "#dc8a78",
		},
		"Catppuccin-Frappe": {
			Black:         "#51576d",
			Red:           "#e78284",
			Green:         "#a6d189",
			Yellow:        "#e5c890",
			Blue:          "#8caaee",
			Magenta:       "#f4b8e4",
			Cyan:          "#81c8be",
			White:         "#b5bfe2",
			BrightBlack:   "#626880",
			BrightRed:     "#e78284",
			BrightGreen:   "#a6d189",
			BrightYellow:  "#e5c890",
			BrightBlue:    "#8caaee",
			BrightMagenta: "#f4b8e4",
			BrightCyan:    "#81c8be",
			BrightWhite:   "#a5adce",
			Background:    "#303446",
			Foreground:    "#c6d0f5",
			Selection:     "#626880",
			Cursor:        "#f2d5cf",
		},
		"Catppuccin-Macchiato": {
			Black:         "#494d64",
			Red:           "#ed8796",
			Green:         "#a6da95",
			Yellow:        "#eed49f",
			Blue:          "#8aadf4",
			Magenta:       "#f5bde6",
			Cyan:          "#8bd5ca",
			White:         "#b8c0e0",
			BrightBlack:   "#5b6078",
			BrightRed:     "#ed8796",
			BrightGreen:   "#a6da95",
			BrightYellow:  "#eed49f",
			BrightBlue:    "#8aadf4",
			BrightMagenta: "#f5bde6",
			BrightCyan:    "#8bd5ca",
			BrightWhite:   "#b8c0e0",
			Background:    "#24273a",
			Foreground:    "#cad3f5",
			Selection:     "#5b6078",
			Cursor:        "#f4dbd6",
		},
		"Catppuccin-Mocha": {
			Black:         "#45475a",
			Red:           "#f38ba8",
			Green:         "#a6e3a1",
			Yellow:        "#f9e2af",
			Blue:          "#89b4fa",
			Magenta:       "#f5c2e7",
			Cyan:          "#94e2d5",
			White:         "#bac2de",
			BrightBlack:   "#585b70",
			BrightRed:     "#f38ba8",
			BrightGreen:   "#a6e3a1",
			BrightYellow:  "#f9e2af",
			BrightBlue:    "#89b4fa",
			BrightMagenta: "#f5c2e7",
			BrightCyan:    "#94e2d5",
			BrightWhite:   "#a6adc8",
			Background:    "#1e1e2e",
			Foreground:    "#cdd6f4",
			Selection:     "#585b70",
			Cursor:        "#f5e0dc",
		},
	}
}

// Theme is a terminal theme for xterm.js
// It is used for marshalling between the xterm.js readable json format and a
// valid go struct.
// https://xtermjs.org/docs/api/terminal/interfaces/itheme/
type Theme struct {
	Background    string `json:"background"`
	Foreground    string `json:"foreground"`
	Selection     string `json:"selection"`
	Cursor        string `json:"cursor"`
	CursorAccent  string `json:"cursorAccent"`
	Black         string `json:"black"`
	BrightBlack   string `json:"brightBlack"`
	Red           string `json:"red"`
	BrightRed     string `json:"brightRed"`
	Green         string `json:"green"`
	BrightGreen   string `json:"brightGreen"`
	Yellow        string `json:"yellow"`
	BrightYellow  string `json:"brightYellow"`
	Blue          string `json:"blue"`
	BrightBlue    string `json:"brightBlue"`
	Magenta       string `json:"magenta"`
	BrightMagenta string `json:"brightMagenta"`
	Cyan          string `json:"cyan"`
	BrightCyan    string `json:"brightCyan"`
	White         string `json:"white"`
	BrightWhite   string `json:"brightWhite"`
}

func (t Theme) String() string {
	ts, err := json.Marshal(t)
	if err != nil {
		dts, _ := json.Marshal(DefaultTheme)
		return string(dts)
	}
	return string(ts)
}

// DefaultTheme is the default theme to use for recording demos and
// screenshots.
//
// Taken from https://github.com/meowgorithm/dotfiles.
var DefaultTheme = Theme{
	Background:    Background,
	Foreground:    Foreground,
	Cursor:        Foreground,
	CursorAccent:  Background,
	Black:         Black,
	BrightBlack:   BrightBlack,
	Red:           Red,
	BrightRed:     BrightRed,
	Green:         Green,
	BrightGreen:   BrightGreen,
	Yellow:        Yellow,
	BrightYellow:  BrightYellow,
	Blue:          Blue,
	BrightBlue:    BrightBlue,
	Magenta:       Magenta,
	BrightMagenta: BrightMagenta,
	Cyan:          Cyan,
	BrightCyan:    BrightCyan,
	White:         White,
	BrightWhite:   BrightWhite,
}

const margin = 2

// GlamourTheme is the theme for printing out the manual page.
// $ vhs man
var GlamourTheme = ansi.StyleConfig{
	Document: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			BlockPrefix: "\n",
			BlockSuffix: "\n",
		},
		Margin: uintPtr(margin),
	},
	Heading: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			BlockSuffix: "\n",
			Color:       stringPtr("99"),
			Bold:        boolPtr(true),
		},
	},
	Item:     ansi.StylePrimitive{Prefix: "· "},
	Emph:     ansi.StylePrimitive{Color: stringPtr(BrightBlack)},
	Strong:   ansi.StylePrimitive{Bold: boolPtr(true)},
	Link:     ansi.StylePrimitive{Color: stringPtr("42"), Underline: boolPtr(true)},
	LinkText: ansi.StylePrimitive{Color: stringPtr("207")},
	Code:     ansi.StyleBlock{StylePrimitive: ansi.StylePrimitive{Color: stringPtr("204")}},
}

func boolPtr(b bool) *bool       { return &b }
func stringPtr(s string) *string { return &s }
func uintPtr(u uint) *uint       { return &u }
