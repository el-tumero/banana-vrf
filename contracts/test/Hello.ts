import { expect } from "chai";
import { Hello } from "../typechain-types";
import { ethers } from "hardhat";


let contract:Hello

describe("Hello contract", () => {

  it("Deploy", async() => {
    const Hello = await ethers.getContractFactory("Hello")
    contract = await Hello.deploy()
  })

  it("hello()", async() => {
    const res = await contract.hello()
    expect(res).to.be.equal("Hello")
  })

});



