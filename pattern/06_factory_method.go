package pattern

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Фабрика (фабричный метод) — это порождающий паттерн проектирования, который определяет интерфейс для создания объектов определённого типа.
Паттерн предлагает создавать объекты не напрямую, а через вызов специального фабричного метода. Все объекты должны иметь общий интерфейс,
который отвечает за их создание.

Плюсы:
- Разделение создания объектов от их использования
- Упрощенное добавление новых объектов -> принцип закрытости открытости

Минусы:
- Создание дополнительных иерархий, усложнение читаемости кода
- Появляется "божественный" конструктор, который нужно поддерживать

Паттерн используется, когда:
- заранее неизвестны типы объектов — фабричный метод отделяет код создания объектов от остального кода, где они используются;
- нужна возможность расширять части существующей системы.

Мой пример: подключение к базам данных
*/

type DatabaseConnector interface {
	Query(q string) error
}

type MysqlConnector struct {
}

func newMysqlConnector(dsn string) *MysqlConnector {
	fmt.Println("Connect to mysql")
	return &MysqlConnector{}
}

func (c *MysqlConnector) Query(q string) error {
	fmt.Printf("Query to mysql: %s\n", q)
	return nil
}

type PostgresqlConnector struct {
}

func newPostgresqlConnector(dsn string) *PostgresqlConnector {
	fmt.Println("Connect to postgresql")
	return &PostgresqlConnector{}
}

func (c *PostgresqlConnector) Query(q string) error {
	fmt.Printf("Query to postgresql: %s\n", q)
	return nil
}

// NewConnector реализует фабричный метод.
func NewConnector(dsn string) DatabaseConnector {
	switch {
	case strings.HasPrefix(dsn, "mysql://"):
		return newMysqlConnector(dsn)
	case strings.HasPrefix(dsn, "postgresql://"):
		return newPostgresqlConnector(dsn)
	default:
		panic(fmt.Sprintf("unknown dsn protocol: %s", dsn))
	}
}

func mainFactory() {
	mysql := NewConnector("mysql://...")
	mysql.Query("SELECT something FROM list")

	pg := NewConnector("postgresql://...")
	pg.Query("SELECT something FROM list")
}
