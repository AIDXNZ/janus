package transformer

import (
	"encoding/json"
	"testing"
)

func TestProtocolVersionReturnsHardcodedValue(t *testing.T) {
	//preparing the request
	requestParams := []json.RawMessage{} //eth_protocolVersion has no params
	request, err := prepareEthRPCRequest(1, requestParams)
	if err != nil {
		t.Fatal(err)
	}

	proxyEth := ETHProtocolVersion{}
	got, err := proxyEth.Request(request)
	if err != nil {
		t.Fatal(err)
	}

	expected := "0x41"

	if got != expected {
		t.Errorf(
			"error\nwant: %s\ngot: '%v'",
			expected,
			got,
		)
	}
}
