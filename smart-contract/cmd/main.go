// generate-abi: solc --abi smart-contract/src/go/basic.sol -o smart-contract/src/artifacts/abi

//generate-type-safe-go: abigen --abi smart-contract/src/artifacts/abi/Basic.abi --pkg main --type Basic --out smart-contract/src/go/Basic.go

//generate-bin: solc --bin smart-contract/src/go/basic.sol -o smart-contract/src/artifacts/bin

//generate-deployable: abigen --abi smart-contract/src/artifacts/abi/Basic.abi --pkg main --type Basic --out smart-contract/src/go/Basic.go --bin smart-contract/src/artifacts/bin/Basic.bin

