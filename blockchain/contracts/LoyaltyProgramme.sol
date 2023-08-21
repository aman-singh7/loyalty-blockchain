// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

// import "@openzeppelin/contracts/metatx/ERC2771Context.sol";

// contract LoyaltyProgramme is ERC2771Context {
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
        uint couponId; // will be indexed onto the vector
        address issuerBusiness;
        uint superCoins; // value of superCoin in each coupon
        uint count;
        uint discount; // multipiled by a factor of `decimalFactor`
        uint productCategory;
        uint thresholdValue;
        uint productId;
        CouponType couponType;
        uint expiryTime;
        bool active;
    }

    // Coupon NFT
    struct CouponHolding {
        uint holdingId;
        uint couponId;
        uint expiryDate;
        bool active;
    }

    struct Account {
        uint superCoins;
        uint currentHoldingId;
        mapping(uint => uint) couponHoldings; // holdings id to index mapping
        CouponHolding[] coupons; // will be available only for CONSUMER and others has empty
        AccountType accountType; // has default value unregistered
    }

    // state variables

    uint constant decimalFactor = 1000;
    uint constant couponRatio = 10; //(100/10 = 10%)
    address immutable OWNER_ADDRESS;
    mapping(address => Account) accounts;
    uint currentCouponId = 0;
    Coupon[] public couponList;

    // errors

    // supercoin transaction errors (identified uniquely by unique transactionId)

    /// senders account has less super coin than requested.
    /// needed `requested` supercoins but only `available` supercoins are available
    /// @param transactionId id to uniquely identify the transaction
    /// @param requested requested transfer amount
    /// @param available available amount in the senders account
    error InSufficientFunds(uint transactionId, uint requested, uint available);

    /// Only requestor with `validAuthority` can do this action
    /// @param transactionId id to uniquely identify the transaction
    /// @param validAuthority authority by whom action is allowed only
    /// @param senderAuthority authority of the requester
    error InValidAuthority(
        uint transactionId,
        AccountType validAuthority,
        AccountType senderAuthority
    );

    /// the `requestor` is not allowed to access this resource
    /// @param transactionId id to uniquely identify the transaction
    /// @param requestor address of the requestor
    error UnAuthorized(uint transactionId, address requestor);

    /// Only recipient with `validAuthority` can get this transaction money
    /// @param transactionId id to uniquely identify the transaction
    /// @param recieverAuthority authority of the recipient
    error InValidRecipient(uint transactionId, AccountType recieverAuthority);

    /// account type of the involved member is invalid for this transaction
    /// @param transactionId id to uniquely identify the transaction
    /// @param account wallet address of the involved member with invalid account type
    /// @param validAuthority expected account authority type for the member
    /// @param senderAuthority recieved account authority for the member
    error UnexpectedAccount(
        uint transactionId,
        address account,
        AccountType validAuthority,
        AccountType senderAuthority
    );

    /// Coupon is either already expired or invalid couponId is provided
    /// @param transactionId id to uniquely identify the transaction
    /// @param holdingId holdingId of the coupon holding
    error InValidCouponHolding(uint transactionId, uint holdingId);

    /// Coupon is either already expired or invalid couponId is provided
    /// @param transactionId id to uniquely identify the transaction
    /// @param couponId couponId of the coupon holding
    error InValidCoupon(uint transactionId, uint couponId);

    /// Coupons holdings are expired
    /// @param transactionId id to uniquely identify the transaction
    /// @param holdingId id of the coupon holding
    /// @param couponId couponId of the coupon
    /// @param expiryDate expiry date of the coupon
    error CouponHoldingExpired(
        uint transactionId,
        uint holdingId,
        uint couponId,
        uint expiryDate
    );

    /// senders account has less coupons than requested.
    /// needed `requested` coupons but only `available` coupons are available
    /// @param transactionId id to uniquely identify the transaction
    /// @param couponId requested couponId
    /// @param requested requested coupons
    /// @param available available coupons
    error InSufficientCoupons(
        uint transactionId,
        uint couponId,
        uint requested,
        uint available
    );

    /// consumer has applied coupon on wrong business.
    /// coupon is being applied on `receivedBusiness` but coupon is issued by `expectedBusiness`
    /// @param transactionId id to uniquely identify the transaction
    /// @param couponId requested couponId
    /// @param expectedBusiness business by whom the coupon is issued
    /// @param receivedBusiness business to whom the coupon is applied
    error InValidCouponHoldingApplication(
        uint transactionId,
        uint couponId,
        address expectedBusiness,
        address receivedBusiness
    );

    /// Invalid regisration
    /// User already has an account with `existingAuthority`
    /// @param transactionId id to uniquely identify the transaction
    /// @param existingAuthority existing authority of the user
    error InValidRegistration(
        uint transactionId,
        AccountType existingAuthority
    );

    error UnRegisteredAccount(uint transactionId, address person);

    /// Integrity constraints failed
    /// @param transactionId id to uniquely identify the transaction
    /// @param message message for the error
    error IntegrityConstraintFailed(uint transactionId, string message);

    // events

    /// event emittend when a transaction of `amount` supercoin is done
    /// from `sender` account to `reciever` account
    /// @param transactionId id to uniquely identify the transaction
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
    /// @param transactionId id to uniquely identify the transaction
    /// @param businessSender sender account
    /// @param consumerReceiver reciever account
    /// @param couponId couponId
    /// @param count count of coupons
    event CouponTransactionComplete(
        uint transactionId,
        address businessSender,
        address consumerReceiver,
        uint couponId,
        uint count
    );

    /// event emittend when a transaction of `count` coupons with `couponId` is done
    /// from `sender` account to `reciever` account
    /// @param transactionId id to uniquely identify the transaction
    /// @param consumerSender sender account
    /// @param businessReceiver reciever account
    /// @param holdingId holdingId of the coupon holding
    event CouponHoldingTransactionComplete(
        uint transactionId,
        address consumerSender,
        address businessReceiver,
        uint holdingId
    );

    /// event emitted when a new member is registered at address `address`
    /// with authority `authority`
    /// @param transactionId id to uniquely identify the transaction
    /// @param memberAddress address of the new joinee
    /// @param authority authority of the new joinee
    event MemberRegistered(
        uint transactionId,
        address memberAddress,
        AccountType authority
    );

    /// event emitted when new coupon are created
    /// @param couponId id of coupon which is registerd
    /// @param transactionId id to uniquely identify the transaction
    event CouponCreated(uint transactionId, uint couponId);

    // modifiers

    modifier checkOwner(address account, uint transactionId) {
        if (msg.sender != OWNER_ADDRESS && msg.sender != account) {
            revert UnAuthorized(transactionId, account);
        }
        _;
    }

    modifier checkSuperCoinFunds(
        uint transactionId,
        address sender,
        uint request
    ) {
        if (accounts[sender].superCoins < request) {
            revert InSufficientFunds(
                transactionId,
                request,
                accounts[sender].superCoins
            );
        }
        _;
    }

    modifier checkBusinessCouponFunds(
        uint transactionId,
        address businessAddress, // assuming it is passed after checking
        uint couponId,
        uint count
    ) {
        Coupon storage coupon = couponList[couponId];

        if (!coupon.active) {
            revert InValidCoupon(transactionId, couponId);
        }

        if (coupon.count < count) {
            revert InSufficientCoupons(
                transactionId,
                couponId,
                count,
                coupon.count
            );
        }
        _;
    }

    modifier checkConsumerCouponValidity(
        uint transactionId,
        address consumerAddress, // assuming it is passed after checking
        address businessAddress,
        uint holdingId
    ) {
        uint holdingIndex = accounts[consumerAddress].couponHoldings[holdingId];

        // all holdingIndexes start from 1
        if (holdingIndex == 0) {
            revert InValidCouponHolding(transactionId, holdingId);
        }

        CouponHolding storage consumerHolding = accounts[consumerAddress]
            .coupons[holdingIndex];
        Coupon storage coupon = couponList[consumerHolding.couponId];

        // coupon should be applied to only the business issuing the coupon
        if (coupon.issuerBusiness != businessAddress) {
            revert InValidCouponHoldingApplication(
                transactionId,
                consumerHolding.couponId,
                coupon.issuerBusiness,
                businessAddress
            );
        }

        assert(consumerHolding.active != false);

        if (consumerHolding.expiryDate < block.timestamp) {
            _removeCouponHolding(holdingIndex, consumerAddress);
            revert CouponHoldingExpired(
                transactionId,
                holdingId,
                consumerHolding.couponId,
                consumerHolding.expiryDate
            );
        }
        _;
    }

    modifier checkAccess(uint transactionId, AccountType allowedAuthority) {
        if (accounts[msg.sender].accountType != allowedAuthority) {
            revert InValidAuthority(
                transactionId,
                allowedAuthority,
                accounts[msg.sender].accountType
            );
        }
        _;
    }

    modifier checkRegistered(uint transactionId, address person) {
        if (accounts[person].accountType == AccountType.UNREGISTERED) {
            revert UnRegisteredAccount(transactionId, person);
        }
        _;
    }

    modifier checkAccountType(
        uint transactionId,
        address person,
        AccountType expectedAuthoirty
    ) {
        if (accounts[person].accountType != expectedAuthoirty) {
            revert UnexpectedAccount(
                transactionId,
                person,
                expectedAuthoirty,
                accounts[person].accountType
            );
        }
        _;
    }

    // internal functions

    function _transferSuperCoins(
        uint transactionId,
        address sender,
        address receiver,
        uint amount
    )
        internal
        checkRegistered(transactionId, sender)
        checkRegistered(transactionId, receiver)
        checkSuperCoinFunds(transactionId, sender, amount)
    {
        accounts[sender].superCoins -= amount;
        accounts[receiver].superCoins += amount;
        emit SuperCoinTransactionComplete(
            transactionId,
            sender,
            receiver,
            amount
        );
    }

    function _removeCouponHolding(
        uint holdingIndex,
        address consumerAddress
    ) internal {
        uint lastIndex = accounts[consumerAddress].coupons.length - 1;

        uint expiredHoldingId = accounts[consumerAddress]
            .coupons[holdingIndex]
            .holdingId;

        accounts[consumerAddress].coupons[holdingIndex] = accounts[
            consumerAddress
        ].coupons[lastIndex];

        uint holdingId = accounts[consumerAddress]
            .coupons[holdingIndex]
            .holdingId;

        accounts[consumerAddress].coupons.pop();

        accounts[consumerAddress].couponHoldings[holdingId] = holdingIndex;
        delete accounts[consumerAddress].couponHoldings[expiredHoldingId]; // will be set to 0
    }

    /**
     * Set the trustedForwarder address either in constructor or
     * in other init function in your contract
     */
    // public functions
    // constructor(
    //     address _trustedForwarder
    // ) public ERC2771Context(_trustedForwarder) {
    constructor() {
        OWNER_ADDRESS = msg.sender;
        Account storage onwerAccount = accounts[OWNER_ADDRESS]; // spelling
        onwerAccount.superCoins = 0;
        onwerAccount.accountType = AccountType.OWNER;
    }

    /**
     * Override this function.
     * This version is to keep track of BaseRelayRecipient you are using
     * in your contract.
     */
    // function versionRecipient() external view virtual returns (string memory) {
    //     return "1";
    // }

    // public consumer functions

    function consumerTokenPayment(
        uint transactionId,
        uint amount
    ) public checkAccess(transactionId, AccountType.CONSUMER) {
        _transferSuperCoins(transactionId, msg.sender, OWNER_ADDRESS, amount);
    }

    function consumerCouponPayment(
        uint transactionId,
        uint holdingId,
        address businessAddress
    )
        public
        checkAccess(transactionId, AccountType.CONSUMER)
        checkAccountType(transactionId, businessAddress, AccountType.BUSINESS)
        checkConsumerCouponValidity(
            transactionId,
            msg.sender,
            businessAddress,
            holdingId
        )
    {
        uint holdingIndex = accounts[msg.sender].couponHoldings[holdingId];
        _removeCouponHolding(holdingIndex, msg.sender);

        emit CouponHoldingTransactionComplete(
            transactionId,
            msg.sender,
            businessAddress,
            holdingId
        );
    }

    function purchaseCoupon(
        uint transactionId,
        address businessAddress,
        uint couponId,
        uint count
    )
        public
        checkAccess(transactionId, AccountType.CONSUMER)
        checkAccountType(transactionId, businessAddress, AccountType.BUSINESS)
        checkBusinessCouponFunds(
            transactionId,
            businessAddress,
            couponId,
            count
        )
    {
        Account storage consumerAccount = accounts[msg.sender];
        uint holdingId = consumerAccount.currentHoldingId;
        uint holdingIndex = consumerAccount.coupons.length;
        Coupon storage coupon = couponList[couponId];
        CouponHolding[] storage couponHoldings = accounts[msg.sender].coupons;

        // push coupons into consumer coupons list
        for (uint i = 0; i < count; i++) {
            couponHoldings.push(
                CouponHolding(
                    holdingId,
                    couponId,
                    block.timestamp + coupon.expiryTime,
                    true
                )
            );
            consumerAccount.couponHoldings[holdingId] = holdingIndex;
            holdingId = holdingId + 1;
            holdingIndex = holdingIndex + 1;
        }
        consumerAccount.currentHoldingId = holdingId;
        coupon.count -= count;

        // transfer coins into e-commerce account
        uint totalCost = count * coupon.superCoins;
        _transferSuperCoins(
            transactionId,
            msg.sender,
            OWNER_ADDRESS,
            totalCost
        );

        emit CouponTransactionComplete(
            transactionId,
            businessAddress,
            msg.sender,
            couponId,
            count
        );
    }

    // public business functions

    function redeemTokens(uint amount) public {
        //  better flow requires confirmation amountof both the parties
    }

    function deactiveCoupon(
        uint transactionId,
        uint couponId
    ) public checkAccess(transactionId, AccountType.BUSINESS) {
        if (!couponList[couponId].active) return;

        if (couponList[couponId].issuerBusiness != msg.sender) {
            revert UnAuthorized(transactionId, msg.sender);
        }

        couponList[couponId].active = false;
    }

    // public owner functions

    function createCoupons(
        uint transactionId,
        address issuerBusiness,
        uint count,
        uint cost,
        uint discount,
        uint productCategory,
        uint thresholdValue,
        uint productId,
        CouponType couponType,
        uint expiryTime,
        uint creationCost
    ) public checkAccess(transactionId, AccountType.OWNER) returns (uint) {
        // checks

        if (count <= 0) {
            revert IntegrityConstraintFailed(
                transactionId,
                "count should be > 0"
            );
        }

        if (cost < 0) {
            revert IntegrityConstraintFailed(
                transactionId,
                "cost should be >= 0"
            );
        }

        if (discount < 0 || discount > 100) {
            revert IntegrityConstraintFailed(
                transactionId,
                "discount should be >= 0 and <= 100"
            );
        }

        if (couponType == CouponType.UNIQUE) {
            if (productId < 0)
                revert IntegrityConstraintFailed(
                    transactionId,
                    "when coupon type is unique then productId can't be negative"
                );
        } else if (couponType == CouponType.CATEGORY) {
            if (productCategory < 0) {
                revert IntegrityConstraintFailed(
                    transactionId,
                    "when coupon type is category wise then productCategories can't be negative"
                );
            }

            if (thresholdValue < 0) {
                revert IntegrityConstraintFailed(
                    transactionId,
                    "when coupon type is category wise then threshold can't be negative"
                );
            }
        }

        if (expiryTime <= 0) {
            revert IntegrityConstraintFailed(
                transactionId,
                "expiryTime must be > 0"
            );
        }

        uint couponId = currentCouponId;
        couponList.push(
            Coupon(
                couponId,
                issuerBusiness,
                cost,
                count,
                discount,
                productCategory,
                thresholdValue,
                productId,
                couponType,
                expiryTime,
                true
            )
        );

        accounts[issuerBusiness].superCoins -= creationCost;
        currentCouponId = currentCouponId + 1;
        return couponId;
    }

    function registerMember(
        uint transactionId,
        address memberAddress,
        AccountType accountType
    ) public checkAccess(transactionId, AccountType.OWNER) {
        if (accounts[memberAddress].accountType != AccountType.UNREGISTERED) {
            revert InValidRegistration(
                transactionId,
                accounts[memberAddress].accountType
            );
        }
        Account storage memberAccount = accounts[memberAddress];
        memberAccount.superCoins = 0;
        memberAccount.accountType = accountType;
        memberAccount.coupons.push(CouponHolding(0, 0, 0, false)); // pass 0th index as invalid element ( active = false )
        memberAccount.currentHoldingId = 1;
        emit MemberRegistered(transactionId, memberAddress, accountType);
    }

    function pay(
        uint transactionId,
        address account,
        uint amount
    ) public checkAccess(transactionId, AccountType.OWNER) {
        if (accounts[account].accountType == AccountType.OWNER) {
            revert InValidRecipient(
                transactionId,
                accounts[account].accountType
            );
        }
        _transferSuperCoins(transactionId, OWNER_ADDRESS, account, amount);
    }

    function mintTokens(
        uint transactionId,
        uint amount
    ) public checkAccess(transactionId, AccountType.OWNER) returns (uint) {
        accounts[msg.sender].superCoins += amount;
        return accounts[msg.sender].superCoins;
    }

    function getAccountBalance(
        uint transactionId,
        address person
    )
        public
        view
        checkOwner(person, transactionId)
        checkRegistered(transactionId, person)
        returns (uint)
    {
        return accounts[person].superCoins;
    }

    function getAllCoupons(
        uint transactionId
    )
        public
        view
        checkRegistered(transactionId, msg.sender)
        returns (Coupon[] memory)
    {
        Coupon[] memory results;
        uint j = 0;
        for (uint i = 0; i < couponList.length; i++) {
            if (couponList[i].active) {
                results[j++] = couponList[i];
            }
        }
        return results;
    }

    function getBrandCoupons(
        uint transactionId,
        address brandAddress
    )
        public
        view
        checkAccountType(transactionId, brandAddress, AccountType.BUSINESS)
        returns (Coupon[] memory)
    {
        Coupon[] memory results;
        uint j = 0;
        for (uint i = 0; i < couponList.length; i++) {
            if (
                couponList[i].active &&
                couponList[i].issuerBusiness == brandAddress
            ) {
                results[j++] = couponList[i];
            }
        }
        return results;
    }
}
