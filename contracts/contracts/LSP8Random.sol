// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {
    LSP8Mintable
} from "../node_modules/@lukso/lsp-smart-contracts/contracts/LSP8IdentifiableDigitalAsset/presets/LSP8Mintable.sol";
import {
    _LSP8_TOKENID_TYPE_NUMBER
} from "../node_modules/@lukso/lsp-smart-contracts/contracts/LSP8IdentifiableDigitalAsset/LSP8Constants.sol";
import "./VRFConsumer.sol";

bytes32 constant _LSP4_TOKEN_TYPE_DATA_KEY = 0xe0261fa95db2eb3b5439bd033cda66d56b96f92f243a8228fd87550ed7bdfdb3;

enum TokenType {
    TOKEN,
    NFT,
    COLLECTION
}

contract LSP8Random is LSP8Mintable, VRFConsumer {
    constructor(address _VRFHostAddress) LSP8Mintable("RandomToken", "RNDT", msg.sender, _LSP8_TOKENID_TYPE_NUMBER) VRFConsumer(_VRFHostAddress){
        _setData(_LSP4_TOKEN_TYPE_DATA_KEY, abi.encode(1));
    }

    mapping(address => uint256) preMints;

    function preMint() public returns (uint256) {
        uint256 ref = saveRandomValue();
        preMints[msg.sender] = ref;
        return ref;
    }

    function getPreMint() public view returns(uint256) {
        return refs[preMints[msg.sender]];
    }

    function mint() public returns (bytes32) {
        uint256 id = preMints[msg.sender];
        uint256 rnd = readRandomValue(id);
        bytes32 tokenId = bytes32(rnd);
        _mint(msg.sender, tokenId, true, "0x");
        return tokenId;
    }

}