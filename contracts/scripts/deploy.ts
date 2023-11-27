import { ethers } from "hardhat";

const VRFHostAddress = "0xD061CEb1F6BE5b6822762893e229FFce5C62C283"

async function main() {
  const LSP8Random = await ethers.getContractFactory("LSP8Random")
  const contract = await LSP8Random.deploy(VRFHostAddress)

  const address = await contract.getAddress()
  console.log(address)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
