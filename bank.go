package example

type Account struct {
	Money float64
}

type Bank struct {
	Accounts map[string]Account
}
