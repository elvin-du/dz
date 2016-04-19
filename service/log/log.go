package log

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Init() {
}

func Debugln(v ...interface{}) {
	log.Println(v...)
}

func Infoln(v ...interface{}) {
	log.Println(v...)
}

func Errorln(v ...interface{}) {
	log.Println(v...)
}

func Warnln(v ...interface{}) {
	log.Println(v...)
}

func Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func Warnf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
