package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/infobloxopen/atlas-contacts-app/pkg/pb"
)

const (
	endpoint = "http://localhost:8080/v1/networks"
	token    = "Token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEedA7dGULI3e"
)

func TestCreateNetwork_REST(t *testing.T) {

	network := pb.Network{
		Addr: "127.0.0.1",
		Cidr: 24,
		Fixed: []*pb.IPv4{
			&pb.IPv4{
				Addr: "192.168.0.0",
			},
		},
		Comment: "Localhost",
	}

	jsonNetwork, err := json.Marshal(network)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonNetwork))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatal("Incorrect response code", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	jsonResp := json.Unmarshal(body, &pb.Network{})
	fmt.Println(jsonResp)
}
