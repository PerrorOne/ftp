package main

import (
	"github.com/PerrorOne/ftp"
	"time"
	"fmt"
)

func main() {
	f := ftp.NewFtp("127.0.0.1:8080", "toor", "235689", 1*time.Second)
	fmt.Println(f.Conn())
	d, _:= f.List("/")
	for _, v:=range d{
		fmt.Println(v.Name, v.Type, v.Size)
	}
}
