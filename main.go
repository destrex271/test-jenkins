package main

import(
    "fmt"
    "net"
    "time"
)

func handleRequest(conn net.Conn){
    var data []byte;
    conn.Read(data)
    fmt.Println(time.Now().String() + " " + string(data))
}

func main(){
    fmt.Printf("Initalizing Server...\n");

    ln, err := net.Listen("tcp", ":8090");
    if err != nil{
        fmt.Println("Error: ", err);
    }

    fmt.Println("Waiting for connections...");
    for{
        conn, err := ln.Accept()
        if err != nil{
            fmt.Println("Error: ", err);
        }
        go handleRequest(conn);
    }
}
