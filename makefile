contract:
	solc --optimize --abi ./contracts/VRFHost.sol --bin ./contracts/VRFHost.sol -o ./contracts/build && abigen --abi=./contracts/build/VRFHost.abi --bin=./contracts/build/VRFHost.bin --pkg=contract --out ./client/contract/VRFHost.go
deploy:


clean:
	rm ./contracts/build/VRFHost.abi rm ./contracts/build/VRFHost.bin
node:
	ganache -s hello