# Deploy curve on MVM

1. Install brownie

2. git clone https://github.com/curvefi/curve-contract

3. edit script/deploy.py

```
# 1. In accounts.add('XXX'), Replace XXX with your account's private key

# 2. In POOL_NAME, Replace with the name of pool you want to deploy. In POOL_OWNER,POOL_OWNER, Replace it with your address.

# 3. In _tx_params(), Replace GasNowScalingStrategy() with number, (e.g. 1000000000) (1Gwei)

# 4. In token_deployer.deploy(token_args["name"], token_args["symbol"], _tx_params()), 
# Replace it with token = token_deployer.deploy(token_args["name"], token_args["symbol"], '18', '100000000', _tx_params())
# ('18' is the decimal, '100000000' is total supply of the token.)
```

4. add network: `$ brownie networks add MVM mvm-main name="Mainnet" chainid=73927 host=https://geth.mvm.dev/ explorer="https://scan.mvm.dev/"`

5. `$ brownie run deploy --network mvm-main`

Then it's deployed. Then call the funtion in console


6. `$ brownie console`

7. `>>> accounts`, if `[]` then `accounts.add('REPLACE WITH YOUR PRIVATE KEY')`

8. `>>> network.show_active()`, if `'development'` then `network.disconnect()`, `network.connect('mvm-main')`

9. `>>> c = Contract('REPLACE WITH YOUR STABLE SWAP CONTRACT ADDRESS ON MVM')`

10. `>>> c.A()`, returns `10`

11. `>>> c.add_liquidity([1,1,1], 1, {'from':accounts[0], 'gas_limit':210000})`

`ValueError: Execution reverted during call: 'execution reverted'. This transaction will likely revert. If you wish to broadcast, include `allow_revert:True` as a transaction parameter.`

I ended up here. Have no idea how to make it work TAT.

