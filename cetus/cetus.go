package cetus

import (
	"encoding/json"

	"github.com/ipoluianov/ceta/sui/client"
)

func GetGool() string {
	cl := client.NewClient(client.MAINNET_URL)
	var options client.GetObjectShowOptions
	options.ShowContent = true
	obj, err := cl.GetObject("0x763f63cbada3a932c46972c6c6dcf1abd8a9a73331908a1d7ef24c2232d85520", options)
	if err != nil {
		return err.Error()
	}
	bs, _ := json.MarshalIndent(obj, "", "    ")
	return string(bs)
}
