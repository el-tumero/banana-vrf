contract:
	./solc --optimize --abi ./contracts/contracts/VRFHost.sol -o ./contracts/build && abigen --abi=./contracts/build/VRFHost.abi --pkg=contract --out ./client/contract/VRFHost.go
clean:
	rm ./contracts/build/VRFHost.abi