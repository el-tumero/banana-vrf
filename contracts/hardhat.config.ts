import { HardhatUserConfig } from "hardhat/config";
import { config as LoadEnv } from 'dotenv';
import "@nomicfoundation/hardhat-toolbox";

LoadEnv();

const config: HardhatUserConfig = {
  solidity: "0.8.19",
  networks: {
    hardhat: {
      forking: {
        url: 'https://rpc.testnet.lukso.network',
        blockNumber: 1466233,
      }
    },
    luksoTestnet: {
      url: 'https://rpc.testnet.lukso.network',
      chainId: 4201,
      accounts: [process.env.PRIVATE_KEY as string]
    }
  }
};

export default config;
