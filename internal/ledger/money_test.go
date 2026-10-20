package ledger

import "testing"

func TestValidCurrency(t *testing.T) {
	cases := map[string]bool{
		"INR":  true,
		"EUR":  true,
		"MXN":  true,
		"usd":  false, // must be upper-case
		"US":   false, // too short
		"INRR": false, // too long
		"US1":  false, // digits not allowed
		"":     false,
	}
	for in, want := range cases {
		if got := validCurrency(in); got != want {
			t.Errorf("validCurrency(%q) = %v, want %v", in, got, want)
		}
	}
}
