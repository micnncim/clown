package emoji

import (
	"encoding/json"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Emoji struct {
	Emoji          string   `json:"emoji"`
	Description    string   `json:"description"`
	Category       string   `json:"category"`
	Aliases        []string `json:"aliases"`
	Tags           []string `json:"tags"`
	UnicodeVersion string   `json:"unicode_version"`
	IosVersion     string   `json:"ios_version"`
}

type FilterOption struct {
	// Number is used to specify how many emojis are returned after though all other filters.
	// If Number is not specified, functions return the all filtered emojis.
	Number int
	Tag    string
}

func NewEmojis(r io.Reader) ([]*Emoji, error) {
	var emojis []*Emoji

	if err := json.NewDecoder(r).Decode(&emojis); err != nil {
		return nil, err
	}

	return emojis, nil
}

func FilterEmojis(emojis []*Emoji, opt FilterOption) []*Emoji {
	filtered := make([]*Emoji, 0, len(emojis))

	if opt.Tag != "" {
		for _, e := range emojis {
			for _, tag := range e.Tags {
				if tag == opt.Tag {
					filtered = append(filtered, e)
				}
			}
		}
	}

	// If all the options are not specified, copy all the emojis to filtered emojis.
	if opt.Tag == "" {
		filtered = make([]*Emoji, len(emojis))
		copy(filtered, emojis)
	}

	if opt.Number > 0 {
		picked := make([]*Emoji, 0, len(filtered))
		for i := 0; i < opt.Number; i++ {
			picked = append(picked, filtered[rand.Intn(len(filtered))])
		}
		return picked
	}

	return filtered
}
