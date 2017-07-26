package ripego

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

type Org struct {
	Code string `xml:"handle,attr"`
	Name string `xml:"name,attr"`
}

type Net struct {
	Name         string `xml:"name"`
	EndAddress   string `xml:"endAddress"`
	StartAddress string `xml:"startAddress"`
	OrgInfo      Org    `xml:"orgRef"`
	LastModified string `xml:"updateDate"`
	Created      string `xml:"registrationDate"`
}

func ArinCheck(search string, whoisServer string) (*WhoisInfo, error) {

	resp, err := http.Get("https://" + whoisServer + "/rest/ip/" + search)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed to parse reply body, got: " + err.Error())
	}

	return parseArinReply(body)
}

func parseArinReply(xmlreply []byte) (*WhoisInfo, error) {
	v := Net{}
	err := xml.Unmarshal(xmlreply, &v)
	if err != nil {
		return nil, err
	}

	return &WhoisInfo{
		Noc:          "ARIN",
		Inetnum:      v.StartAddress + " - " + v.EndAddress,
		Netname:      v.Name,
		Organization: v.OrgInfo.Name,
		Created:      v.Created,
		LastModified: v.LastModified,
	}, nil
}
