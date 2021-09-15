package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)




func formatAddress(_address string) common.Address{
	// 使用以太坊地址，要先转成common.Address类型
	address := common.HexToAddress(_address)
	//fmt.Println(address.Hex())
	//fmt.Println(address.Hash().Hex())
	//fmt.Println(address.Bytes())
	return address
}

func getBalance(client *ethclient.Client, _address string) {
	address := formatAddress(_address)

	client.BalanceAt(context.Background(), address, nil)
}

func main()  {
	ipcUrl := "/Users/linsiting/Library/Ethereum/geth.ipc"
	client, _ := ethclient.Dial(ipcUrl)

	fmt.Println("we have a connection")
	_ = client

	account := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
	formatAddress(account)

}
