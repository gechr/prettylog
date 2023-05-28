package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	writer := zerolog.NewConsoleWriter(
		func(w *zerolog.ConsoleWriter) {
			w.FormatLevel = formatLevel(w.NoColor)
			w.PartsExclude = []string{zerolog.CallerFieldName}
			w.TimeFormat = time.TimeOnly
		},
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		reader := bytes.NewReader(scanner.Bytes())
		if _, err := io.Copy(writer, reader); err != nil {
			fmt.Println(scanner.Text())
		}
	}
}
