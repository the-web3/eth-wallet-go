package main

import (
	"fmt"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/the-web3/eth-wallet-go/addresss"
)

func Foo() *int {
	x := 42
	return &x // x 逃逸到堆上
}

func ArraySize() {
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("元素: %d, 地址: %p\n", arr[i], &arr[i])
	}
	fmt.Printf("数组占用的内存大小: %d bytes\n", unsafe.Sizeof(arr))
}

func main01() {
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

func main02() {
	ArraySize()
}

func main03() {
	s := make([]byte, 200)
	ptr := unsafe.Pointer(&s[0])
	fmt.Println(ptr)
}
