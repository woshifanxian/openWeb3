// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity ^0.8.24;

contract ERC1820Registry {
    bytes4 internal constant INVALID_ID = 0xffffffff;
    bytes4 internal constant ERC165ID = 0x01ffc9a7;
    // 标准定义的一个常量，如果合约实现了某地址的接口，则返回这个常量
    bytes32 internal constant ERC1820_ACCEPT_MAGIC =
        keccak256(abi.encodePacked("ERC1820_ACCEPT_MAGIC"));
    // 保存地址、接口到实现合约地址的映射关系，对应着注册表记录的3 个内容
    mapping(address => mapping(bytes32 => address)) internal interfaces;
    // 映射地址到管理者
    mapping(address => address) internal managers;
    // 每个地址和 ERC165 接口的 flag，指示是否被缓存
    mapping(address => mapping(bytes4 => bool)) internal erc165Cached;
    // 接口实现事件
    event InterfaceImplementerSet(
        address indexed addr,
        bytes32 indexed interfaceHash,
        address indexed implementer
    );
    // 更改管理事件
    event ManagerChanged(address indexed addr, address indexed newManager); // 获取给定地址及接口的实现合约地址。

    function getInterfaceImplementer(address _addr, bytes32 _interfaceHash)
        external
        view
        returns (address)
    {
        address addr = _addr == address(0) ? msg.sender : _addr;
        if (isERC165Interface(_interfaceHash)) {
            bytes4 erc165InterfaceHash = bytes4(_interfaceHash);
            return
                implementsERC165Interface(addr, erc165InterfaceHash)
                    ? addr
                    : address(0);
        }
        return interfaces[addr][_interfaceHash];
    }

    // 设置某个地址的接口由哪个合约实现，需要由管理员来设置
    function setInterfaceImplementer(
        address _addr,
        bytes32 _interfaceHash,
        address _implementer
    ) external {
        address addr = _addr == address(0) ? msg.sender : _addr;
        require(getManager(addr) == msg.sender, "Not the manager");

        require(
            !isERC165Interface(_interfaceHash),
            "Must not be an ERC165 hash"
        );
        if (_implementer != address(0) && _implementer != msg.sender) {
            require(
                ERC1820ImplementerInterface(_implementer)
                    .canImplementInterfaceForAddress(_interfaceHash, addr) ==
                    ERC1820_ACCEPT_MAGIC,
                "Does not implement the interface"
            );
        }
        interfaces[addr][_interfaceHash] = _implementer;
        emit InterfaceImplementerSet(addr, _interfaceHash, _implementer);
    }

    // 为地址_addr 设置新的管理员地址
    function setManager(address _addr, address _newManager) external {
        require(getManager(_addr) == msg.sender, "Not the manager");
        managers[_addr] = _newManager == _addr ? address(0) : _newManager;
        emit ManagerChanged(_addr, _newManager);
    }

    // 获取地址_addr 的管理员
    function getManager(address _addr) public view returns (address) {
        // By default the manager of an address is the same address
        if (managers[_addr] == address(0)) {
            return _addr;
        } else {
            return managers[_addr];
        }
    }

    // 返回接口的 keccak256 哈希值
    function interfaceHash(string calldata _interfaceName)
        external
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(_interfaceName));
    }

    /* --- ERC165 相关方法 --- */
    // 更新合约是否实现了 ERC165 接口的缓存。
    function updateERC165Cache(address _contract, bytes4 _interfaceId)
        external
    {
        interfaces[_contract][_interfaceId] = implementsERC165InterfaceNoCache(
            _contract,
            _interfaceId
        )
            ? _contract
            : address(0);
        erc165Cached[_contract][_interfaceId] = true;
    }

    // 检查合约是否实现了 ERC165 接口
    function implementsERC165Interface(address _contract, bytes4 _interfaceId)
        public
        view
        returns (bool)
    {
        if (!erc165Cached[_contract][_interfaceId]) {
            return implementsERC165InterfaceNoCache(_contract, _interfaceId);
        }
        return interfaces[_contract][_interfaceId] == _contract;
    }

    // 在不使用缓存的情况下检查合约是否实现 ERC165 接口
    function implementsERC165InterfaceNoCache(
        address _contract,
        bytes4 _interfaceId
    ) public view returns (bool) {
        uint256 success;
        uint256 result;

        (success, result) = noThrowCall(_contract, ERC165ID);
        if (success == 0 || result == 0) {
            return false;
        }

        (success, result) = noThrowCall(_contract, INVALID_ID);
        if (success == 0 || result != 0) {
            return false;
        }

        (success, result) = noThrowCall(_contract, _interfaceId);
        if (success == 1 && result == 1) {
            return true;
        }
        return false;
    }

    // 检查_interfaceHash 是否为 ERC165 接口
    function isERC165Interface(bytes32 _interfaceHash)
        internal
        pure
        returns (bool)
    {
        return
            _interfaceHash &
                0x00000000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF ==
            0;
    }

    // 调用合约接口，如果函数不存在也不抛出异常
    function noThrowCall(address _contract, bytes4 _interfaceId)
        internal
        view
        returns (uint256 success, uint256 result)
    {
        bytes4 erc165ID = ERC165ID;
        assembly {
            let x := mload(0x40) // Find empty storagelocation using "free memory pointer"
            mstore(x, erc165ID) // Place signature atbeginning of empty storagecanImplementInterfaceForAddress
            mstore(add(x, 0x04), _interfaceId) // Place first argumentdirectly next to signature
            success := staticcall(
                30000, // 30k gas
                _contract, // To addr
                x, // Inputs are stored atlocation x
                0x24, // Inputs are 36 (4 + 32)bytes long
                x, // Store output over input(saves space)
                0x20 // Outputs are 32 byteslong
            )
            result := mload(x) // Load the result
        }
    }
}

interface ERC1820ImplementerInterface {
    /// @notice Indicates whether the contract implements the interface 'interfaceHash' for the address 'addr' or not.
    /// @param interfaceHash keccak256 hash of the name of the interface
    /// @param addr Address for which the contract will implement the interface
    /// @return ERC1820_ACCEPT_MAGIC only if the contract implements 'interfaceHash' for the address 'addr'.
    function canImplementInterfaceForAddress(
        bytes32 interfaceHash,
        address addr
    ) external view returns (bytes32);
}
