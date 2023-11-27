import { ethers } from "hardhat";
import { VRFHost } from "../typechain-types";

const VRFHostAddress = "0xD061CEb1F6BE5b6822762893e229FFce5C62C283"

async function main() {
  const VRFHost = await ethers.getContractFactory("VRFHost")
  const contract = await VRFHost.attach(VRFHostAddress).waitForDeployment() as VRFHost

  const currentRoundId = await contract.getCurrentRoundId()
  console.log("current round id", currentRoundId)

  const data = await contract.getRound(currentRoundId)
    
  console.log(data)

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
