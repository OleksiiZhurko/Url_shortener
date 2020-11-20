package response

import (
	"URL_shortener/pkg/config"
	"encoding/json"
)

type URLs struct {
	Link    string `json:"link"`    // full link
	Shorten string `json:"shorten"` // shorten link
}

// Returns Urls struct
func NewURLs() *URLs {
	return &URLs{}
}

// URLs.Link getter
func (links URLs) GetFullLink() string {
	return links.Link
}

// URLs.Shorten getter
func (links URLs) GetShortenedLink() string {
	return links.Shorten
}

// URLs.Link and URLs.Shorten setter
func (links *URLs) SetParams(fullLink string, shortLink string) {
	links.Link = fullLink
	links.Shorten = shortLink
}

// Comparing by URLs.Shorten
func (links URLs) CompareByShorten(shorten string) bool {
	return links.Shorten == shorten
}

// Returns URLs struct in JSON type
func (links URLs) ConvertToJSON() []byte {
	links.Shorten = config.GetConfigHOST() + config.GetConfigPORT() + "/" + links.Shorten
	result, _ := json.Marshal(links)
	return result
}
