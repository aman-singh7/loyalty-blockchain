// Help Truffle find `LoyaltyProgramme.sol` in the `/contracts` directory
const LoyaltyProgramme = artifacts.require("LoyaltyProgramme");

module.exports = function(deployer) {
  // Command Truffle to deploy the Smart Contract
  deployer.deploy(LoyaltyProgramme);
};