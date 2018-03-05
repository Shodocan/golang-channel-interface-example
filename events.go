package example

type Event interface{}

type Deposit struct {
	Account string
	Amount  float64
}

type Withdrawn struct {
	Account string
	Amount  float64
}

type Transfer struct {
	From   string
	To     string
	Amount float64
}
