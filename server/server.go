package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
    "strings"
)

func main() {
    var mode string
    fmt.Print("Enter mode (client/server): ")
    fmt.Scanln(&mode)

    if mode == "server" {
        Server()
    } else if mode == "client" {
        var serverIP string
        fmt.Print("Enter server IP address: ")
        fmt.Scanln(&serverIP)
        Client(serverIP)
    } else {
        log.Fatal("Invalid mode. Please specify 'client' or 'server'.")
    }
}


func Server() {
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
        newmessage := strings.ToUpper(message)
        conn.Write([]byte(newmessage + "\n"))
    }
}

func Client(IP string) {
    conn, err := net.Dial("tcp", IP+":8000")
    if err != nil {
        log.Fatal(err)
        
    }
    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Text to send: ")
        text, _ := reader.ReadString('\n')
        fmt.Fprintf(conn, text+"\n")
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Print("Message from server: " + message)
    }

    
}