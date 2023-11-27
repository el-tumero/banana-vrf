contract:
	make clean && solc --optimize --abi ./contracts/contracts/VRFHost.sol --bin ./contracts/contracts/VRFHost.sol -o ./contracts/build && abigen --abi=./contracts/build/VRFHost.abi --bin=./contracts/build/VRFHost.bin --pkg=contract --out ./client/contract/VRFHost.go
clean:
	rm -r ./contracts/build
node:
	ganache -s hello
testapi:
	cd tests && go run . -rpc https://rpc.testnet.lukso.network -chain_id 4201 -contract 0xD061CEb1F6BE5b6822762893e229FFce5C62C283