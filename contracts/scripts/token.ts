import { ethers } from "hardhat";
import { LSP8Random } from "../typechain-types";

const randomTokenAddress = "0xc6848d02f1d35a81999A913f1E50ebD36800Fc3A"

async function main() {
  const LSP8Random = await ethers.getContractFactory("LSP8Random")
  const contract = await LSP8Random.attach(randomTokenAddress).waitForDeployment() as LSP8Random
  
  const [addr0] = await ethers.getSigners()

  const tx = await contract["mint()"]()
  const recp = await tx.wait()

  const data = await contract.tokenIdsOf(addr0.address)
  console.log(data)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});