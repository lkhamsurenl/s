package flickr

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("flickr", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for flickr.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.flickr.com/search/?q=%s", url.QueryEscape(q))
}
