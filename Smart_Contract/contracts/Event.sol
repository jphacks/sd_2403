pragma solidity ^0.8.0;

interface Event {
    event Confirm(
        bytes32 indexed productionID,
        uint32 indexed productionUnitID,
        uint8 transPortState,
        uint8 freshState,
        uint256 timeStamp
    );
    event UnitRegister(uint32 indexed productionUnitID, uint256 timeStamp);
    event BindCategory(bytes32 indexed productionID, uint8 indexed category);
}
