package main

import (
	"io"
	"os"
	"os/user"
	"fmt"
	"gopkg.in/gomail.v2"
	"github.com/BurntSushi/toml"
)

type Config struct {
	From string
	To string
	MailServer string
	Port int
	Username string
	Password string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: meiwaku <spam mail filename>")
		return
	}

	var conf Config
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	if _, err := toml.DecodeFile(usr.HomeDir + "/.config/meiwaku.toml", &conf); err != nil {
		panic(err)
	}

	d := gomail.NewDialer(conf.MailServer, conf.Port, conf.Username, conf.Password)

	for _, spamfilename := range os.Args[1:] {
		m := gomail.NewMessage()
		m.SetHeader("From", conf.From)
		m.SetHeader("To" , conf.To)
		m.SetHeader("Subject", "Forwarded Message")
		m.SetBody("text/plain", "Forwarded Message\r\n")
		f := func(w io.Writer) error {
			h, err := os.Open(spamfilename)
			if err != nil {
				return err
			}
			if _, err := io.Copy(w, h); err != nil {
				h.Close()
				return err
			}
			return h.Close()
		}
		m.AddAlternativeWriter("message/rfc822", f, gomail.SetPartEncoding(gomail.Unencoded))

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}
