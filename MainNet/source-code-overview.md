# Mixin Mainnet source code overview

1. What is required to add a domain

- 1. Find the commit they made, and compare what they added.

- Answer 1:


Case starcoin (https://github.com/MixinNetwork/mixin/commit/8dc38c1fd59005240ddbe1413a61067667bdb111)

common/asset.go:	add case
common/deposit.go:	add case
common/withdrawal.go:	add case
domains/NAME/validation.go	add file

This file might requires to implement
```
var (
	StarcoinAssetKey  string
	StarcoinChainBase string
	StarcoinChainId   crypto.Hash
)
func init()
func VerifyAssetKey(assetKey string)
func VerifyAddress(address string)
func VerifyTransactionHash(hash string)
func GenerateAssetId(assetKey string)
```


Case terra (https://github.com/MixinNetwork/mixin/commit/01c080692fd1d2ba9b674b5ed19e86e0c72a7a82)


