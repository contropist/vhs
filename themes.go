//go:generate sh sync-themes.sh
//go:generate go run . themes --markdown THEMES.md
package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

var (
	//go:embed themes.json
	themesBts []byte

	//go:embed themes_custom.json
	customThemesBts []byte

	themesOnce sync.Once
	allThemes  = map[string]Theme{}
)

func sortedThemeNames() []string {
	keys := make([]string, 0, len(Themes()))
	for k := range Themes() {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})
	return keys
}

// Themes return the list of themes.
func Themes() map[string]Theme {
	themesOnce.Do(func() {
		var themes []windowsTerminalTheme
		for _, bts := range [][]byte{themesBts, customThemesBts} {
			if err := json.Unmarshal(bts, &themes); err != nil {
				fmt.Fprintf(os.Stderr, "could not load themes.json: %s\n", err)
				os.Exit(1)
			}
			for _, theme := range themes {
				allThemes[theme.Name] = Theme{
					Background:    theme.Background,
					Foreground:    theme.Foreground,
					Selection:     theme.SelectionBackground,
					Cursor:        theme.CursorColor,
					CursorAccent:  theme.CursorAccent,
					Black:         theme.Black,
					BrightBlack:   theme.BrightBlack,
					Red:           theme.Red,
					BrightRed:     theme.BrightRed,
					Green:         theme.Green,
					BrightGreen:   theme.BrightGreen,
					Yellow:        theme.Yellow,
					BrightYellow:  theme.BrightYellow,
					Blue:          theme.Blue,
					BrightBlue:    theme.BrightBlue,
					Magenta:       theme.Purple,
					BrightMagenta: theme.BrightPurple,
					Cyan:          theme.Cyan,
					BrightCyan:    theme.BrightCyan,
					White:         theme.White,
					BrightWhite:   theme.BrightWhite,
				}
			}
		}
	})
	return allThemes
}

type windowsTerminalTheme struct {
	Name                string `json:"name"`
	Background          string `json:"background"`
	Foreground          string `json:"foreground"`
	SelectionBackground string `json:"selectionBackground"`
	CursorColor         string `json:"cursorColor"`
	CursorAccent        string `json:"cursorAccent"`
	Black               string `json:"black"`
	BrightBlack         string `json:"brightBlack"`
	Red                 string `json:"red"`
	BrightRed           string `json:"brightRed"`
	Green               string `json:"green"`
	BrightGreen         string `json:"brightGreen"`
	Yellow              string `json:"yellow"`
	BrightYellow        string `json:"brightYellow"`
	Blue                string `json:"blue"`
	BrightBlue          string `json:"brightBlue"`
	Purple              string `json:"purple"`
	BrightPurple        string `json:"brightPurple"`
	Cyan                string `json:"cyan"`
	BrightCyan          string `json:"brightCyan"`
	White               string `json:"white"`
	BrightWhite         string `json:"brightWhite"`
}
