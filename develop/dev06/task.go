package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cfg Config

type Config struct {
	fields    string
	delimiter string
	separated bool
}

func (cfg *Config) parce() {
	flag.StringVar(&cfg.fields, "f", "", "fields to select")
	flag.StringVar(&cfg.delimiter, "d", "\t", "delimiter")
	flag.BoolVar(&cfg.separated, "s", false, "only separated lines")

	flag.Parse()
}

func main() {
	cfg.parce()

	selectedFields := make(map[int]bool)
	if cfg.fields != "" {
		fields := strings.Split(cfg.fields, ",")
		for _, field := range fields {
			index, err := strconv.Atoi(field)
			if err != nil {
				log.Fatalln("can't parse column to int: ", err.Error())
				os.Exit(1)
			}
			selectedFields[index] = true
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if cfg.separated && !strings.Contains(line, cfg.delimiter) {
			continue
		}
		var builder strings.Builder

		parts := strings.Split(line, cfg.delimiter)
		for i, part := range parts {
			if len(selectedFields) == 0 || selectedFields[i+1] {
				builder.WriteString(part)
				builder.WriteString(" ")
			}
		}
		fmt.Println(builder.String())
	}
}
