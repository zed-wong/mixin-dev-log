# Deploy llamapay

## Contract part

1. Deploy the factory contract

2. Modify pay contract before deploying (Replace 'ERC20.SafeTransferFrom' to 'ERC20.transferFrom', replace all safe function to normal function)

3. Call factory contract with erc20 address to deploy pay contract

4. Approve pay contract and call the 'deposit' function of the pay contract.

5. If deposit succeed, then the contract deployment was succeed.

## Subgraph

1. git clone $GRAPH_URL

2. cd $GRAPH

3. yarn codegen && yarn build 

4. graph create --node https://graph.mvg.finance/deploy/ llamapay-subgraph-mvm-1

5. graph deploy --ipfs http://graph.mvg.finance:5001 --node https://graph.mvg.finance/deploy/ llamapay-subgraph-mvm-1

##  TokenList

0. See (https://github.com/zed-wong/mvm-tokenlist) for more details

1. Collect and generate the token list in such format.

## Interface

0. See (https://github.com/zed-wong/llamapay-interface-mvm/commit/7c3d4a406fc11c5b59eff74d72203f1a7a0b6693) for more details

1. Add MVM network

2. Modify gas price when calling contract

3. Add token list for mvm
