package d2svg

import (
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/svg"
)

// Copied private functions from chroma. Their public functions do too much (write the whole SVG document)
// https://github.com/alecthomas/chroma
// >>> BEGIN

var svgEscaper = strings.NewReplacer(
	`&`, "&amp;",
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&quot;",
	` `, "&#160;",
	`	`, "&#160;&#160;&#160;&#160;",
)

func styleToSVG(style *chroma.Style) map[chroma.TokenType]string {
	converted := map[chroma.TokenType]string{}
	// NOTE this is in the original source code, but it just makes unhighlightable code turn into the bg color
	// Which I don't understand, and I get the results I want when I remove it.
	// bg := style.Get(chroma.Background)
	for t := range chroma.StandardTypes {
		entry := style.Get(t)
		// if t != chroma.Background {
		//   entry = entry.Sub(bg)
		// }
		if entry.IsZero() {
			continue
		}
		converted[t] = svg.StyleEntryToSVG(entry)
	}
	return converted
}

func styleAttr(styles map[chroma.TokenType]string, tt chroma.TokenType) string {
	if _, ok := styles[tt]; !ok {
		tt = tt.SubCategory()
		if _, ok := styles[tt]; !ok {
			tt = tt.Category()
			if _, ok := styles[tt]; !ok {
				return ""
			}
		}
	}
	// Custom code
	out := styles[tt]
	classes := []string{}

	if strings.Contains(out, `font-weight="bold"`) {
		out = strings.Replace(out, `font-weight="bold"`, ``, 1)
		classes = append(classes, "text-mono-bold")
	}
	if strings.Contains(out, `font-style="italic"`) {
		out = strings.Replace(out, `font-style="italic"`, ``, 1)
		classes = append(classes, "text-mono-italic")
	}

	if len(classes) > 0 {
		out += `class="` + strings.Join(classes, " ") + `"`
	}

	return strings.TrimSpace(out)
}

// <<< END
