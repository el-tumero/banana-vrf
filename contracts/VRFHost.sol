// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

contract VRFHost {

    uint32 private round = 0; // current round id
    uint256 private prevRandNumber = 0xf4af66a743b631375dbe319fa457aa07e08b0401f9a822f6c797722938143db2; // prev random number
    uint256 private currRandNumber = 0; // current random number

    function verifySignature(bytes32 message, uint8 _v, bytes32 _r, bytes32 _s) public pure returns (address) {
        bytes32 hash = keccak256(abi.encode(message));
        if (_v < 27) {
            _v += 27;
        }
        address signer = ecrecover(hash, _v, _r, _s);
        return signer;
    }

    function getPreviousRandomNumber() public view returns (uint256) {
        return prevRandNumber;
    }

    // debug
    function getValue() public pure returns (uint256) {
        return 1;
    }
    
    // function setNewRandomNumber(uint8 _v, bytes32 _r, bytes32 _s) public{

    // }

    // function hello() public pure returns (string memory) {
    //     return "Hello";
    // }


}