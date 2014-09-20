package kutil

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type KLog struct {
	logFile *os.File
	log     *log.Logger
	wr      io.Writer
	day     string

	prefix string
	path   string
	name   string
}

func (k *KLog) New(path string, name string, prefix string) (err error) {
	k.prefix = prefix
	k.path = path
	k.name = name
	return k.init()
}

func (k *KLog) init() (err error) {
	if len(k.name) == 0 {
		return errors.New("not init.")
	}
	day := time.Now().Format("20060102")
	if k.day == day {
		return nil
	}
	k.Close()
	file := filepath.Clean(k.path + "/" + k.name + day + ".log")
	if !filepath.IsAbs(file) {
		file = filepath.Clean(ProcessPath() + "/" + file)
	}
	path := filepath.Dir(file)
	if err = os.MkdirAll(path, 0777); err != nil {
		return err
	}
	k.logFile, err = os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		return err
	}
	if os.Stdout == nil {
		k.log = log.New(k.logFile, k.prefix, log.Ldate|log.Ltime)
	} else {
		k.wr = io.MultiWriter(os.Stdout, k.logFile)
		k.log = log.New(k.wr, k.prefix, log.Ldate|log.Ltime)
	}
	k.day = day
	return nil
}

func (k *KLog) Close() {
	if k.logFile != nil {
		k.logFile.Close()
		k.logFile = nil
		k.log = nil
		k.wr = nil
		k.day = ""
	}
}

func (k *KLog) Fatal(v ...interface{}) error {
	if err := k.init(); err != nil {
		return err
	}
	k.log.Fatal(v...)
	return nil
}

func (k *KLog) Fatalf(format string, v ...interface{}) error {
	if err := k.init(); err != nil {
		return err
	}
	k.log.Fatalf(format, v...)
	return nil
}

func (k *KLog) Fatalln(v ...interface{}) error {
	if err := k.init(); err != nil {
		return err
	}
	k.log.Fatalln(v...)
	return nil
}

func (k *KLog) Panic(v ...interface{}) error {
	if err := k.init(); err != nil {
		return err
	}
	k.log.Panic(v...)
	return nil
}

func (k *KLog) Panicf(format string, v ...interface{}) error {
	if err := k.init(); err != nil {
		return err
	}
	k.log.Panicf(format, v...)
	return nil
}

func (k *KLog) Panicln(v ...interface{}) error {
	if err := k.init(); err != nil {
		return err
	}
	k.log.Panicln(v...)
	return nil
}

func (k *KLog) Print(v ...interface{}) error {
	if err := k.init(); err != nil {
		return err
	}
	k.log.Print(v...)
	return nil
}

func (k *KLog) Printf(format string, v ...interface{}) error {
	if err := k.init(); err != nil {
		return err
	}
	k.log.Printf(format, v...)
	return nil
}

func (k *KLog) Println(v ...interface{}) error {
	if err := k.init(); err != nil {
		return err
	}
	k.log.Println(v...)
	return nil
}
