package vic

import (
	"log"

	"github.com/aziflaj/caboose/sarge"
)

func HandleRequest(store *KVStore, data string) string {
	// Step 1: parse RESP
	req, err := sarge.Deserialize(data)
	if err != nil {
		panic(err)
	}

	reqArr := req.([]string)

	// Step 2: Very Intricate AI (deeply nested if-else)
	command, args := reqArr[0], reqArr[1:]
	log.Println(command, args)
	res := Execute(store, command, args)

	// Step 3: respond with RESP
	// Step 4: ???
	// Step 5: Profit
	return res
}
