# To build a evm-based cross-network bridge

1. A contract in the native network, user can call this contract to stake their token(ETH).

2. A contract in the target network, monitoring contract1, issue token(WETH).

3. An AMM(curve) for turning WETH to ETH in target network, transfer to user address.
