// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract LoyaltyProgramme {
    // Contract constructs

    enum AccountType {
        UNREGISTERED,
        CONSUMER,
        BUSINESS,
        OWNER
    }

    enum CouponType {
        CATEGORY,
        UNIQUE
    }

    struct Coupon {
        address issuerBusiness;
        uint superCoins; // value of superCoin in each coupon
        uint count;
        uint discount; // multipiled by a factor of `decimalFactor`
        string productCategory;
        uint thresholdvalue;
        string productId;
        CouponType couponType;
        uint expiry;
        bool active;
    }

    struct Account {
        uint superCoins;
        mapping(uint => Coupon) coupons;
        AccountType accountType; // has default value unregistered
    }

    // state variables

    uint constant decimalFactor = 1000;
    address immutable OWNER_ADDRESS;
    mapping(address => Account) accounts;

    // errors

    /// senders account has less super coin than requested.
    /// needed `requested` supercoins but only `available` supercoins are available
    /// @param requested requested transfer amount
    /// @param available available amount in the senders account
    error InSufficientFunds(uint requested, uint available);

    /// Only requestor with `validAuthority` can do this action
    /// @param validAuthority authority by whom action is allowed only
    /// @param senderAuthority authority of the requester
    error UnAuthorized(AccountType validAuthority, AccountType senderAuthority);

    /// Only recipient with `validAuthority` can get this transaction money
    /// @param validAuthority authority to whom transaction is allowed only
    /// @param recieverAuthority authority of the recipient
    error InValidRecipient(
        AccountType validAuthority,
        AccountType recieverAuthority
    );

    /// account type of the involved member is invalid for this transaction
    /// @param account wallet address of the involved member with invalid account type
    /// @param expectedType expected account type for the member
    /// @param recievedType recieved account type for the member
    error UnexpectedAccount(
        address account,
        AccountType expectedType,
        AccountType recievedType
    );

    /// Coupon is either already expired or invalid couponId is provided
    /// @param couponId provided couponId
    error InvalidCoupons(uint couponId);

    /// Coupons are expired
    /// @param expiry expiry date of the coupon
    /// @param couponId couponId of the coupon
    error CouponsExpired(uint expiry, uint couponId);

    /// senders account has less coupons than requested.
    /// needed `requested` coupons but only `available` coupons are available
    /// @param couponId requested couponId
    /// @param requested requested coupons
    /// @param available available coupons
    error InSufficientCoupons(uint couponId, uint requested, uint available);

    /// Invalid regisration
    /// User already has an account with `existingAuthority`
    /// @param existingAuthority existing authority of the user
    error InValidRegistration(AccountType existingAuthority);

    // events

    /// event emittend when a transaction of `amount` supercoin is done
    /// from `sender` account to `reciever` account
    /// @param transactionId transactionId of the transaction
    /// @param sender sender account
    /// @param receiver reciever account
    /// @param amount amount transferred
    event SuperCoinTransactionComplete(
        uint transactionId,
        address sender,
        address receiver,
        uint amount
    );

    /// event emittend when a transaction of `count` coupons with `couponId` is done
    /// from `sender` account to `reciever` account
    /// @param transactionId transactionId of the transaction
    /// @param sender sender account
    /// @param receiver reciever account
    /// @param couponId couponId
    /// @param count count of coupons
    event CouponTransactionComplete(
        uint transactionId,
        address sender,
        address receiver,
        uint couponId,
        uint count
    );

    /// event emitted when a new member is registered at address `address`
    /// with authority `authority`
    /// @param memberAddress address of the new joinee
    /// @param authority authority of the new joinee
    event MemberRegistered(address memberAddress, AccountType authority);

    // event emitted when new coupon are created
    event CouponCreated(uint couponId);

    // modifiers

    modifier checkSuperCoinFunds(address sender, uint request) {
        if (accounts[sender].superCoins < request) {
            revert InSufficientFunds(request, accounts[sender].superCoins);
        }
        _;
    }

    modifier checkCouponFunds(
        address sender,
        uint couponId,
        uint count
    ) {
        if (!accounts[sender].coupons[couponId].active) {
            revert InvalidCoupons(couponId);
        }
        if (accounts[sender].coupons[couponId].count < count) {
            revert InSufficientCoupons(
                couponId,
                count,
                accounts[sender].coupons[couponId].count
            );
        }
        _;
    }

    modifier checkCouponExpiry(address sender, uint couponId) {
        if (!accounts[sender].coupons[couponId].active) {
            revert InvalidCoupons(couponId);
        }
        if (accounts[sender].coupons[couponId].expiry > block.timestamp) {
            uint expiry = accounts[sender].coupons[couponId].expiry;
            delete accounts[sender].coupons[couponId];
            revert CouponsExpired(expiry, couponId);
        }
        _;
    }

    modifier checkAccess(AccountType allowedAuthority) {
        if (accounts[msg.sender].accountType != allowedAuthority) {
            revert UnAuthorized(
                allowedAuthority,
                accounts[msg.sender].accountType
            );
        }
        _;
    }

    modifier checkAccountType(address person, AccountType expectedType) {
        if (accounts[person].accountType != expectedType) {
            revert UnexpectedAccount(
                person,
                expectedType,
                accounts[person].accountType
            );
        }
        _;
    }

    // internal functions

    // add transaction id
    function _transferSuperCoins(
        uint transactionId,
        address sender,
        address receiver,
        uint amount
    ) internal checkSuperCoinFunds(sender, amount) {
        accounts[sender].superCoins -= amount;
        accounts[receiver].superCoins += amount;
        emit SuperCoinTransactionComplete(
            transactionId,
            sender,
            receiver,
            amount
        );
    }

    // set expiry after purchase and check expiry
    // reclaim option
    // add transactio id
    function _transferCoupons(
        uint transactionId,
        address sender,
        address receiver,
        uint couponId,
        uint count
    ) internal checkCouponFunds(sender, couponId, count) returns (uint) {
        Coupon storage senderCoupons = accounts[sender].coupons[couponId];
        Coupon memory receiverCoupons = accounts[receiver].coupons[couponId];

        senderCoupons.count -= count;

        if (receiverCoupons.active) {
            receiverCoupons.count += count;
        } else {
            receiverCoupons.count = count;
        }

        accounts[receiver].coupons[couponId] = receiverCoupons; // storage variable assigned to memory, data gets copied
        emit CouponTransactionComplete(
            transactionId,
            sender,
            receiver,
            couponId,
            count
        );
        return count * senderCoupons.superCoins;
    }

    // public functions

    constructor() {
        OWNER_ADDRESS = msg.sender;
        Account storage onwerAccount = accounts[OWNER_ADDRESS]; // spelling
        onwerAccount.superCoins = 0;
        onwerAccount.accountType = AccountType.OWNER;
    }

    // public consumer functions

    function consumerTokenPayment(
        uint transactionId,
        uint amount
    ) public checkAccess(AccountType.CONSUMER) {
        _transferSuperCoins(transactionId, msg.sender, OWNER_ADDRESS, amount);
    }

    function consumerCouponPayment(uint transactionId, uint couponId) public {
        // delete coupon
    }

    // remove check expiry
    function purchaseCoupon(
        uint transactionId,
        address businessAddress,
        uint couponId,
        uint count
    )
        public
        checkAccess(AccountType.CONSUMER)
        checkAccountType(businessAddress, AccountType.BUSINESS)
        checkCouponFunds(businessAddress, couponId, count)
        checkCouponExpiry(businessAddress, couponId)
    {
        _transferCoupons(
            transactionId,
            msg.sender,
            businessAddress,
            couponId,
            count
        );
        // flipkart token transfer and org token transfer
        // set expiry
    }

    // public business functions

    function redeemTokens(uint amount) public {
        //  better flow requires confirmation of both the parties
    }

    function createCoupons(
        uint couponId,
        uint count,
        uint cost,
        uint discount,
        string calldata productCategory,
        uint thresholdvalue,
        string calldata productId,
        CouponType couponType,
        uint expiry
    ) public checkAccess(AccountType.BUSINESS) {
        // checks
        require(count > 0, "count of tokens should be greater than 0");

        if (couponType == CouponType.UNIQUE) {
            require(
                bytes(productId).length != 0,
                "productId should not be empty when coupon type is unique product"
            );
        } else {
            require(
                bytes(productCategory).length != 0,
                "product category should not be empty when coupon type is category"
            );

            require(
                thresholdvalue >= 0,
                "threshold value for the product should be always greater than 0"
            );
        }

        require(
            discount >= 0 * decimalFactor && discount <= 100 * decimalFactor,
            "discount should be valid ( between 0 to 100 )"
        );

        // update no checks
        require(
            expiry > block.timestamp,
            "expiry date of the coupons must be after the "
        );

        Coupon storage coupons = accounts[msg.sender].coupons[couponId];
        coupons.superCoins = cost;
        coupons.issuerBusiness = msg.sender;
        coupons.count = count;
        coupons.discount = discount;
        coupons.productCategory = productCategory;
        coupons.thresholdvalue = thresholdvalue;
        coupons.productId = productId;
        coupons.couponType = couponType;
        coupons.expiry = expiry;

        emit CouponCreated(couponId);
    }

    // public owner functions

    function registerMember(
        address memberAddress,
        AccountType accountType
    ) public checkAccess(AccountType.OWNER) {
        if (accounts[memberAddress].accountType != AccountType.UNREGISTERED) {
            revert InValidRegistration(accounts[memberAddress].accountType);
        }
        Account storage memberAccount = accounts[memberAddress];
        memberAccount.superCoins = 0;
        memberAccount.accountType = accountType;
        emit MemberRegistered(memberAddress, accountType);
    }

    function payBusiness(
        uint transactionId,
        address business,
        uint amount
    ) public checkAccess(AccountType.OWNER) {
        if (accounts[business].accountType != AccountType.BUSINESS) {
            revert InValidRecipient(
                AccountType.BUSINESS,
                accounts[business].accountType
            );
        }
        _transferSuperCoins(transactionId, OWNER_ADDRESS, business, amount);
    }

    function mintTokens(
        uint amount
    ) public checkAccess(AccountType.OWNER) returns (uint) {
        accounts[msg.sender].superCoins += amount;
        return accounts[msg.sender].superCoins;
    }
}
