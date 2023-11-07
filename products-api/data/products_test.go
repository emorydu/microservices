package data

import "testing"

func TestProduct_Validate(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 1,
		SKU:   "abc-abc-abc",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
