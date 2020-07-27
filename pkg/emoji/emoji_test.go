package emoji

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFilterEmoji(t *testing.T) {
	type args struct {
		emojis []*Emoji
		opt    FilterOption
	}
	tests := []struct {
		name string
		args args
		want []*Emoji
	}{
		{
			name: "filter emojis by tag",
			args: args{
				emojis: []*Emoji{
					{
						Emoji:          "😀",
						Description:    "grinning face",
						Category:       "Smileys & Emotion",
						Aliases:        []string{"grinning"},
						Tags:           []string{"smile", "happy"},
						UnicodeVersion: "6.1",
						IosVersion:     "6.0",
					},
					{
						Emoji:          "🙏",
						Description:    "folded hands",
						Category:       "People & Body",
						Aliases:        []string{"pray"},
						Tags:           []string{"please", "hope", "wish"},
						UnicodeVersion: "6.0",
						IosVersion:     "6.0",
					},
				},
				opt: FilterOption{
					Tag: "happy",
				},
			},
			want: []*Emoji{
				{
					Emoji:          "😀",
					Description:    "grinning face",
					Category:       "Smileys & Emotion",
					Aliases:        []string{"grinning"},
					Tags:           []string{"smile", "happy"},
					UnicodeVersion: "6.1",
					IosVersion:     "6.0",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := FilterEmoji(tt.args.emojis, tt.args.opt)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}
