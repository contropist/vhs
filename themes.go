//go:generate make all
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
		for _, bts := range [][]byte{themesBts, customThemesBts} {
			var themes []Theme
			if err := json.Unmarshal(bts, &themes); err != nil {
				fmt.Fprintf(os.Stderr, "could not load themes.json: %s\n", err)
				os.Exit(1)
			}
			for _, theme := range themes {
				allThemes[theme.Name] = theme
			}
		}
	})
	return allThemes
}
