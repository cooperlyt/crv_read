package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"syscall"

	"github.com/axgle/mahonia"
)

type Preson struct {
	Name        string  `json:"name"`
	Sex         string  `json:"sex"`
	Ethnic      string  `json:"ethnic"`
	Birthday    string  `json:"birthday"`
	Address     string  `json:"address"`
	Id          string  `json:"id"`
	Agency      string  `json:"agency"`
	ExpireStart *string `json:"expire_start"`
	ExpireEnd   *string `json:"expire_end"`
	Picture     string  `json:"picture"`
}

type Content struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Person  Preson `json:"person"`
}

var mu sync.Mutex

var deveice_init = -1

func init_deveice(handle syscall.Handle) int {
	p_init, err := syscall.GetProcAddress(handle, "CVR_InitComm")
	if err != nil {
		fmt.Println("Error getting init proc address:", err)
		return -1
	}
	r := -1
	port := 1001
	for r != 1 && port <= 1016 {
		r_init, r2, err := syscall.SyscallN(p_init, uintptr(port))
		r = int(r_init)
		fmt.Println(port, r, r2, err)
		port++

	}
	return r
}

func convert_expire(expire string) *string {
	re := regexp.MustCompile(`(\d{4})(\d{2})(\d{2})`)
	matches := re.FindStringSubmatch(expire)

	if len(matches) != 4 {
		fmt.Println("Invalid date format")
		return nil
	}
	result := fmt.Sprintf("%04s-%02s-%02sT00:00:00.000+08:00", matches[1], matches[2], matches[3])
	return &result

}

func read_content_file() Content {
	contentBytes, err := os.ReadFile("wz.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return Content{6, "Error reading file!", Preson{}}
	}
	contentStr := mahonia.NewDecoder("GBK").ConvertString(string(contentBytes))
	fields := strings.Split(contentStr, "\r\n")

	f, err := os.ReadFile("zp.bmp")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return Content{7, "Error reading photo file!", Preson{}}
	}
	base64Encoding := "data:image/bmp;base64," + base64.StdEncoding.EncodeToString(f)

	expire := strings.Split(fields[7], "-")

	fmt.Println("Expire:", expire[0], expire[1])

	return Content{0, "Success!", Preson{fields[0], fields[1], fields[2], fields[3] + "T00:00:00.000+08:00", fields[4], fields[5], fields[6], convert_expire(expire[0]), convert_expire(expire[1]), base64Encoding}}
}

func read_content(handle syscall.Handle) Content {

	var errOccurred bool

	defer func() {
		if r := recover(); r != nil {
			errOccurred = true
			fmt.Println("Error occurred:", r)
		}
	}()

	if deveice_init != 1 {
		deveice_init = init_deveice(handle)
	}
	if deveice_init != 1 {
		return Content{1, "Deveice not found!", Preson{}}
	}
	p_authenticate, err := syscall.GetProcAddress(handle, "CVR_AuthenticateForNoJudge")
	if err != nil {
		fmt.Println("Error getting authenticate proc address:", err)
		return Content{2, "Error getting authenticate proc address!", Preson{}}
	}

	r_authenticate, _, err := syscall.SyscallN(p_authenticate)
	fmt.Println("Authenticate:", r_authenticate, err)

	if r_authenticate != 1 {
		fmt.Println("Error: Authenticate failed!", err)
		return Content{3, "Error: Authenticate failed!", Preson{}}
	}

	p_read_content, err := syscall.GetProcAddress(handle, "CVR_Read_FPContent")
	if err != nil {
		fmt.Println("Error getting read content proc address:", err)
		return Content{4, "Error getting read content proc address!", Preson{}}
	}

	r_read_content, _, err := syscall.SyscallN(p_read_content)
	fmt.Println("Read content:", r_read_content, err)
	if r_read_content != 1 {
		fmt.Println("Error: Read content failed!", err)
		return Content{5, "Error: Read content failed!", Preson{}}
	}

	if errOccurred {
		return Content{8, "Error occurred!", Preson{}}
	}

	return read_content_file()
}

var lib_handle syscall.Handle

func read(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization")
	w.Header().Set("Content-Type", "application/json")

	if req.Method == http.MethodOptions {
		// 处理预检请求（OPTIONS 请求），返回成功状态
		w.WriteHeader(http.StatusOK)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	content := read_content(lib_handle)

	json.NewEncoder(w).Encode(content)

}

func main() {
	fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
	handle, err := syscall.LoadLibrary("Termb.dll")

	if err != nil {
		fmt.Println("Error loading DLL:", err)
		return
	}
	lib_handle = handle

	fmt.Println("DLL loaded successfully handle:", handle)
	defer syscall.FreeLibrary(handle)

	http.HandleFunc("/read", read)

	http.ListenAndServe(":56541", nil)

}
