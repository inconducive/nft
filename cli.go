package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

/*
var (
	ca = []string{
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
	response        = map[string]interface{}{}
	contractAddress = ""
)

const (
	APIkey     = "WtwzAuh31WD6CR1u8p9sMv46ShxLJKuw"
	webSockets = "wss://eth-mainnet.g.alchemy.com/v2/"
	uri        = "https://eth-mainnet.g.alchemy.com/v2/"
	// testuri    = uri + "0xcdb7c1a6fe7e112210ca548c214f656763e13533"
	rateLimit = 4 // requests/second
	cr        = "\n"
)
*/

func Tmain() {
	// Run("A> ", os.Stdin, os.Stdout)
	// os.Exit(0)

	url := "https://eth-mainnet.g.alchemy.com/nft/v2/" + APIkey + "/getFloorPrice?contractAddress=" + contractAddress
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	response, err := http.DefaultClient.Do(req)
	log.Println("Ignore below error :", cr, err, "============", cr, "<>")
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	log.Println(response)
	log.Println(string(body))

	// caRes := exec.Command(curler)

}

func Run(prompt string, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		log.Println(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" || line == "" || line == "quit" {
			os.Exit(0)
		}
		parts := strings.Split(line, "")
		if len(parts) == 0 {
			os.Exit(0)
		}

		cmd := strings.TrimSpace(parts[0])
		args := parts[1:]
		runner := exec.Command(cmd, args...)
		cmdResult, err := runner.CombinedOutput()

		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		out.Write(cmdResult)
	}
}

func Runcli(prompt string, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		log.Println(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" || line == "" || line == "quit" {
			os.Exit(0)
		}
		parts := strings.Split(line, "")
		if len(parts) == 0 {
			os.Exit(0)
		}

		cmd := strings.TrimSpace(parts[0])
		args := parts[1:]
		runner := exec.Command(cmd, args...)
		cmdResult, err := runner.CombinedOutput()

		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		out.Write(cmdResult)
	}
}
