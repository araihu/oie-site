package main

import (
	"regexp"
	"strings"
)

// Margo v0.0.35 leaves these search and code-copy strings in English.
var docsLabels = strings.NewReplacer(
	`aria-label="Clear search"`, `aria-label="Limpar busca"`,
	`>Clear</button>`, `>Limpar</button>`,
	`>⌘ K</kbd>`, `>Ctrl/⌘ + K</kbd>`,
	`data-search-kind="Page"`, `data-search-kind="Página"`,
	`<span>Page</span>`, `<span>Página</span>`,
	`data-search-section="default"`, `data-search-section="Documentação"`,
	`>default</span>`, `>Documentação</span>`,
	`aria-label="Copy sh code"`, `aria-label="Copiar código"`,
	`data-code-block-idle-text="Copy"`, `data-code-block-idle-text="Copiar"`,
	`data-code-block-success-text="Copied!"`, `data-code-block-success-text="Copiado!"`,
	`data-code-block-error-text="Unable to copy"`, `data-code-block-error-text="Falha ao copiar"`,
	`aria-live="polite">Copy</span>`, `aria-live="polite">Copiar</span>`,
)

var searchLabels = strings.NewReplacer(
	`"No matching pages."`, `"Nenhuma página encontrada."`,
	`" result available."`, `" resultado encontrado."`,
	`" results available."`, `" resultados encontrados."`,
)

// Goshtoso applies the ranking as CSS order. Use it for Margo's
// aria-activedescendant so it follows the same result as the keyboard.
const searchResultOrder = `    return Array.prototype.filter.call(listbox.querySelectorAll('[role="option"]'), visible);`

const rankedSearchResultOrder = `    return Array.prototype.filter.call(listbox.querySelectorAll('[role="option"]'), visible)
      .sort(function (a, b) { return Number(a.style.order) - Number(b.style.order); });`

var copyLabels = strings.NewReplacer(
	`"Copy"`, `"Copiar"`,
	`"Copy code"`, `"Copiar código"`,
	`"Copied!"`, `"Copiado!"`,
	`"Code copied"`, `"Código copiado"`,
	`"Copy failed"`, `"Falha ao copiar"`,
)

var sectionHeading = regexp.MustCompile(`<h2 id="([^"]+)" data-toc-heading="">([^<]+)</h2>`)

func prepareDocs(content []byte) []byte {
	page := docsLabels.Replace(string(content))
	headings := sectionHeading.FindAllStringSubmatchIndex(page, -1)
	if len(headings) < 2 {
		return []byte(page)
	}

	// Reuse Margo's heading IDs and HTML-escaped labels. The inline index
	// stays with the article during navigation and follows its reading order.
	var toc strings.Builder
	toc.WriteString(`<details class="oie-mobile-toc"><summary>Nesta página</summary><ul>`)
	for _, heading := range headings {
		toc.WriteString(`<li><a href="#` + page[heading[2]:heading[3]] + `">` + page[heading[4]:heading[5]] + `</a></li>`)
	}
	toc.WriteString(`</ul></details>`)

	start := headings[0][0]
	return []byte(page[:start] + toc.String() + page[start:])
}
