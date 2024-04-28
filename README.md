# blockchain-project
BLockchain Border.io implementation


## Deploy the chaincode

We will run the auction smart contract using the Fabric test network. Open a command terminal and navigate to the test network directory:
```
cd blockchain-project/test-network
```

You can then run the following command to deploy the test network.
```
./network.sh up createChannel -ca
```

Note that we use the `-ca` flag to deploy the network using certificate authorities. We will use the CA to register and enroll our sellers and buyers.

Run the following command to deploy the auction smart contract. We will override the default endorsement policy to allow any channel member to create an auction without requiring an endorsement from another organization.
```
./network.sh deployCC -ccn borderpay -ccp ../border-pay/chaincode-go/ -ccl go -ccep "OR('Org1MSP.peer')"
```

## Install the application dependencies

We will interact with the auction smart contract through a set of Node.js applications. Change into the `application-javascript` directory:
```
cd blockchain-project/application-javascript
```

From this directory, run the following command to download the application dependencies:
```
npm install
```

## Register and enroll the application identities

To interact with the network, you will need to enroll the Certificate Authority administrators of Org1 and Org2. You can use the `enrollAdmin.js` program for this task. Run the following command to enroll the Org1 admin:
```
node enrollAdmin.js org1
```

We can use the CA admin to register and enroll the identities of the seller that will create the auction and the bidders who will try to purchase the painting.

Run the following command to register and enroll the seller identity that will create the auction. The seller will belong to Org1.
```
node registerEnrollUser.js org1 seller
```

You should see the logs of the seller wallet being created as well. Run the following commands to register and enroll 2 bidders from Org1 and another 2 bidders from Org2:
```
node registerEnrollUser.js org1 bidder1
node registerEnrollUser.js org1 bidder2
node registerEnrollUser.js org2 bidder3
node registerEnrollUser.js org2 bidder4
```
