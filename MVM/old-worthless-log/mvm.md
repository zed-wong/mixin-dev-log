# MVM

RegistryAddress: "0x65ccF8d1B92AfC6aF2915Cb61e72d93ACdD16556"

RegistryContractURL: https://testnet.mvmscan.com/address/0x65ccF8d1B92AfC6aF2915Cb61e72d93ACdD16556/transactions

Toolbox: http://metamask.test.mixinbots.com/

Official documention: https://mvm.dev

## Invoke with registry contract when error occurs

The registry would refund automaticlly

## Reserved transfer amount

1 satoshi (0.00000001)

## Extra too long?

See: https://github.com/liuzemei/mvm-mvm/blob/6333798d5c737cdcfb77cc64914deed4ebaad3d6/src/components/mvm/uniswap/swap.vue#L67

Add `uploadkey`(lowercase) inside `options` field with any vaule and try again.

## Script for generate mvm invoke payment code ?

See: https://github.com/MixinNetwork/bot-api-nodejs-client/blob/main/example/mvm.js

Or: http://metamask.test.mixinbots.com/ to generate extra.

## Amount correspondence between Mixin asset and MVM asset

Decimal: 100000000 (8 digits)

In solidity, amount type is: uint256 amount

1 CNB(mixin) = 100000000 (solidity uint256)
0.00000001(mixin) = 1 (solidity uint256)


## Since mvm contract functions are "public", is it possible for crackers to steal your money?


Use msg.sender as the only identifier for your users, and by far I suppose it's safe (I didn't find any vulnerabilities, it's logically correct).

## What happens after I scanned a code and paid?

When you as a mixin messenger user scanned a code that is gonna call registry contract, after you transfered the money to MVM nodes, the procedure would be like:
(-> == call)

Registry contract -> Your mvm contract address (mixin messenger user address)

Your mvm contract address -> Traget contract address (Address specified in transfer memo)

Traget contract address -> Your mvm contract address (Skip this step if no money was transfered) 

Your mvm contract address -> Registry contract

## Possible transaction situations

| Transfer amount        | Usage (call solidity func) | Reaction      | State             | Intro                                                        |
| ---------------------- | -------------------------- | ------------- | ----------------- | ------------------------------------------------------------ |
| 0.00000001 (1 satoshi) | Call non-transfer function | Refund anyway | Succeed or Failed | State depends on condition inside the function               |
| 0.00000001 (1 satoshi) | Call transfer function     | Refund anyway | Failed            | Unspported                                                   |
| other amount           | Call transfer function     | Depends       | Succeed           | Reaction depends on how function handles the money           |
| other amount           | Call transfer function     | Refund        | Failed            | Reaction depends on how function handles the money before failed. |
| other amount           | Call non-transfer function | Refund anyway | Succeed or Failed | State depends on condition inside the function.              |



## Any other problems?

Find https://github.com/liuzemei (id:30265) or https://github.com/jadeydi (id:493230) in mixin messenger 

Join group: https://mixin.one/codes/a91c865c-5de7-40c1-a130-f6c92ee89bd7
