// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract Simple0xExchange {
    // 订单结构
    struct Order {
        address maker;           // 订单创建者
        address makerToken;      // Maker 提供的代币地址
        uint256 makerAmount;     // Maker 提供的代币数量
        address takerToken;      // Maker 期望接收的代币地址
        uint256 takerAmount;     // Maker 期望接收的代币数量
        uint256 expiry;          // 订单有效期
        bool filled;             // 订单是否已完成
    }

    // 订单哈希 => 订单数据
    mapping(bytes32 => Order) public orders;

    // 事件
    event OrderCreated(bytes32 orderHash, address indexed maker);
    event OrderFilled(bytes32 orderHash, address indexed taker);
    event OrderCancelled(bytes32 orderHash, address indexed maker);

    // 创建订单
    function createOrder(
        address makerToken,
        uint256 makerAmount,
        address takerToken,
        uint256 takerAmount,
        uint256 expiry
    ) external returns (bytes32) {
        require(makerAmount > 0, "Maker amount must be greater than 0");
        require(takerAmount > 0, "Taker amount must be greater than 0");
        require(expiry > block.timestamp, "Expiry must be in the future");

        // 生成订单哈希
        bytes32 orderHash = keccak256(
            abi.encodePacked(
                msg.sender,
                makerToken,
                makerAmount,
                takerToken,
                takerAmount,
                expiry
            )
        );

        // 确保订单未存在
        require(orders[orderHash].maker == address(0), "Order already exists");

        // 保存订单
        orders[orderHash] = Order({
            maker: msg.sender,
            makerToken: makerToken,
            makerAmount: makerAmount,
            takerToken: takerToken,
            takerAmount: takerAmount,
            expiry: expiry,
            filled: false
        });

        emit OrderCreated(orderHash, msg.sender);
        return orderHash;
    }

    // 填充订单
    function fillOrder(bytes32 orderHash) external {
        Order storage order = orders[orderHash];

        require(order.maker != address(0), "Order does not exist");
        require(!order.filled, "Order already filled");
        require(order.expiry > block.timestamp, "Order expired");

        // 检查 Maker 的代币授权
        IERC20 makerToken = IERC20(order.makerToken);
        require(
            makerToken.allowance(order.maker, address(this)) >= order.makerAmount,
            "Maker token allowance insufficient"
        );

        // 检查 Taker 的代币授权
        IERC20 takerToken = IERC20(order.takerToken);
        require(
            takerToken.allowance(msg.sender, address(this)) >= order.takerAmount,
            "Taker token allowance insufficient"
        );

        // 执行代币转移
        require(
            makerToken.transferFrom(order.maker, msg.sender, order.makerAmount),
            "Maker token transfer failed"
        );
        require(
            takerToken.transferFrom(msg.sender, order.maker, order.takerAmount),
            "Taker token transfer failed"
        );

        // 标记订单为已完成
        order.filled = true;

        emit OrderFilled(orderHash, msg.sender);
    }

    // 取消订单
    function cancelOrder(bytes32 orderHash) external {
        Order storage order = orders[orderHash];

        require(order.maker == msg.sender, "Only the maker can cancel the order");
        require(!order.filled, "Cannot cancel a filled order");

        // 删除订单
        delete orders[orderHash];

        emit OrderCancelled(orderHash, msg.sender);
    }
}
