package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func main() {

	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))

	logrus.SetLevel(logrus.TraceLevel)

	logrus.SetReportCaller(true)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.WithField("name", "dj").Info("info msg")
	logrus.WithFields(logrus.Fields{
		"name": "Foo",
		"age":  18,
	}).Error("error msg")
	requestLogger := logrus.WithFields(logrus.Fields{
		"user_id": 10010,
		"ip":      "192.168.32.15",
	})
	requestLogger.Fatal("fatal msg")
	logrus.Panic("panic msg")

}
