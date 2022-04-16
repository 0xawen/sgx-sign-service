package main

import (
	"log"
	crypto "sgx-sign-service/crypro"
)

var (
	//Parameters:
	//   - `language`：1中文，2英文。
	//   - `strength`：1弱（12个助记词），2中（18个助记词），3强（24个助记词）。
	language = 2
	strength = uint8(1)
)

// 签名服务接口-方便以后拓展
type SGXSignServe interface {
	// 创建账号(返回用户的地址/特定表示符号)
	CreateAccount() string
	// 签名
	Sign(date []byte, account string) ([]byte, error)
}

// 超级链账号
type XuperChainAccount struct {
	// 合约账号
	contractAccount string
	// 地址
	Address string
	// 私钥
	PrivateKey string
	// 公钥
	PublicKey string
	// 助记词
	Mnemonic string
}

func NewXuperchainAccount() *XuperChainAccount {
	return &XuperChainAccount{}
}

// 创建账号
func (x *XuperChainAccount) CreateAccount() string {
	cli := crypto.GetCryptoClient()
	ecdsaAccount, err := cli.CreateNewAccountWithMnemonic(language, strength)
	if err != nil {
		log.Printf("CreateAccount CreateNewAccountWithMnemonic err: %v", err)
		return "CreateAccount CreateNewAccountWithMnemonic err"
	}

	account := &XuperChainAccount{
		Address:    ecdsaAccount.Address,
		PublicKey:  ecdsaAccount.JsonPublicKey,
		PrivateKey: ecdsaAccount.JsonPrivateKey,
		Mnemonic:   ecdsaAccount.Mnemonic,
	}
	// 持久化
	GDB.Add(account.Address, account.Mnemonic)

	return account.Address
}

// 签名
func (x *XuperChainAccount) Sign(address string, data []byte) ([]byte, error) {
	// todo
	// 从数据库中读取account 的信息
	mnemonic, err := GDB.Query(address)
	if err != nil {
		return nil, err
	}
	// 恢复
	cryptoClient := crypto.GetCryptoClient()
	ecdsaAccount, err := cryptoClient.RetrieveAccountByMnemonic(mnemonic, language)
	if err != nil {
		return nil, err
	}
	// 签名
	privateKey, err := cryptoClient.GetEcdsaPrivateKeyFromJsonStr(ecdsaAccount.JsonPrivateKey)
	if err != nil {
		return nil, err
	}
	sign, err := cryptoClient.SignECDSA(privateKey, data)
	if err != nil {
		return nil, err
	}
	return sign, nil
}
