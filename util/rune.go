package util

import "github.com/mattn/go-runewidth"

func ReplaceRuneWithComb(c rune) (r rune, width int, comb []rune) {
	width = runewidth.RuneWidth(c)
	if width == 0 {
		comb = []rune{c}
		c = ' '
		width = 1
	}
	r = c
	return
}
