package main

import (
    "fmt"
    "os/exec"
)

func threadTask() {
    fmt.Println("This is running in a goroutine.")
}

func main() {
    // Example of Goroutine (Thread-like)
    go threadTask()

    // Example of Process (Running external command)
    cmd := exec.Command("go", "version")
    cmd.Run()
    fmt.Println("Process has finished.")
}
