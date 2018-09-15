# Decentralized Video Conferencing Demo

## Build

`cd $GOPATH/src/github.com/gangwn/dvc/`

`dep ensure`

`go build cmd/dvc.go`

`go build tests/client/dvc_proxy.go`

## Run demo

1. Install geth
2. Create test account 
`geth --datadir <testnet_dir> account new`
3. Initialize the balance for your new account in $GOPATH/src/github.com/gangwn/dvc/tests/geth/PrivateTestNetGenesis.json
4. Initialize private blockchain with the genesis block
`geth init --datadir "<testnet_dir>" $GOPATH/src/github.com/gangwn/dvc/tests/geth/PrivateTestNetGenesis.json`
5. Start a node use geth
`geth --identity "<node_name>" --rpc --rpcport "8545" --rpccorsdomain "*" --datadir "<testnet_dir>" --port "30303" --nodiscover --rpcapi "db,eth,net,web3" --networkid 1999 console`
6. Start to miner in the console
`miner.start()`
7.Deploy contract, record the address for ServiceManager <service_manager_addr>
`cd $GOPATH/src/github.com/gangwn/dvc/api/contract`
`truffle compile`
`truffle migrate`
8. Build dvc and dvc_proxy
9. Run dvc and dvc_proxy
`./dvc --config=./configs/dvc.json`
`./dvc_proxy --account=<eth_account> --keystore_dir=<keystore_dir> --contract_address=<service_manager_addr> --geth_rpc_path=<geth_rpc_path>`
10. Use the browser to open $GOPATH/src/github.com/gangwn/dvc/tests/demo/index.html