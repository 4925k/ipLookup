package dnslookup

import (
	"encoding/json"
	"fmt"
	"io"
	"lookup/flatten"
	"net"
	"net/http"
	"strings"
)

func DnsLookup(ip string) (map[string]string, error) {
	//reverse given ip and add .in-addr.arpa for reverse lookup
	ipadd := net.ParseIP(ip)
	if ipadd.To4() != nil {
		temp := strings.Split(ip, ".")
		ip = fmt.Sprintf("%s.%s.%s.%s.in-addr.arpa", temp[3], temp[2], temp[1], temp[0])
	}
	response := make(map[string]interface{})
	//create request
	//the request is sent to google public dns resolver
	request := fmt.Sprintf("https://dns.google/resolve?name=%s&type=PTR", ip)
	resp, err := http.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//unmarshal body in to response struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	//flatten struct
	flat, err := flatten.Flatten(response, "", flatten.DotStyle)
	if err != nil {
		return nil, err
	}

	return flat, nil
}
