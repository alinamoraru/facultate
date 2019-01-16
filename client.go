package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "time"
)

func main() {

    message("Se porneste clientul...")
    conn, err := net.Dial("tcp", "127.0.0.1:4956")
    exitOnError(err)

    before := time.Now()
    message("")
    message("Inainte:  ", before)
    fmt.Fprintf(conn, before.Format(time.RFC3339Nano)+"\n")

    str, err := bufio.NewReader(conn).ReadString('\n')
    exitOnError(err)
    if len(str) < 1 {
        message("Eroare: String invalid!")
        os.Exit(1)
    }

    received, err := time.Parse(time.RFC3339Nano, str[:len(str)-1])
    exitOnError(err)
    message("Primit:", received)

    after := time.Now()
    message("Dupa:   ", after)

    correction := after.Sub(before) / 2

    message("")
    message("Corectie interval: +", correction)
    message("Timpul este", received.Add(correction))
}

func message(a ...interface{}) (n int, err error) {
    return fmt.Print("[C] ", fmt.Sprintln(a...))
}

func exitOnError(err error) {
    if err != nil {
        message("Eroare:", err)
        os.Exit(1)
    }
}
