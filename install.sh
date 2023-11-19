#!/bin/bash
wget https://github.com/ethereum/solidity/releases/download/v0.8.23/solc-static-linux
mv solc-static-linux solc
chmod +x solc
cd client
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/pkg/mod/github.com/ethereum/go-ethereum@v1.13.5
make
make devtools
sudo apt install nodejs -y
sudo apt install npm -y
sudo npm install -g yarn
sudo npm install -g ganache