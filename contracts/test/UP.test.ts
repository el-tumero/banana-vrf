import { ethers } from "hardhat"
import { dataLength, ethers as eth } from "ethers"
import { Signer } from "ethers"
import { config as LoadEnv } from 'dotenv';
import { LSPFactory } from "@lukso/lsp-factory.js";
import UniversalProfile from '@lukso/lsp-smart-contracts/artifacts/UniversalProfile.json'

LoadEnv()

let impSigner:Signer
const rpc = "http://0.0.0.0:8545"
const chainId = 31337
const address = "0xA9b06d5a84cE114871ffCaBA39d933816fAAE070"

const VRFHostAddr = "0xD061CEb1F6BE5b6822762893e229FFce5C62C283"

async function createUniversalProfile() {
    const lspFactory = new LSPFactory(rpc, {
        deployKey: process.env.PRIVATE_KEY,
        chainId: chainId
    })

    const deployedContracts = await lspFactory.UniversalProfile.deploy({
        controllerAddresses: [address],
        lsp3Profile: {
            name: "My super cool UP",
            description: "Very very cool",
            tags: ["Public profile"],
            links: [
                {
                    title: "My website",
                    url: "https://example.com"
                }
            ]
        }
    })

    return deployedContracts.LSP0ERC725Account.address
}

describe("LSP8 tests", () => {
    it("Deployment of contracts", async () => {
        impSigner = await ethers.getImpersonatedSigner(address);
        const provider = new eth.JsonRpcProvider(rpc)
        const upAddress = await createUniversalProfile()

        const universalProfile = new ethers.Contract(
            upAddress,
            UniversalProfile.abi,
            provider
        )

        const LSP8Random = await ethers.getContractFactory("LSP8Random")
        const randomToken = await LSP8Random.deploy(VRFHostAddr)

        const callData = randomToken.interface.encodeFunctionData(
            'preMint'
        )

        const tx0 = await randomToken.connect(impSigner).preMint()
        const recp0 = await tx0.wait()
        const data0 = await randomToken.connect(impSigner).getRefValue(address)
        console.log(data0.toString())

        
        // const tx = await impSigner.sendTransaction({
        //     from: address,
        //     to: upAddress,
        //     data: callData
        // })

        // const recp = await tx.wait()
        // console.log(recp)

        // const data = await randomToken.connect(impSigner).getRefValue(upAddress)
        // console.log(data.toString())






        // console.log(universalProfile)


    })


}) 