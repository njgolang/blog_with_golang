package main

import (
    "fmt"
    "os/exec"
)

func runcmd(cmdString string) string {
    cmd := exec.Command("/bin/sh", "-c", cmdString)
    bytes, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
        return ""
    }   
    data := string(bytes)
    return data
}
