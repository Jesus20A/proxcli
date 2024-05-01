package request

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"proxcli/config"
)

// Make the different requests to the Proxmox API
func NewRequest(url, method string) ([]byte, int) {
	config := config.InitConfig()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	auth := fmt.Sprintf("PVEAPIToken=%s@%s!%s=%s", config["user"], config["realm"], config["tokenid"], config["token"])
	req.Header.Add("Authorization", auth)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body, res.StatusCode
}
