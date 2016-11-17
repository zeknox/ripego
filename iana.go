package ripego

import (
	"io/ioutil"
	"net"
	"strings"
	"time"
	"errors"
)

const (
	// IANAWhois server
	IANAWhois = "whois.iana.org:43"
)

/*
refer:        whois.ripe.net

inetnum:      213.0.0.0 - 213.255.255.255
organisation: RIPE NCC
status:       ALLOCATED

whois:        whois.ripe.net

changed:      1993-10
source:       IANA
*/

// WhoisData struct
type WhoisData struct {
	refer        string
	inetnum      string
	organisation string
	status       string
	whois        string
	changed      string
	source       string
	result       string
}

// Query function
func Query(ipaddr string) (result string, err error) {
	// TODO:
	//  - Validate IPv4 and IPv6 addresses
	// server = IANAWhois

	conn, err := net.DialTimeout("tcp", IANAWhois, time.Second * 30)
	if err != nil {
		return "", errors.New("failed to establish connection with whois server" + IANAWhois)
	}
	defer conn.Close()

	conn.Write([]byte(ipaddr + "\r\n"))
	var buffer []byte
	buffer, err = ioutil.ReadAll(conn)
	if err != nil {
		return "", errors.New("failed to read data from IANA server")
	}

	result = string(buffer[:])

	return
}

// Server function
func Server(whois string) (whoisserver string, whoisorg string, err error) {
	afnic := "whois.afrinic.net"
	apnic := "whois.apnic.net"
	arin := "whois.arin.net"
	lacnic := "whois.lacnic.net"
	ripe := "whois.ripe.net"

	switch {
	case strings.Contains(whois, afnic):
		return afnic, "afnic", nil
	case strings.Contains(whois, apnic):
		return apnic, "apnic", nil
	case strings.Contains(whois, arin):
		return arin, "arin", nil
	case strings.Contains(whois, lacnic):
		return lacnic, "lacnic", nil
	case strings.Contains(whois, ripe):
		return ripe, "ripe", nil
	}

	return "none", "none", errors.New("No such NIC found")
}
