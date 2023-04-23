package api

import (
	"io"
	"log"
	"net/http"
)

var (
	CAs = []string{
		"0xf87e31492faf9a91b02ee0deaad50d51d56d5d4d",
		"0xff9c1b15b16263c61d017ee9f65c50e4ae0113d7",
		"0x273f7f8e6489682df756151f5525576e322d51a3",
		"0x1a92f7381b9f03921564a437210bb9396471050c",
		"0x219b8ab790decc32444a6600971c7c3718252539",
		"0xd4e4078ca3495de5b1d4db434bebc5a986197782",
		"0x6afcdefce0e8d1c774789b21c11e2bd5d4509a16",
		"0x97597002980134bea46250aa0510c9b90d87a587",
		"0x236672ed575e1e479b8e101aeeb920f32361f6f9",
		"0x90ca8a3eb2574f937f514749ce619fdcca187d45",
		"0x521f9c7505005cfa19a8e5786a9c3c9c9f5e6f42",
		"0x63c0691d05f441f42915ca6ca0a6f60d8ce148cd",
		"0x454cbc099079dc38b145e37e982e524af3279c44",
		"0x6d4530149e5b4483d2f7e60449c02570531a0751",
		"0xcdb7c1a6fe7e112210ca548c214f656763e13533",
	}
	// response = map[string]interface{}{}
	Results = []string{}
)

const (
	APIkey     = "WtwzAuh31WD6CR1u8p9sMv46ShxLJKuw"
	WebSockets = "wss://eth-mainnet.g.alchemy.com/v2/"
	// url        = "https://eth-mainnet.g.alchemy.com/v2/"
	// testuri    = uri + "0xcdb7c1a6fe7e112210ca548c214f656763e13533"
	RateLimit = 4 // requests/second
	cr, tab   = "\n", "\t"
)

func GetFloorPrice(apikey, contractAddress string) string {
	uri := `https://eth-mainnet.g.alchemy.com/nft/v2/` + apikey + `/getFloorPrice?contractAddress=` + contractAddress
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("accept", "application/json")
	res, er := http.DefaultClient.Do(req)
	if er != nil {
		log.Fatal(er)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return string(body)
}

func GetFloorPrices(apikey string, allCAs []string) *[]string {
	for _, v := range allCAs {
		uri := `https://eth-mainnet.g.alchemy.com/nft/v2/` + apikey + `/getFloorPrice?contractAddress=` + v
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Add("accept", "application/json")
		res, er := http.DefaultClient.Do(req)
		if er != nil {
			log.Fatal(er)
		}
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		Results = append(Results, string(body))
	}
	return &Results
}

func If(condition bool, truth, falsity interface{}) interface{} {
	if condition {
		return truth
	}
	return falsity
}
