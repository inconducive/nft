package api

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
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
	curlResults := *CurlRunner(CAs)
	log.Println(cr, "^^^^ vvvv ", curlResults)
}

func CurlRunner(ca []string) *curlResultSet {
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
		results = append(results, fmt.Sprintf("%s", c))
	}
	return &results
}
