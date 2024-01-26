package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func command(args []string) error {
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("missing argument")
		}
		err := os.Chdir(args[1])
		if err != nil {
			return err
		}
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		if len(args) < 2 {
			return errors.New("missing argument")
		}
		pid := args[1]
		err := exec.Command("kill", pid).Run()
		if err != nil {
			return err
		}
	case "ps":
		cmd := exec.Command("ps")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	case "exec":
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	default:
		return errors.New("unknowing command")
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		user, _ := user.Current()
		dir, _ := os.Getwd()
		fmt.Printf("%s:%s$ ", user.Username, dir)
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		input = strings.TrimSpace(input)
		commands := strings.Split(input, "|")
		for _, com := range commands {
			args := strings.Fields(com)
			if len(args) == 0 {
				continue
			}
			err = command(args)
			if err != nil {
				log.Printf("can't execute command %s, err: %s", args[0], err.Error())
				os.Exit(1)
			}
		}
	}
}
