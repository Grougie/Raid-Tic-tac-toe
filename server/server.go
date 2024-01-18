package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
    ln, err := net.Listen("tcp", ":8000")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Listening on port 8000")
    conn, err := ln.Accept()
    if err != nil {
        log.Fatal(err)
    }
    for {
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }
        fmt.Print("Message Received:", string(message))
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Text to send: ")
        text, _ := reader.ReadString('\n')
        conn.Write([]byte(text + "\n"))
    }
}