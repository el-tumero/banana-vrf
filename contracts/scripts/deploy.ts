import { ethers } from "hardhat";

async function deploy() {
    const VRFHost = await ethers.getContractFactory("VRFHost")
    const contract = await VRFHost.deploy()

    console.log(await contract.getAddress())
}

deploy()