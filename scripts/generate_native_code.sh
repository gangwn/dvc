#!/bin/bash

# Generate contract code via below commands

cur_dir=$(cd `dirname $0` && pwd)
root_dir=$(cd $cur_dir/.. && pwd)

contract_root_dir=$root_dir/api/contract
source_dir=$contract_root_dir/contracts
dest_dir=$root_dir/pkg/eth/contracts

solc --overwrite --abi openzeppelin-solidity=$contract_root_dir/node_modules/openzeppelin-solidity $source_dir/ServiceManager.sol -o $root_dir/build/contracts
solc --overwrite --abi openzeppelin-solidity=$contract_root_dir/node_modules/openzeppelin-solidity $source_dir/ConferenceService.sol -o $root_dir/build/contracts
solc --overwrite --abi openzeppelin-solidity=$contract_root_dir/node_modules/openzeppelin-solidity $source_dir/FixedSupplyToken.sol -o $root_dir/build/contracts
solc --overwrite --abi openzeppelin-solidity=$contract_root_dir/node_modules/openzeppelin-solidity LibCLL=$contract_root_dir/node_modules/LibCLL $source_dir/CCSService.sol -o $root_dir/build/contracts
solc --overwrite --abi openzeppelin-solidity=$contract_root_dir/node_modules/openzeppelin-solidity $source_dir/IService.sol -o $root_dir/build/contracts
solc --overwrite --abi openzeppelin-solidity=$contract_root_dir/node_modules/openzeppelin-solidity $source_dir/Utils.sol -o $root_dir/build/contracts

abigen --abi $root_dir/build/contracts/ServiceManager.abi --pkg contracts --type ServiceManager --out $dest_dir/service_manager.go
abigen --abi $root_dir/build/contracts/ConferenceService.abi --pkg contracts --type ConferenceService --out $dest_dir/conference_service.go
abigen --abi $root_dir/build/contracts/CCSService.abi --pkg contracts --type CCSService --out $dest_dir/ccs_service.go
abigen --abi $root_dir/build/contracts/FixedSupplyToken.abi --pkg contracts --type FixedSupplyToken --out $dest_dir/fixed_supply_token.go

#abigen --abi $root_dir/build/contracts/IService.abi --pkg contracts --out $dest_dir/iservice.go

