package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)


const {
	privateKey      = "77e7e7d4db6930c32590eb803d74265679c64212de31fbdb178eeb377f0c2525"
	contractAddress = "0x9bC2523031B11ed566F9499e25Ff91E04448eB4F"
	toAddress       = "0xC214f25Caa39A2Bb1BE3FF4BfAe9C79E5EbC148a" //这里我使用transfer方法作为案例，所以需要一个接收转账地址
}

func transfer(client *ethclient.Client, privateKey, toAddress, contract string) (string, error) {
		//从私钥推导出 公钥
		privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			fmt.Println("crypto.HexToECDSA error ,", err)
			return "", err
		}

		publicKey := crypto.Hex

		publicKey := privateKeyECDSA.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			fmt.Println("publicKeyECDSA error ,", err)
			return "", err
		}
		
		//从公钥推导出钱包地址
		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		fmt.Println("钱包地址：", fromAddress.Hex())
		var data []byte
		methodName := crypto.Keccak256([]byte("transfer(address,uint256)"))[:4]
		paddedToAddress := common.LeftPadBytes(common.HexToAddress(toAddress).Bytes(), 32)
		amount, _ := new(big.Int).SetString("100000000000000000000", 10)
		paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
		data = append(data, methodName...)
		data = append(data, paddedToAddress...)
		data = append(data, paddedAmount...)

}
