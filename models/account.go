package models

type Account struct {
	Type  string `json:"type"`
	Value struct {
		Address string `json:"address"`
		Coins   []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"coins"`
		PublicKey struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"public_key"`
		AccountNumber string `json:"account_number"`
		Sequence      string `json:"sequence"`
	} `json:"value"`
}
