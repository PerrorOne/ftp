# goftp #

[![Build Status](https://travis-ci.org/jlaffaye/ftp.svg?branch=master)](https://travis-ci.org/jlaffaye/ftp)
[![Coverage Status](https://coveralls.io/repos/jlaffaye/ftp/badge.svg?branch=master&service=github)](https://coveralls.io/github/jlaffaye/ftp?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/jlaffaye/ftp)](http://goreportcard.com/report/jlaffaye/ftp)

一个golang ftp包。该项目基于https://github.com/jlaffaye/ftp 修改完成

## 安装 ##

```
go get -u github.com/PerrorOne/ftp
```
## 使用 ##
```
package main

import (
	"time"
	"fmt"
	"github.com/PerrorOne/ftp"
)

func main() {
	c := ftp.NewFtp("127.0.0.1:8080", "toor", "235689", 1*time.Second)
	if err := c.Conn(); err != nil{
      fmt.Println(err)
      return
  }
	d, _:= c.List("/")
	for _, v:=range d{
		fmt.Println(v.Name, v.Type, v.Size)
	}
}
```

## 文档 ##

http://godoc.org/github.com/jlaffaye/ftp
