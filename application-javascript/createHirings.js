/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const { buildCCPOrg1, buildCCPOrg2, buildWallet, prettyJSONString} = require('../test-application/javascript/AppUtil.js');

const myChannel = 'mychannel';
const myChaincodeName = 'borderpay';

async function createAuction(ccp,wallet,user,hiringID,AccountID,salary,variablepay,currency,company) {
	try {

		const gateway = new Gateway();

		// Connect using Discovery enabled
		await gateway.connect(ccp,
			{ wallet: wallet, identity: user, discovery: { enabled: true, asLocalhost: true } });

		const network = await gateway.getNetwork(myChannel);
		const contract = network.getContract(myChaincodeName);

		let statefulTxn = contract.createTransaction('CreateHiring');

		console.log('\n--> Submit Transaction: Propose a new Hiring');
		await statefulTxn.submit(hiringID,AccountID,salary,variablepay,currency,company);
		console.log('*** Result: committed');

		console.log('\n--> Evaluate Transaction: query the auction that was just created');
		// let result = await contract.evaluateTransaction('QueryAuction',auctionID);
		// console.log('*** Result: Auction: ' + prettyJSONString(result.toString()));

		gateway.disconnect();
	} catch (error) {
		console.error(`******** FAILED to submit bid: ${error}`);
	}
}

async function main() {
	try {

		if (process.argv[2] === undefined || process.argv[3] === undefined ||
            process.argv[4] === undefined || process.argv[5] === undefined ||
            process.argv[6] === undefined || process.argv[7] === undefined||
            process.argv[8] === undefined || process.argv[9] === undefined) {
			console.log('Usage: node createAuction.js org userID auctionID item');
			process.exit(1);
		}

		const org = process.argv[2];
		const user = process.argv[3];
		const hiringID = process.argv[4];
		const AccountID = process.argv[5];
		const salary = process.argv[6];
		const variablepay = process.argv[7];
		const currency = process.argv[8];
		const company = process.argv[9];

		if (org === 'Org1' || org === 'org1') {
			const ccp = buildCCPOrg1();
			const walletPath = path.join(__dirname, 'wallet/org1');
			const wallet = await buildWallet(Wallets, walletPath);
			await createAuction(ccp,wallet,user,hiringID,AccountID,salary,variablepay,currency,company);
		}  else {
			console.log('Usage: node createAuction.js org userID auctionID item');
			console.log('Org must be Org1 or Org2');
		}
	} catch (error) {
		console.error(`******** FAILED to run the application: ${error}`);
	}
}


main();
