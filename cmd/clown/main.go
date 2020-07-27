package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rakyll/statik/fs"

	"github.com/micnncim/clown/pkg/emoji"
	_ "github.com/micnncim/clown/pkg/statik"
)

var (
	nFlag = flag.Int("n", 1, "Number of emojis.")
	oFlag = flag.String("o", "", "Optional format of output for emojis. json is available.")
)

func main() {
	flag.Parse()

	if err := run(*nFlag, *oFlag); err != nil {
		log.Fatal(err)
	}
}

func run(n int, output string) error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	r, err := statikFS.Open("/emoji.json")
	if err != nil {
		return err
	}
	defer r.Close()

	emojis, err := emoji.NewEmojis(r)
	if err != nil {
		return err
	}

	filtered := emoji.FilterEmoji(emojis, emoji.FilterOption{
		Number: n,
	})

	switch output {
	case "":
		for _, e := range filtered {
			fmt.Printf("%s ", e.Emoji)
		}
		fmt.Printf("\n")
	case "json":
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(filtered); err != nil {
			return err
		}
	}

	return nil
}
