# Self-hosted graph node

Guide: [graph-node/getting-started.md at master · graphprotocol/graph-node · GitHub](https://github.com/graphprotocol/graph-node/blob/master/docs/getting-started.md)

Hardware requirement: https://thegraph.academy/indexers/indexer-requirements/

1. Get a server with 4 vCPU, 8GB RAM and 1TB ROM at least. I choosed a Ubuntu server.

2. Install Rust 
   
   - `curl --proto '=https' --tlsv1.2 -sSf [https://sh.rustup.rs](https://sh.rustup.rs) | sh`

3. Download IPFS ([Command-line | IPFS Docs](https://docs.ipfs.tech/install/command-line/#official-distributions))
   
   - After downloaded, run
   
   - `ipfs init`
   
   - `ipfs daemon`

4. Install postgres [Install postgres on Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-20-04)
   
   - After downloaded, run
   
   - ```shell
     initdb -D .postgres
     pg_ctl -D .postgres -l logfile start
     createdb <POSTGRES_DB_NAME>
     ```

           In my case, the `initdb` was added to the $PATH, so I ran `find / -name initdb` and ran it with path. After this, I ran `pg_ctl` and it failed due to the postgresql is already running. So I just ran `createdb graph` directly.

4.  Run graph node
   
   - ```
     cargo run -p graph-node --release -- \
       --postgres-url postgresql://graph:PASSWORD@localhost:5432/graph \
       --ethereum-rpc mvm:https://geth.mvm.dev \
       --ipfs 127.0.0.1:5001 \
       --debug
     ```

5. Deploy a contract and write a subgraph, then deploy that subgraph to the graph node and test the API.


