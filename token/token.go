package token

type Token struct {
	Balances map[string]float64
	Stakes   map[string]float64
}

func New() *Token {
	return &Token{
		Balances: map[string]float64{"a": 100, "b": 100},
		Stakes:   map[string]float64{},
	}
}

func (t *Token) Transfer(from, to string, amount float64) {
	t.Balances[from] -= float64(amount)
	t.Balances[to] += float64(amount)
}

func (t *Token) Staking(from string, amount float64) {
	t.Balances[from] -= amount
	t.Stakes[from] += amount
}

func (t *Token) Mint(from string, amount float64) {
	t.Stakes[from] += amount
}

func (t *Token) Slash(from string, amount float64) {
	t.Stakes[from] -= amount
}

func (t *Token) BalanceOf(from string) (amount float64) {
	amount = t.Balances[from]
	return
}
