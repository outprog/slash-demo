package schema

type Tx struct {
	Action string
	From   string
	To     string
	Amount float64
	Signer string
}
