package wolframalpha

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("wolframalpha", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for WolframAlpha.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.wolframalpha.com/input/?i=%s", url.QueryEscape(q))
}
