package kutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type UtilError string

func (err UtilError) Error() string { return "Redis Error: " + string(err) }

func ProcessName() string {
	file := os.Args[0]
	i := strings.LastIndex(file, "\\")
	j := strings.LastIndex(file, "/")
	if j < i {
		file = file[i+1:]
	} else if j > i {
		file = file[j+1:]
	}
	i = strings.LastIndex(file, ".")
	if i > 0 {
		file = file[0:i]
	}
	return file
}

func DefaultPidFileName() string {
	file := ProcessName()
	pidFile := "/var/run/"
	if runtime.GOOS == "windows" {
		pidFile = "c:\\" + file + ".pid"
	} else {
		pidFile += file + ".pid"
	}
	return pidFile
}

func WritePidFile(myFile string, pid int) (err error) {
	return ioutil.WriteFile(myFile, []byte(fmt.Sprintf("%d", pid)), 0644)
}

func CheckWritePidPermission(pidFile string) error {
	if len(pidFile) <= 0 {
		pidFile = DefaultPidFileName()
	}
	if err := ioutil.WriteFile(pidFile, []byte(fmt.Sprintf("%d", 0)), 0644); err != nil {
		fmt.Sprintln("had no permission to write pid file: ", pidFile, err)
		return UtilError("sdfdsf")
	}
	return nil
}

func CreateProcess(background bool, file string, args []string) error {
	filePath, _ := filepath.Abs(file)
	cmd := exec.Command(filePath, args...)
	if background {
		cmd.Stdin = nil //给新进程设置文件描述符，可以重定向到文件中
		cmd.Stdout = nil
		cmd.Stderr = nil
	} else {
		cmd.Stdin = os.Stdin //给新进程设置文件描述符，可以重定向到文件中
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Start()
}

func CreateProcess2(background bool, file string, args []string) (*os.Process, error) {
	filePath, _ := filepath.Abs(file)
	if background {
		return os.StartProcess(filePath, args, &os.ProcAttr{Files: []*os.File{nil, nil, nil}})
	}
	return os.StartProcess(filePath, args, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
}

func WaitProcess(background bool, file string, args []string) ([]byte, error) {
	filePath, _ := filepath.Abs(file)
	cmd := exec.Command(filePath, args...)
	if background {
		cmd.Stdin = nil //给新进程设置文件描述符，可以重定向到文件中
		cmd.Stdout = nil
		cmd.Stderr = nil
	} else {
		cmd.Stdin = os.Stdin //给新进程设置文件描述符，可以重定向到文件中
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	//cmd.Run()不能返回内容，而Output可以返回内容。
	return cmd.Output()
}
