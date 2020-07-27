# 🤡 clown

clown gives you emojis randomly.

## Installation

```console
$ go get github.com/micnncim/clown/cmd/clown
```

## Usage

```console
$ clown -h
Usage of clown:
  -n int
        Number of emojis. (default 1)
  -o string
        Optional format of output for emojis. json is available.
```

## Examples

```console
$ clown
😀
$ clown -n 5
🇸🇱 🤦‍♀️ 🥖 🧑‍⚕️ 🇯🇲
$ clown -n 3 -o json
[
  {
    "emoji": "🥣",
    "description": "bowl with spoon",
    "category": "Food \u0026 Drink",
    "aliases": [
      "bowl_with_spoon"
    ],
    "tags": [],
    "unicode_version": "11.0",
    "ios_version": "12.1"
  },
  {
    "emoji": "↘️",
    "description": "down-right arrow",
    "category": "Symbols",
    "aliases": [
      "arrow_lower_right"
    ],
    "tags": [],
    "unicode_version": "",
    "ios_version": "6.0"
  },
  {
    "emoji": "🇱🇻",
    "description": "flag: Latvia",
    "category": "Flags",
    "aliases": [
      "latvia"
    ],
    "tags": [],
    "unicode_version": "6.0",
    "ios_version": "8.3"
  }
]
```

## Notice

[`assets/emoji.json`](https://github.com/micnncim/clown/blob/master/assets/emoji.json) is copied from https://github.com/github/gemoji/blob/master/db/emoji.json.
