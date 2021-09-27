package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
)




// 将地址转换为common.Address类型
func formatAddress(_address string) common.Address{
	address := common.HexToAddress(_address)

	fmt.Println(address) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F (地址在evm中是160bits = 20 bytes = 40 hex)
	fmt.Println(address.Hash().Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes()) // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111] - 20bytes
	return address
}

// 查询账户可用余额
func getBalance(client *ethclient.Client, _address string, blockNumber *int64) {
	var _blockNumber *big.Int;
	if blockNumber != nil {
		_blockNumber = big.NewInt(*blockNumber)
	}

	address := formatAddress(_address)

	// 第三个参数- [区块号] - nil 表示获取最新区块
	balance, err := client.BalanceAt(context.Background(), address, _blockNumber)
	if err != nil {
		return
	}

	fmt.Println(balance) // 100000000000000000000

	// 余额换算
	t, _ := new(big.Float).SetString(balance.String())
	balanceWITHETH := new(big.Float).Quo(t, big.NewFloat(math.Pow10(18)))
	fmt.Println(balanceWITHETH) // 100
}

// 查询用户可用余额（包含被锁定为作为正在上链的交易的余额）
func getPendingBalance(client * ethclient.Client, account string) {

	_address := formatAddress(account)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), _address)
	if err != nil {
		return
	}

	fmt.Println(pendingBalance) // 100000000000000000000

}

func main()  {
	//ipcUrl := "/Users/linsiting/Library/Ethereum/geth.ipc"
	ipcUrl := "http://localhost:8545"
	client, _ := ethclient.Dial(ipcUrl)

	fmt.Println("we have a connection")
	_ = client

	account := "0x0D73960e65740983e2f4f5c969f9c2A076910366"

	//formatAddress(account)

	//getBalance(client, account, nil)

	getPendingBalance(client, account)

}
