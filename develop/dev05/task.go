package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var cfg Config

type Config struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func readRowsFromFile(filename string) ([]string, error) {
	var res []string
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		return res, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res, nil
}

func (cfg *Config) parce() {
	flag.IntVar(&cfg.after, "A", 0, "print +N lines after match")
	flag.IntVar(&cfg.before, "B", 0, "print +N lines until match")
	flag.IntVar(&cfg.context, "C", 0, "print ±N lines around match")
	flag.BoolVar(&cfg.count, "c", false, "number of lines")
	flag.BoolVar(&cfg.ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&cfg.invert, "v", false, "instead of matching, exclude")
	flag.BoolVar(&cfg.fixed, "F", false, "exact match to string, not a pattern")
	flag.BoolVar(&cfg.lineNum, "n", false, "print line number")
	flag.Parse()
}

func main() {
	cfg.parce()
	fmt.Println(cfg.after)

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("missed filename or pattern")
		os.Exit(1)
	}

	filename := args[1]
	pattern := args[0]

	lines, err := readRowsFromFile(filename)
	if err != nil {
		fmt.Println("can't read file:", err)
		os.Exit(1)
	}

	var re *regexp.Regexp
	if cfg.fixed {
		re = regexp.MustCompile(regexp.QuoteMeta(pattern))
	} else {
		if cfg.ignoreCase {
			re = regexp.MustCompile("(?i)" + pattern)
		} else {
			re = regexp.MustCompile(pattern)
		}
	}

	matchedLines := make([]string, 0)
	for i, line := range lines {
		matched := re.MatchString(line)
		if (cfg.invert && !matched) || (!cfg.invert && matched) {
			matchedLines = append(matchedLines, line)
			if cfg.count {
				continue
			}
			if cfg.lineNum {
				fmt.Printf("№ %d: ", i+1)
			}
			fmt.Println(line)
			if cfg.after > 0 {
				for j := 1; j <= cfg.after && i+j < len(lines); j++ {
					fmt.Println(lines[i+j])
				}
			}
			if cfg.before > 0 {
				for j := 1; j <= cfg.before && i-j >= 0; j++ {
					fmt.Println(lines[i-j])
				}
			}
			if cfg.context > 0 {
				for j := 1; j <= cfg.context && i+j < len(lines); j++ {
					fmt.Println(lines[i+j])
				}
				for j := 1; j <= cfg.context && i-j >= 0; j++ {
					fmt.Println(lines[i-j])
				}
			}
		}
	}

	if cfg.count {
		fmt.Println("count of strings:", len(matchedLines))
	}
}
