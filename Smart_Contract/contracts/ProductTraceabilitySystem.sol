pragma solidity ^0.8.0;
import "./Event.sol";
import "./Model.sol";

contract ProductTraceabilitySystem is Event, Model {
    address admin;
    bool closed;
    mapping(uint32 => ProductionUnit) IDToUnit;
    mapping(address => uint32) addressToUnitID;
    mapping(address => bool) unitRegistered;
    mapping(uint32 => uint16) registerCounter;

    constructor() {
        closed = false;
        admin = msg.sender;
    }

    function generateProductID(uint256 _time) public view returns (bytes32) {
        require(!closed);
        require(unitRegistered[msg.sender]);

        bytes32 ID = keccak256(
            abi.encodePacked(addressToUnitID[msg.sender], _time)
        );
        return ID;
    }

    function produce(
        bytes32 productionID,
        uint8 _freshState,
        uint8 _category
    ) public {
        require(!closed);
        require(unitRegistered[msg.sender]);

        emit BindCategory(productionID, _category);
        emit Confirm(
            productionID,
            addressToUnitID[msg.sender],
            0,
            _freshState,
            block.timestamp
        );
    }

    function confirm(
        bytes32 productionID,
        uint8 _transPortState,
        uint8 _freshState,
        uint8
    ) public {
        require(!closed);
        require(unitRegistered[msg.sender]);
        emit Confirm(
            productionID,
            addressToUnitID[msg.sender],
            _transPortState,
            _freshState,
            block.timestamp
        );
    }

    function stringToBytes32(
        string memory source
    ) internal pure returns (bytes32 result) {
        assembly {
            result := mload(add(source, 32))
        }
    }

    function unitRegister(
        address _unitAddress,
        uint32 _addrCode,
        string calldata _name
    ) public returns (uint32) {
        require(!closed);
        require(msg.sender == admin);
        require(!unitRegistered[_unitAddress]);
        bytes32 name = stringToBytes32(_name);
        uint32 _ID = getNewUintID(_addrCode);
        addressToUnitID[_unitAddress] = _ID;
        IDToUnit[_ID] = ProductionUnit({ID: _ID, name: name});
        emit UnitRegister(_ID, block.timestamp);
        return _ID;
    }

    function getNewUintID(uint32 _addrCode) internal returns (uint32) {
        uint32 result = (_addrCode << 12) + registerCounter[_addrCode];
        registerCounter[_addrCode] += 1;
        return result;
    }
}