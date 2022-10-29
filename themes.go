package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

//go:embed themes.json
var themesBts []byte

var (
	themesOnce sync.Once
	themesMap  = map[string]Theme{}
)

// Themes return the list of themes.
var Themes = func() map[string]Theme {
	themesOnce.Do(func() {
		var themes []Theme
		if err := json.Unmarshal(themesBts, &themes); err != nil {
			fmt.Fprintf(os.Stderr, "could not load themes.json: %s\n", err)
			os.Exit(1)
		}
		for _, theme := range themes {
			themesMap[theme.Name] = theme
		}
	})
	return themesMap
}
