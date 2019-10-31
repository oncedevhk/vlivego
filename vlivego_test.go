package vlivego_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/keneuming/vlivego"
)

// 7.664s
func TestClient(t *testing.T) {
	seq := vlivego.GetChannelSeq("EDBF")
	client := vlivego.NewClient(seq)
	client.RefreshAll()
	b, _ := json.MarshalIndent(client, "", "	")
	fmt.Println(string(b))
}
