package main

import (
	srslog "github.com/artkescha/syslog"
	"log"
)

func main() {
	//w, err := srslog.DialWithTLSCertPath("tcp+tls", "127.0.0.1:514", srslog.LOG_ERR, "testtag", "/path/to/servercert.pem")
	w, err := srslog.Dial("udp", "127.0.0.1:514", srslog.LOG_ERR, "testtag")
	if err != nil {
		log.Fatal("failed to connect to syslog:", err)
	}
	defer w.Close()
	w.SetFormatter(srslog.RFC5424Formatter)

	w.Alert("this is an alert")
	w.Crit("this is critical")
	w.Err("this is an error")
	w.Warning("this is a warning")
	w.Notice("this is a notice")
	w.Info("this is info")
	w.Debug("this is debug")

	w.Write([]byte("Hello world!"))
}
