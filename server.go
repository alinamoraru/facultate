package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "time"
)

func main() {

    message("Se porneste serverul...")
    ln, err := net.Listen("tcp", ":4956")
    exitOnError(err)
    for {
        message("Asteapta conexiunea...")
        conn, err := ln.Accept()
        exitOnError(err)
        go handleConnection(conn)
    }

    message("Finalizat")

}

func handleConnection(conn net.Conn) {

    message("")
    message("Se prelucreaza conexiunea...")
    str, err := bufio.NewReader(conn).ReadString('\n')
    exitOnError(err)
    if len(str) > 0 {
        t := time.Now()

        conn.Write([]byte(t.Format(time.RFC3339Nano) + "\n"))

        received, err := time.Parse(time.RFC3339Nano, str[:len(str)-1])
        exitOnError(err)

        message("Primit:", received)
        message("Trimis:    ", t)
    }

    conn.Close()
}

func message(a ...interface{}) (n int, err error) {
    return fmt.Print("[S] ", fmt.Sprintln(a...))
}

func exitOnError(err error) {
    if err != nil {
        message("Eroare:", err)
        os.Exit(1)
    }
}
