# ripego

[![Build Status](https://travis-ci.org/alxark/ripego.svg?branch=master)](https://travis-ci.org/alxark/ripego)

This package detects the coordination center for the IPv4 address and makes IP Whois lookup. Retrieve all informations about IP adress. Supported registries: ripe, arin, apnic, afrinic, lacnic

### Install

```bash
$ go get github.com/c1982/ripego
```

### Usage

```go
package main

import (
	"fmt"
	"log"
	"ripego"
)

func main() {

	w, err := ripego.IPLookup("178.18.196.250")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inetnum: " + w.Inetnum)
	fmt.Println("Desc: " + w.Descr)
}

//Output:
//Inetnum: 178.18.192.0 - 178.18.207.255
//Desc: Vargonen Teknoloji ve Bilisim Sanayi Ticaret Anonim Sirketi
```

### Contact

aspsrc@gmail.com

Oğuzhan
