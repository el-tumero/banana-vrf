import { expect } from "chai";
import { VRFHost } from "../typechain-types";
import { ethers } from "hardhat";


let contract:VRFHost

describe("VRFHost contract", () => {

  it("Deploy", async() => {
    const VRFHost = await ethers.getContractFactory("VRFHost")
    contract = await VRFHost.deploy()
  })

  it("verifySignature()", async() => {

    const [user0] = await ethers.getSigners()


    // const res = await contract.verifySignature()
    // expect(res).to.be.equal("Hello")
  })

});



