package main

import (
    "fmt"
    "os"
    "time"
)


func main() {
    today := time.Now().Format("2006-01-02")

    filename := today + ".md"

    _, err := os.Stat(filename)

    if os.IsNotExist(err) {
        file, err := os.Create(filename)
        if err != nil {
            fmt.Printf("Error creating file: %v\n", err)
            os.Exit(1);
        }
        defer file.Close()
        fmt.Printf("created note: %s\n", filename);
    } else if err != nil {
        fmt.Printf("Error checking file %v\n", err)
        os.Exit(1)
    } else {
        fmt.Printf("Note for today (%s) already exists\n", filename)
    }
}
