package internal

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type OsError struct {
	message string
}

// Error returns the error message.
func (e *OsError) Error() string {
	return fmt.Sprintf("os error: %s", e.message)
}

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func GetIP() (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return "", err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return "", err
	}
	defer l.Close()
	return string(l.Addr().(*net.TCPAddr).IP), nil
}

func GetMyIP() (string, error) {
	resp, err := http.Get("http://ifconfig.me/ip")
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return string(ip), nil
	// var ip string
	// os := runtime.GOOS
	// fmt.Println("The os is ", os)
	// var cmd *exec.Cmd
	// if os == "windows" {
	// 	cmd = exec.Command("Invoke-RestMethod", "ifconfig.me/ip")

	// } else if os == "linux" {
	// 	cmd = exec.Command("curl", "ifconfig.me")
	// } else {
	// 	return "", &OsError{
	// 		message: os + " not found",
	// 	}
	// }
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// err := cmd.Run()
	// if err != nil {
	// 	return "", err
	// }
	// ip = out.String()
	// return ip, nil
}
