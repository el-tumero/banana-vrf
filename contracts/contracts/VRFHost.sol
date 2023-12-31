// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

contract VRFHost {

    enum RoundState {EMPTY, PROPOSAL, FINAL}
    struct Round {
        address proposer;
        uint256 randomNumber;
        bytes32 randomNumberHash;
        RoundState state;
        uint256 blockHeight;
    }
    mapping(uint32 => Round) private rounds;

    uint32 private currentRoundId = 1;
    uint256 private constant BLOCK_NUMBER_THRESHOLD = 5;
    uint256 private constant MIN_STAKE = 100;
    uint256 private constant LAST_INIT_ROUND_ID = 5;

    struct Operator {
        uint256 stake;
        uint32 sinceRound;
    }
    mapping (address => Operator) private operators;

    event NewRound(uint32 indexed _id);


    constructor(){
        // genesis random number
        rounds[0] = Round(
            address(0),
            0xf4af66a743b631375dbe319fa457aa07e08b0401f9a822f6c797722938143db2, 
            0x4d5f7dd6652bc469c41cc53d0c515d0254469f469272633b5b34ad32e2a10a0c, 
            RoundState.FINAL,
            block.number
        );
        rounds[1].blockHeight = block.number;
    }

    function verifySignature(bytes32 message, uint8 _v, bytes32 _r, bytes32 _s) public pure returns (address) {
        if (_v < 27) {
            _v += 27;
        }
        address signer = ecrecover(message, _v, _r, _s);
        return signer;
    }

    function verifyProposal(uint8 _v, bytes32 _r, bytes32 _s) public view returns (bool) {
        address signer = verifySignature(rounds[currentRoundId-1].randomNumberHash, _v, _r, _s);
        if(signer != address(0x0)){
            return true;
        }
        return false;
    }


    function getPreviousRandomNumber() public view returns (uint256) {
        return rounds[currentRoundId-1].randomNumber;
    }

    function getCurrentRandomNumber() public view returns (uint256) {
        return rounds[currentRoundId].randomNumber;
    }

    function setRandomNumber(uint8 _v, bytes32 _r, bytes32 _s) public {
        address signer = verifySignature(rounds[currentRoundId-1].randomNumberHash, _v, _r, _s);
        require(signer == msg.sender, "Wrong signature!");
        require(operators[signer].stake > MIN_STAKE, "No required stake!");
        require(currentRoundId - operators[signer].sinceRound > 3 || currentRoundId <= LAST_INIT_ROUND_ID, "You need to wait!");
        uint256 num = uint256(_r) << 128 | uint256(_s) >> 128;
        Round storage currentRound = rounds[currentRoundId];
        if(currentRound.state == RoundState.PROPOSAL) {
            require(num < currentRound.randomNumber, "Wrong signature or number is not valid!");
        }
        currentRound.randomNumber = uint256(_r) << 128 | uint256(_s) >> 128;
        currentRound.randomNumberHash = keccak256(abi.encode(currentRound.randomNumber));
        currentRound.state = RoundState.PROPOSAL;
        currentRound.proposer = signer;
    }

    function nextRound() public {
        Round storage currentRound = rounds[currentRoundId];
        require(
            block.number > currentRound.blockHeight + BLOCK_NUMBER_THRESHOLD && 
            msg.sender == currentRound.proposer &&
            currentRound.state == RoundState.PROPOSAL, 
            "Not permitted!");
        currentRound.state = RoundState.FINAL;
        currentRoundId++;
        rounds[currentRoundId].blockHeight = block.number;
        emit NewRound(currentRoundId);
    }

    function nextRoundLate() public {
        Round storage currentRound = rounds[currentRoundId];
        require(
            block.number > currentRound.blockHeight + BLOCK_NUMBER_THRESHOLD * 2 &&
            currentRound.state == RoundState.PROPOSAL,
            "Not permitted!");
        currentRound.state = RoundState.FINAL;
        currentRoundId++;
        rounds[currentRoundId].blockHeight = block.number;
        emit NewRound(currentRoundId);
    }

    function getRound(uint32 id) external view returns (Round memory) {
        return rounds[id];
    }

    function addStake() public payable {
        require(msg.value > 0, "Not permitted!");
        operators[msg.sender].stake += msg.value;
        operators[msg.sender].sinceRound = currentRoundId;
    }

    function withdrawStake() public {
        uint256 amount = operators[msg.sender].stake;
        require(amount > 0, "No funds to withdraw!");
        operators[msg.sender].stake = 0;
        (bool sent,) = payable(msg.sender).call{value: amount}("");
        require(sent, "Failed to send Ether");
    }

    function checkStake(address operator) public view returns(uint256) {
        return operators[operator].stake;
    }

    // debug
    function getValue() public pure returns (uint256) {
        return 1;
    }

    function getCurrentRoundId() external view returns (uint32) {
        return currentRoundId;
    }

    function isOperatorActive(address addr) public view returns (bool) {
        Operator memory operator = operators[addr];
        if(operator.stake > MIN_STAKE && (currentRoundId - operator.sinceRound > 3 || currentRoundId <= LAST_INIT_ROUND_ID)) return true;
        return false;
    }

    
}