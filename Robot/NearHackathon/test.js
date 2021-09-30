const { keyStores, KeyPair } = require("near-api-js");
const nearAPI = require("near-api-js");
const keyStore = new keyStores.InMemoryKeyStore();
const Big = require("big-js");
const toNear = (value = '0') => Big(value).times(10 ** 24).toFixed()
const fs = require("fs");
const { send } = require("process");
const { CreateAccount } = require("near-api-js/lib/transaction");
const ACCOUNT_ID = "dispa1r.testnet"; // NEAR account tied to the keyPair
const NETWORK_ID = "testnet";
// path to your custom keyPair location (ex. function access key for example account)
const KEY_PATH = './dispa1r.testnet.json';
const credentials = JSON.parse(fs.readFileSync(KEY_PATH));
keyStore.setKey(NETWORK_ID, "dispa1r.testnet", KeyPair.fromString(credentials.private_key));

const { connect } = nearAPI;
const config = {
    networkId: "testnet",
    keyStore, // optional if not signing transactions
    nodeUrl: "https://rpc.testnet.near.org",
    walletUrl: "https://wallet.testnet.near.org",
    helperUrl: "https://helper.testnet.near.org",
    explorerUrl: "https://explorer.testnet.near.org",
};

async function deployContract() {
    const near = await connect(config);
    const account = await near.account("dispa1r.testnet");
    const response = await account.deployContract(fs.readFileSync('./linkdrop.wasm'));
    console.log(response);
}
const public_key = [];
const nearPriAccount = [];
async function createKeyPair(newAccountId, num) {
    const keyStore1 = new keyStores.UnencryptedFileSystemKeyStore("./");
    //const creatorAccount = await near.account(creatorAccountId);

    for (i = 0; i < num; i++) {
        const keyPair = KeyPair.fromRandom("ed25519");
        await keyStore1.setKey(config.networkId, newAccountId + i, keyPair);
        const KEY_PATH = './testnet/' + newAccountId + i + ".json";
        const credentials = JSON.parse(fs.readFileSync(KEY_PATH));
        //keyStore.setKey(NETWORK_ID, ACCOUNT_ID, KeyPair.fromString(credentials.private_key));
        public_key.push(keyPair.publicKey.toString().replace('ed25519:', ''));
        nearPriAccount.push(credentials.private_key.replace('ed25519:', ''));
    }
    //console.log(public_key);
    console.log(nearPriAccount);
    //await keyStore.setKey(config.networkId, "testnmsl1.testnet", keyPair)
}

async function getContract(viewMethods = [], changeMethods = [], secretKey) {
    const near = await connect(config);
    if (secretKey) {
        await keyStore.setKey(
            NETWORK_ID, "dispa1r.testnet",
            nearAPI.KeyPair.fromString(secretKey)
        )
    }
    const tmpAccount = await near.account("dispa1r.testnet");
    const contract = new nearAPI.Contract(tmpAccount, "dispa1r.testnet", {
        viewMethods,
        changeMethods,
        sender: "dispa1r.testnet"
    })
    return contract
}
async function getContract1(viewMethods = [], changeMethods = []) {
    const near = await connect(config);
    const tmpAccount = await near.account("dispa1r.testnet");
    const signAccount = await near.account("dispa1r1.testnet");
    const contract1 = new nearAPI.Contract(tmpAccount, "dispa1r.testnet", {
        viewMethods,
        changeMethods,
        sender: tmpAccount
    })
    return contract1
}

async function callSend(public_key, deposit) {
    const contract = await getContract1([], ['send'])
    const depositNum = toNear(deposit)
    await contract.send({
            public_key,
        }, 200000000000000, depositNum)
        .then(() => {
            console.log('Drop claimed')
        })
        .catch((e) => {
            console.log(e)
            console.log('Unable to claim drop. The drop may have already been claimed.')
        })
}


async function callSendLuck(nearAmount, num) {
    await createKeyPair("test.testnet", num);
    const contract = await getContract1([], ['send_luck'])
    const deposit = toNear(nearAmount);
    //console.log(public_key);
    await contract.send_luck({
            public_key,
            num,
        }, 200000000000000, deposit)
        .then(() => {
            console.log('Drop claimed')
        })
        .catch((e) => {
            console.log(e)
            console.log('Unable to claim drop. The drop may have already been claimed.')
        })
}

async function claimLuck(account_id, private_Key) {
    const contract = await getContract([], ['claim', 'create_account_and_claim'], private_Key)
        // return funds to current user
    await contract.claim({
            account_id,
        }, 200000000000000)
        .then(() => {
            console.log('Drop claimed')
        })
        .catch((e) => {
            console.log(e)
            console.log('Unable to claim drop. The drop may have already been claimed.')
        })
}


async function claimDrop(account_id, privKey) {
    const contract = await getContract([], ['claim', 'create_account_and_claim'], privKey)
        // return funds to current user
    await contract.claim({
            account_id,
        }, 200000000000000)
        .then(() => {
            console.log('Drop claimed')
        })
        .catch((e) => {
            console.log(e)
            console.log('Unable to claim drop. The drop may have already been claimed.')
        })
}

async function withDrawMoney(new_account_id,num){
    const deposit = toNear(num);
    const near = await connect(config);
    const account = await near.account("dispa1r.testnet");
    await account.sendMoney(
        new_account_id, // receiver account
        deposit // amount in yoctoNEAR
    );
}





//deployContract();
//callSendLuck();

if (process.argv.length == 5) {
    var arguments = process.argv.splice(2);
    const callFuntionName = arguments[0];
    const nearAmount = arguments[1];
    const num = arguments[2];
    console.log(nearAmount, num)
    if (callFuntionName == "callSendLuck") {
        callSendLuck(Number(nearAmount), Number(num))
    } else if (callFuntionName == "claimLuck") {
        const new_account_id = arguments[1];
        const private_Key = arguments[2];
        claimLuck(new_account_id, private_Key);
    } else if (callFuntionName == "callSend") {
        const public_key = arguments[1];
        const num = arguments[2];
        callSend(public_key, Number(num))
    } else if (callFuntionName == "claimDrop") {
        const new_account_id = arguments[1];
        const privateKey = arguments[2];
        claimDrop(new_account_id, privateKey);
    }else if (callFuntionName == "withdraw") {
        const new_account_id = arguments[1];
        const num = Number(arguments[2]);
        withDrawMoney(new_account_id,num);
    }
}





//deployContract()

//claimLuck()