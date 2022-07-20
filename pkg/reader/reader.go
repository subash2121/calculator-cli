package reader

import (
	"bufio"
	"io"
)

func Reader(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
