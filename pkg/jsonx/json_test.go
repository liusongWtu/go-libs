package jsonx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSortedKeyJsonShallow(t *testing.T) {

	data := `
	{
		"sign":"test",
		"payment_id":5077125051,
		"no":"111",
		"pay_amount":155.38559757,
		"actually_paid":0,
		"order_id":"2",
		"order_description":"Apple Macbook Pro 2019 x 1",
		"purchase_id":"6084744717",
		"created_at":"2021-04-12T14:22:54.942Z",
		"outcome_amount":1131.7812095
	}`

	want := `{"actually_paid":0,"created_at":"2021-04-12T14:22:54.942Z","no":"111","order_description":"Apple Macbook Pro 2019 x 1","order_id":"2","outcome_amount":1131.7812095,"pay_amount":155.38559757,"payment_id":5077125051,"purchase_id":"6084744717"}`

	ignoreKeys := map[string]struct{}{
		"sign": {},
	}
	target, err := ToSortedKeyJsonShallow(data, ignoreKeys)
	assert.NoError(t, err)
	assert.True(t, target == want)

}
