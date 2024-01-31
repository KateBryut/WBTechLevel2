package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var cfg Config

type Config struct {
	timeout     string
	host        string
	port        string
	timeoutTime time.Duration
}

func (cfg *Config) parce() {
	flag.StringVar(&cfg.timeout, "timeout", "10s", "connection timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("use: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	cfg.host = args[0]
	cfg.port = args[1]
}

func main() {
	cfg.parce()
	timeout, err := time.ParseDuration(cfg.timeout)
	if err != nil {
		fmt.Println("invalid timeout value:", err)
		os.Exit(1)
	}

	cfg.timeoutTime = timeout
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	telnet(ctx)
	<-ctx.Done()
}

func telnet(ctx context.Context) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", cfg.host, cfg.port), cfg.timeoutTime)
	if err != nil {
		fmt.Println("error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	//чтение данных из сокета и вывод в STDOUT
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			conn.Close()
			log.Fatalf("error reading from server: %x", err.Error())
			os.Exit(1)
		}
	}()

	//чтение данных из STDIN и запись в сокет,
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data := scanner.Text()
			_, err := fmt.Fprintln(conn, data)
			if err != nil {
				log.Fatalf("error writing to server: %x", err)
				os.Exit(1)
			}
			if err := scanner.Err(); err != nil {
				log.Fatalf("error writing to server: %x", err)
				os.Exit(1)
			}
		}
	}()

	<-ctx.Done()
}
