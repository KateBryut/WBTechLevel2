package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var cfg Config

const filename = "test.txt"

type Config struct {
	sortColumn    int
	sortByNumber  bool
	sortReverse   bool
	sortNotRepeat bool
}

func (cfg *Config) parce() {
	flag.IntVar(&cfg.sortColumn, "k", -1, "sort column")
	flag.BoolVar(&cfg.sortByNumber, "n", false, "sort by number")
	flag.BoolVar(&cfg.sortReverse, "r", false, "sort reverse")
	flag.BoolVar(&cfg.sortNotRepeat, "u", false, "sort not repeat")
	flag.Parse()
}

func sortByNumbers(rows []string) ([]string, error) {
	var res []string
	var rowsInt []int

	for _, v := range rows {
		n, err := strconv.Atoi(v)

		if err != nil {
			return nil, err
		}
		rowsInt = append(rowsInt, n)
	}

	sort.Ints(rowsInt)
	for _, v := range rowsInt {
		res = append(res, strconv.Itoa(v))
	}
	return res, nil
}

func sortByColumn(rows []string, column int) ([]string, error) {
	var res []string
	var resColumn [][]string

	for _, v := range rows {
		columns := strings.Split(v, " ")
		if len(columns)-1 < column {
			return nil, fmt.Errorf("number of column > string length")
		}
		resColumn = append(resColumn, columns)
	}

	sort.Slice(resColumn, func(i, j int) bool {
		return resColumn[i][column] < resColumn[j][column]
	})

	for _, v := range resColumn {
		res = append(res, strings.Join(v, " "))
	}

	return res, nil
}

func sortReverse(rows []string) ([]string, error) {
	for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
		rows[i], rows[j] = rows[j], rows[i]
	}
	return rows, nil
}

func uniqueRows(rows []string) ([]string, error) {
	var res []string
	uniqueData := make(map[string]string)
	for _, v := range rows {
		if len(uniqueData[v]) == 0 {
			uniqueData[v] = v
		}
	}
	for _, v := range uniqueData {
		res = append(res, v)
	}
	return res, nil
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

func main() {
	cfg.parce()
	fmt.Println(cfg.sortByNumber)
	rows, err := readRowsFromFile(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if cfg.sortNotRepeat {
		rows, err = uniqueRows(rows)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	if cfg.sortColumn != -1 {
		rows, err = sortByColumn(rows, cfg.sortColumn)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	if cfg.sortByNumber {
		rows, err = sortByNumbers(rows)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	if !cfg.sortByNumber && cfg.sortColumn == -1 {
		sort.Strings(rows)
	}
	if cfg.sortReverse {
		rows, err = sortReverse(rows)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	for i := 0; i < len(rows); i++ {
		fmt.Println(rows[i])
	}
}
