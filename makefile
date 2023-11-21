contract:
	make clean && solc --optimize --abi ./contracts/VRFHost.sol --bin ./contracts/VRFHost.sol -o ./contracts/build && abigen --abi=./contracts/build/VRFHost.abi --bin=./contracts/build/VRFHost.bin --pkg=contract --out ./client/contract/VRFHost.go
clean:
	rm -r ./contracts/build
node:
	ganache -s hello