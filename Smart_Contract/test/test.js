const ProductTrace = artifacts.require("ProductTraceabilitySystem")

module.exports = function (deployer) {
    deployer.deploy(ProductTrace);
};
