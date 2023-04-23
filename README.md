# NFT
This is the get floor price API called for the task assigned by Ahmed shah- related NFT GET querying

go run main.go and Code will execute

## Ajay's comments before you run codeRun main.go to see the results on CLI / terminal screen. [tested on Win, shouldn't have issues on Linux shell/ Bash]
1. getfloorprice.go is the main API file with all essential functionality being called in main.go
2. curl.go is the CLI based CURL functionality for getFloorPrice assimilated into Go array.
3. get.go is the Go inbuilt net/http lib used to assimilate the results into Go array.
4. the time rate module and limiter concepts used are for advanced rate limits with concurrent manipulated code execution.
5. Since this task was only about limiting to 4 requests/ sec to avoid IP blacklisting or blocking, I kept it simple with time.Sleep.
