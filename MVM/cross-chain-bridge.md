# To build a evm-based cross-network bridge

## Deposit

1. A contract in the native network, user can call this contract to stake their token(ETH).

2. A contract in the target network, a client monitoring contract1, call contract to issue token(WETH).

3. An AMM(curve) for turning WETH to ETH in target network, transfer to user address.


## Withdraw

1. User call withdraw contract with token(ETH).

2. Withdraw contract call AMM(curve), turn ETH to WETH, call stake contract

3. A client mointoring stake contract, call contract1 to unstake token(ETH), transfer to the address user specified.
