package vm

import (
	"github.com/outprog/slash-demo/ar"
	"github.com/outprog/slash-demo/token"
	"github.com/outprog/slash-demo/vm/schema"
)

type VM struct {
	token *token.Token
	ar    *ar.AR
	owner string
}

func New(owner string, ar *ar.AR) *VM {
	return &VM{token.New(), ar, owner}
}

func (v *VM) Exec(tx schema.Tx, dryrun bool) {
	switch tx.Action {
	case "transfer":
		v.token.Transfer(tx.From, tx.To, tx.Amount)
	case "staking":
		v.token.Staking(tx.From, tx.Amount)
	case "qryBalance":
		defer func() {
			// send msg to ar
			if !dryrun {
				reply := schema.Tx{
					Action: "replyBalance",
					From:   tx.From,
					Amount: v.token.BalanceOf(tx.From),
					Signer: v.owner,
				}
				v.ar.Txs = append(v.ar.Txs, reply)
				v.Exec(reply, true)
			}
		}()
	case "replyBalance":
		if v.token.Balances[tx.From] == tx.Amount {
			v.token.Mint(tx.Signer, 1.0)
		} else {
			v.token.Slash(tx.Signer, 1.0)
		}
	}
	if !dryrun {
		v.ar.Txs = append(v.ar.Txs, tx)
	}
}

func (v *VM) Recover() {
	for _, tx := range v.ar.Txs {
		v.Exec(tx, true)
	}
}
