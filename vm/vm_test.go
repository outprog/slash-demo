package vm

import (
	"fmt"
	"testing"

	"github.com/outprog/slash-demo/ar"
	"github.com/outprog/slash-demo/vm/schema"
)

func TestTransfer(t *testing.T) {
	immutable := &ar.AR{
		Txs: []schema.Tx{},
	}
	node1 := New("a", immutable)

	fmt.Println("node1, init balances")
	fmt.Println("a", node1.token.Balances["a"])
	fmt.Println("b", node1.token.Balances["b"])

	node1.Exec(schema.Tx{"transfer", "a", "b", 1.0, "", "a"}, false)

	fmt.Println("node1, a transfer 1.0 to b")
	fmt.Println("a", node1.token.Balances["a"])
	fmt.Println("b", node1.token.Balances["b"])

	///// node recover
	node2 := New("b", immutable)
	node2.Recover()
	fmt.Println("node2 status")
	fmt.Println("a", node2.token.Balances["a"])
	fmt.Println("b", node2.token.Balances["b"])
}

func TestInvlidNode(t *testing.T) {
	immutable := &ar.AR{
		Txs: []schema.Tx{},
	}
	node1 := New("a", immutable)

	fmt.Println("node1, init balances")
	fmt.Println("a", node1.token.Balances["a"])
	fmt.Println("b", node1.token.Balances["b"])
	fmt.Println("a stake", node1.token.Stakes["a"])
	fmt.Println("a stake", node1.token.Stakes["a"])

	node1.Exec(schema.Tx{"staking", "a", "", 50, "", "a"}, false)
	// b not stake
	// node1.Exec(schema.Tx{"staking", "b", "", 50, "b"}, false)
	fmt.Println("node1, stake")
	fmt.Println("a", node1.token.Balances["a"])
	fmt.Println("b", node1.token.Balances["b"])
	fmt.Println("a stake", node1.token.Stakes["a"])
	fmt.Println("b stake", node1.token.Stakes["b"])

	node1.Exec(schema.Tx{"qryBalance", "b", "", 0, "", "b"}, false)
	node1.Exec(schema.Tx{"qryBalance", "b", "", 0, "", "a"}, false)
	fmt.Println("node1, qry balance")
	fmt.Println("a", node1.token.Balances["a"])
	fmt.Println("b", node1.token.Balances["b"])
	fmt.Println("a stake", node1.token.Stakes["a"])
	fmt.Println("b stake", node1.token.Stakes["b"])
}

func TestSlash(t *testing.T) {
	immutable := &ar.AR{
		Txs: []schema.Tx{},
	}
	node1 := New("a", immutable)

	fmt.Println("node1, init balances")
	fmt.Println("a", node1.token.Balances["a"])
	fmt.Println("b", node1.token.Balances["b"])
	fmt.Println("a stake", node1.token.Stakes["a"])
	fmt.Println("b stake", node1.token.Stakes["b"])

	node1.Exec(schema.Tx{"staking", "a", "", 50, "", "a"}, false)
	node1.Exec(schema.Tx{"staking", "b", "", 50, "", "b"}, false)

	fmt.Println("node1, stake")
	fmt.Println("a", node1.token.Balances["a"])
	fmt.Println("b", node1.token.Balances["b"])
	fmt.Println("a stake", node1.token.Stakes["a"])
	fmt.Println("b stake", node1.token.Stakes["b"])

	// qry blance
	node1.Exec(schema.Tx{"qryBalance", "b", "", 0, "", "b"}, false)
	fmt.Println("node1, qry balance")
	fmt.Println("a", node1.token.Balances["a"])
	fmt.Println("b", node1.token.Balances["b"])
	fmt.Println("a stake", node1.token.Stakes["a"])
	fmt.Println("b stake", node1.token.Stakes["b"])

	// fake balance
	node1.ar.Txs = append(node1.ar.Txs, schema.Tx{"replyBalance", "b", "", 49, "", "a"})
	node1.ar.Txs = append(node1.ar.Txs, schema.Tx{"replyBalance", "b", "", 49, "", "a"})
	node1.ar.Txs = append(node1.ar.Txs, schema.Tx{"replyBalance", "b", "", 49, "", "a"})

	///// node recover
	node2 := New("b", immutable)
	node2.Recover()
	fmt.Println("node2 status")
	fmt.Println("a", node2.token.Balances["a"])
	fmt.Println("b", node2.token.Balances["b"])
	fmt.Println("a stake", node2.token.Stakes["a"])
	fmt.Println("b stake", node2.token.Stakes["b"])

	// print immuable tx list
	fmt.Println()
	for _, tx := range immutable.Txs {
		fmt.Printf("%+v\n", tx)
	}
}
