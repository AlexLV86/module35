package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

// Сетевой адрес.
const addr = ":12345"

// Протокол сетевой службы.
const proto = "tcp4"

// Время задержки в секундах
const delay = 3

var proverb = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	// Запуск сетевой службы по протоколу TCP
	// на порту 12345.
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	// Подключения обрабатываются в бесконечном цикле.
	for {
		// Принимаем подключение.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Вызов обработчика подключения.
		go handleConn(conn)
	}
}

// handleConn обработчик соединения с клиентом
func handleConn(conn net.Conn) {
	// Закрытие соединения.
	defer conn.Close()
	for {
		_, err := conn.Write([]byte(proverb[rand.Intn(len(proverb))] + "\r\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second * delay)
	}
}
