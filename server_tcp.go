/*
A very simple TCP server written in Go.
This is a toy project that I used to learn the fundamentals of writing
Go code and doing some really basic network stuff.
Maybe it will be fun for you to read. It's not meant to be
particularly idiomatic, or well-written for that matter.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var addr = flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 9000, "The port to listen on; default is 8000.")

func main() {
	flag.Parse()

	fmt.Println("Starting server...")

	src := *addr + ":" + strconv.Itoa(*port)
	listener, _ := net.Listen("tcp", src)
	fmt.Printf("Listening on %s.\n", src)
	fileLog, err := NewAndRead("logger.log", "./")
	if err != nil {
		logrus.Panicln("Error Create file log: ", err)
	}
	logrus.SetOutput(fileLog)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	scanner := bufio.NewScanner(conn)

	for {
		ok := scanner.Scan()
		time.Now().String()
		if len(scanner.Text()) > 0 {
			fmt.Println("Time : ", time.Now().UTC().String(), "-------- Msg ", scanner.Text())
			logrus.Info("Time: ", time.Now().UTC().String(), " , Msg: ", scanner.Text())
			conn.Write([]byte(scanner.Text() + "--- Time: " + time.Now().UTC().String() + "\n"))
		}
		if !ok {
			break
		}
	}
	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func handleMessage(message string, conn net.Conn) {
	fmt.Println("> " + message)
	if len(message) > 0 && message[0] == '/' {
		switch {
		case message == "/time":
			resp := "It is " + time.Now().String() + "\n"
			fmt.Print("< " + resp)
			conn.Write([]byte(resp))

		case message == "/quit":
			fmt.Println("Quitting.")
			conn.Write([]byte("I'm shutting down now.\n"))
			fmt.Println("< " + "%quit%")
			conn.Write([]byte("%quit%\n"))
			os.Exit(0)

		default:
			conn.Write([]byte("Unrecognized command.\n"))
		}
	}
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// NewAndRead new and read file
func NewAndRead(fileName string, path string) (*os.File, error) {
	var (
		err  error
		fReq *os.File
	)
	if exist := FileExists(fmt.Sprintf("%s/%s", path, fileName)); exist == false {
		if fReq, err = os.Create(path + fileName); err != nil {
			return fReq, err
		}
	}
	fReq, err = os.OpenFile(fmt.Sprintf("%s/%s", path, fileName), os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return fReq, err
	}

	return fReq, err
}
