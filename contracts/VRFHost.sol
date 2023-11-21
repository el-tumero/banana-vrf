// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

contract VRFHost {

    uint32 private round = 0; // current round id
    uint256 private prevRandNumber = 0xf4af66a743b631375dbe319fa457aa07e08b0401f9a822f6c797722938143db2; // prev random number
    bytes32 private hashPrevRandNumber = 0x4d5f7dd6652bc469c41cc53d0c515d0254469f469272633b5b34ad32e2a10a0c; // prev random number hashed§
    uint256 private currRandNumber = 0; // current random number
    bool private isSet = false;

    function verifySignature(bytes32 message, uint8 _v, bytes32 _r, bytes32 _s) public pure returns (address) {
        bytes32 hash = keccak256(abi.encode(message));
        if (_v < 27) {
            _v += 27;
        }
        address signer = ecrecover(hash, _v, _r, _s);
        return signer;
    }

    function verifyProposal(uint8 _v, bytes32 _r, bytes32 _s) public view returns (bool) {
        address signer = verifySignature(hashPrevRandNumber, _v, _r, _s);
        if(signer != address(0x0)){
            return true;
        }
        return false;
    }


    function getPreviousRandomNumber() public view returns (uint256) {
        return prevRandNumber;
    }

    function getCurrentRandomNumber() public view returns (uint256) {
        return currRandNumber;
    }

    function setRandomNumber(uint8 _v, bytes32 _r, bytes32 _s) public {
        bool isVerified = verifyProposal(_v, _r, _s);
        require(isVerified, "Wrong signature!");
        uint256 num = uint256(_r) << 128 | uint256(_s) >> 128;
        if(isSet) {
            require(num <= currRandNumber, "Wrong signature or number is not valid!");
        }
        currRandNumber = uint256(_r) << 128 | uint256(_s) >> 128;
    }


    // debug
    function getValue() public pure returns (uint256) {
        return 1;
    }
    
}