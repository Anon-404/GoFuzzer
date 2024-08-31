package main

import (
    "fmt"
    "net/http"
)

func main() {
    conn, err := http.Get("https://google.com/ome")
    if err != nil {
        panic(err)
    }
    fmt.Println(conn.StatusCode)
}
