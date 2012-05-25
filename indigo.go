package main

import (
"fmt"
	"os/user"
	"runtime"
	"strconv"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	targetUID, err := user.Lookup("rwertman")
	if err != nil {
		panic(err)
		return
	}
	UID, err := strconv.Atoi(targetUID.Uid)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(UID)
	err = syscall.Setuid(UID)
	if err != nil {
		panic(err)
	}
	fmt.Println(syscall.Getuid())
	for {}
}
