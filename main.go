package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	w := zerolog.NewConsoleWriter(
		func(w *zerolog.ConsoleWriter) {
			w.FormatLevel = formatLevel(w.NoColor)
			w.PartsExclude = []string{zerolog.CallerFieldName}
			w.TimeFormat = "15:04:05.000"
		},
	)

	br := bufio.NewReader(os.Stdin)
	for {
		line, err := br.ReadBytes('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}
		if err == io.EOF {
			break
		}
		r := bytes.NewReader(line)
		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(string(line))
		}
	}
}
