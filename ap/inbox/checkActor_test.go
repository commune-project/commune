package inbox_test

import (
	"testing"

	"github.com/commune-project/commune/ap/inbox"
)

func Test_CheckActor(t *testing.T) {
	data := map[string]interface{}{
		"id":    "https://m.hitorino.moe/p/1/activity",
		"actor": "https://m.hitorino.moe/users/misaka4e22",
	}
	data2 := map[string]interface{}{
		"id": "https://m.hitorino.moe/p/1/activity",
		"actor": map[string]interface{}{
			"id": "https://m.hitorino.moe/users/misaka4e22",
		},
	}
	dataDeny := map[string]interface{}{
		"id":    "https://limelight.moe/p/1/activity",
		"actor": "https://m.hitorino.moe/users/misaka4e22",
	}
	if err := inbox.CheckActor().Process(nil, data, nil); err != nil {
		t.Error(err)
	}
	if err := inbox.CheckActor().Process(nil, data2, nil); err != nil {
		t.Error(err)
	}
	if err := inbox.CheckActor().Process(nil, dataDeny, nil); err == nil {
		t.Error(err)
	}
}
