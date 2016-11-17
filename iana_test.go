package ripego

import (
	"testing"
)

func TestIanaQueryAndServer(t *testing.T) {
	whois, err := Query("8.8.8.8")
	if err != nil {
		t.Fatal("Failed to query IANA server")
	}

	t.Logf("IANA reply: \n%s", whois)

	nic_url := ""
	nic_url, _, err = Server(whois)
	if err != nil {
		t.Fatal("Failed to detect IP whois server")
	}

	t.Logf("Detected server %s", nic_url)
}
