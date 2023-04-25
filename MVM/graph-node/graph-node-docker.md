Follow this guide

https://medium.com/coinmonks/deploy-subgraphs-to-any-evm-aaaccc3559f


1. git clone GRAPH-NODE

2. install docker
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-22-04
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04

3. edit graph-node/docker/docker-compose.yml

Keep mainnet, replace RPC with yours

4. ./setup.sh

5. docker-compose up

The setup of the node is done. Then try to deploy a subgraph to it.

1. git clone YOUR-GRAPH

2. edit contract address in subgraph.yaml, network should be `mainnet`

3. graph create --node YOUR-NODE-ADDRESS

4. graph deploy --ipfs YOUR-IPFS-ADDRESS:5001 --node YOUR-NODE-ADDRESS:8020 YOUR-GRAPH-NAME

(Replace all Uppercase with your conf)
