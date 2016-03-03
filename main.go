package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var raw = `
---
hello
---
world
---
wtf
-----
owowow
---
---
123
456
789
---
xyz
---
`

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Printf("%t\t%d\t%s\n", atEOF, len(data), data)
		if atEOF {
			return 0, nil, errors.New("bad luck")
		}

		// Skip leading spaces.
		start := 0
		for width := 0; start < len(data); start += width {
			var r rune
			r, width = utf8.DecodeRune(data[start:])
			if string(r) == "-" {
				break
			}
		}

		return 0, nil, nil
	}

	// input := strings.Repeat("abc ", 5)
	input := raw
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(split)
	buf := make([]byte, 8)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}
