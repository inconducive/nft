package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var (
	contractAddress = ""
	ca              = []string{
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
	response = map[string]interface{}{}
	results  = []string{}
)

const (
	APIkey     = "WtwzAuh31WD6CR1u8p9sMv46ShxLJKuw"
	webSockets = "wss://eth-mainnet.g.alchemy.com/v2/"
	uri        = "https://eth-mainnet.g.alchemy.com/v2/"
	// testuri    = uri + "0xcdb7c1a6fe7e112210ca548c214f656763e13533"
	rateLimit = 4 // requests/second
	cr        = "\n"
	c         = "curl"
)

type curlResultSet []string
type golibResultSet []string

/*
curl GET getFloorPrice example
https://eth-mainnet.g.alchemy.com/nft/v2/{apiKey}/getFloorPrice

curl --request GET \
--url 'https://eth-mainnet.g.alchemy.com/nft/v2/docs-demo/getFloorPrice?contractAddress=0xe785E82358879F061BC3dcAC6f0444462D4b5330' \
--header 'accept: application/json'

// curl --request GET --url "https://eth-mainnet.g.alchemy.com/nft/v2/WtwzAuh31WD6CR1u8p9sMv46ShxLJKuw/
// getFloorPrice?contractAddress=0xe785E82358879F061BC3dcAC6f0444462D4b5330" --header "accept: application/json"

tested_working_uri
curl --request GET --url "https://eth-mainnet.g.alchemy.com/nft/v2/WtwzAuh31WD6CR1u8p9sMv46ShxLJKuw/getFloorPrice?contractAddress=0xe785E82358879F061BC3dcAC6f0444462D4b5330" --header "accept: application/json"
*/

func Qmain() {
	/*
		for _, e := range ca {

			curls := `curl --request GET --url https://eth-mainnet.g.alchemy.com/nft/v2/` + APIkey +
				`/getFloorPrice?contractAddress=` + e + ` --header "accept: application/json"`
			log.Println(curls)
		}
	*/
	curlResults := *curlRunner(ca)
	log.Println(cr, "^^^^ vvvv ", curlResults)
}

func curlRunner(ca []string) *curlResultSet {
	results := curlResultSet{}
	for _, e := range ca {
		curls := `curl --request GET --url https://eth-mainnet.g.alchemy.com/nft/v2/` + APIkey +
			`/getFloorPrice?contractAddress=` + e + ` --header "accept: application/json"`
		// log.Println(curls)
		curlparts := strings.Split(curls, " ")
		curlparameters := curlparts[1:]

		c, err := exec.Command(curlparts[0], curlparameters...).Output()
		// c, err := exec.Command(curlparts[0], curlparameters...).CombinedOutput() // Output()
		if err != nil {
			log.Println(err.Error())
		}
		// fmt.Printf("%s", c)
		results = append(results, fmt.Sprintf("%s", c))
	}
	return &results
}
