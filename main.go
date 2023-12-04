package main

import (
	"fmt"
	"runtime"
	"syscall"
)

func main() {
	fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
	handle, err := syscall.LoadLibrary("Termb.dll")

	if err != nil {
		fmt.Println("Error loading DLL:", err)
		return
	}

	fmt.Println("DLL loaded successfully handle:", handle)
	defer syscall.FreeLibrary(handle)
	p_init, err := syscall.GetProcAddress(handle, "CVR_InitComm")
	if err != nil {
		fmt.Println("Error getting init proc address:", err)
		return
	}
	r := -1
	port := 1001
	for r != 1 && port <= 1016 {
		r_init, r2, err := syscall.SyscallN(p_init, uintptr(port))
		r = int(r_init)
		fmt.Println(port, r, r2, err)
		port++

	}
	if r != 1 {
		fmt.Println("Error: Deveice not found!")
		return
	}

	p_authenticate, err := syscall.GetProcAddress(handle, "CVR_AuthenticateForNoJudge")
	if err != nil {
		fmt.Println("Error getting authenticate proc address:", err)
		return
	}

	r_authenticate, _, err := syscall.SyscallN(p_authenticate)
	fmt.Println("Authenticate:", r_authenticate, err)

	if r_authenticate != 1 {
		fmt.Println("Error: Authenticate failed!", err)
		return
	}

	p_read_content, err := syscall.GetProcAddress(handle, "CVR_Read_FPContent")
	if err != nil {
		fmt.Println("Error getting read content proc address:", err)
		return
	}

	r_read_content, _, err := syscall.SyscallN(p_read_content)
	fmt.Println("Read content:", r_read_content, err)
	if r_read_content != 1 {
		fmt.Println("Error: Read content failed!", err)
		return
	}

	fmt.Println("Read content success!")

	fmt.Println("Hello, World!")
}
