package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var cfg Config

type Config struct {
	host string
	port int
	udp  bool
}

func (cfg *Config) parce() {
	flag.StringVar(&cfg.host, "host", "127.0.0.1", "fields to select")
	flag.IntVar(&cfg.port, "port", 80, "delimiter")
	flag.BoolVar(&cfg.udp, "udp", false, "use udp")

	flag.Parse()
}

func main() {
	cfg.parce()

	address := fmt.Sprintf("%s:%d", cfg.host, cfg.port)

	var conn net.Conn
	var err error

	if cfg.udp {
		conn, err = net.Dial("udp", address)
	} else {
		conn, err = net.Dial("tcp", address)
	}

	if err != nil {
		fmt.Println("can't connect:", err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		io.Copy(conn, os.Stdin)
	}()

	io.Copy(os.Stdout, conn)
}
