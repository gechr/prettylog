package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	writer := zerolog.NewConsoleWriter(
		func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = time.TimeOnly
		},
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		reader := bytes.NewReader(scanner.Bytes())
		_, _ = io.Copy(writer, reader)
	}
}
