package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/the-web3/eth-wallet-go/addresss"
)

func main() {

	addrInfo, err := addresss.CreateAddressFromPrivateKey()
	if err != nil {
		fmt.Println("create address from private fail", "err", err)
		return
	}
	common.HexToHash("sss")
	fmt.Println("address=", addrInfo.PrivateKey)
	fmt.Println("address=", addrInfo.PublicKey)
	fmt.Println("address=", addrInfo.Address)
}
