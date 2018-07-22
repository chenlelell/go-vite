package handler

import (
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
	"math/big"
)
// pack the data for handler
type TokenInfo struct {
	Token *ledger.Token
	TotalAmount *big.Int
}

type UnconfirmedAccount struct {
	AccountAddress *types.Address
	TotalNumber *big.Int
	TokenInfoList []*TokenInfo

}

func (ac *AccountChain) GetUnconfirmedAccountMeta (addr *types.Address) (*ledger.UnconfirmedMeta, error) {
	return ac.uAccess.GetUnconfirmedAccountMeta(addr)
}

//func (ac *AccountChain) GetUnconfirmedBlocks (index int, num int, count int, addr *types.Address, tokenId *types.TokenTypeId) ([]*ledger.AccountBlock, error) {
//	acMeta, err := ac.aAccess.GetAccountMeta(addr)
//	if err != nil {
//		return nil, err
//	}
//	return ac.uAccess.GetUnconfirmedBlocks(index, num, count, acMeta.AccountId, tokenId)
//}

func (ac *AccountChain) GetUnconfirmedHashs (index int, num int, count int, addr *types.Address, tokenId *types.TokenTypeId) ([]*types.Hash, error) {
	acMeta, err := ac.aAccess.GetAccountMeta(addr)
	if err != nil {
		return nil, err
	}
	return ac.uAccess.GetUnconfirmedHashs(index, num, count, acMeta.AccountId, tokenId)
}

func (ac *AccountChain) GetUnconfirmedAccount (addr *types.Address) (*UnconfirmedAccount, error) {
	unconfirmedMeta, err := ac.GetUnconfirmedAccountMeta(addr)
	if err != nil {
		return nil, err
	}
	var tokenInfoList []*TokenInfo
	for _, ti := range unconfirmedMeta.TokenInfoList {
		token, tkErr := ac.tAccess.GetByTokenId(ti.TokenId)
		if tkErr != nil {
			return nil, tkErr
		}
		tokenInfo := &TokenInfo{
			Token:       token,
			TotalAmount: ti.TotalAmount,
		}

		tokenInfoList = append(tokenInfoList, tokenInfo)
	}
	var UnconfirmedAccount = &UnconfirmedAccount{
		AccountAddress: addr,
		TotalNumber:    unconfirmedMeta.TotalNumber,
		TokenInfoList:  tokenInfoList,
	}
	return UnconfirmedAccount, nil
}

func Listener (addr *types.Address, c *chan int32) error {
	return nil
}

func RemoveListener (c *chan int) error {
	close(*c)
	return nil
}