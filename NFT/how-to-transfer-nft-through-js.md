# How to transfer NFT through javascript (Detailed guide)

Basically, there're 5 steps required to transfer an NFT.

1. OAuth user to get user's [JWT token](https://developers.mixin.one/docs/api/oauth)

2. Request [Mixin API](https://developers.mixin.one/docs/api/collectibles/outputs#get-collectiblesoutputs) `GET /collectibles/outputs` with JWT token of the user to get user's collectible outputs(UTXO). Must carry `state` argument as `unspent`.

3. Request [Mixin API](https://developers.mixin.one/docs/api/collectibles/outputs#get-collectiblestokensuuid) `GET /collectibles/tokens/:token_id` , replace `:token_id` with the token_id from the step 2's response.

4. Request [Mixin API](https://developers.mixin.one/docs/api/collectibles/request#post-collectiblesrequests) `POST /collectibles/requests` to create a transaction request. The argument `action` should be `sign`. And the argument `raw` requires some effort to generate. 

For example, there's a function provided in [`mixin-node-sdk`](https://github.com/liuzemei/bot-api-nodejs-client) to generate the `raw`. This [function](https://github.com/liuzemei/bot-api-nodejs-client/blob/main/src/client/collectibles.ts#L68) called `makeCollectibleTransactionRaw`. We can call it like `makeCollectibleTransactionRaw( { output, token, [recipient_user_id], threshold:1 } )`. 

- `output` is from step 2 
- `token` is from step 3 
- `[recipient_user_id]` is the NFT receiver's user id. `[]` is required because it's an array.
- `threshold:1` means the NFT is sent to a single person.

After calling this function, we will have `raw` generated. Then we can create a transfer(multi-sig) request.

This way of generating `raw` might not be the best way to do so. But it doesn't matter much since it does work.

5. Loop [Mixin API](https://developers.mixin.one/docs/api/collectibles/outputs#get-collectiblesoutputs) `GET /collectibles/outputs` with JWT token of the user and the `state` argument as `signed`. Once the output/UTXO is found. Send a mainnet transaction with the argument `raw` to accomplish the final step.

---

Here is the javascript code that is used to do the whole thing.

Assume step 1 and step 2 are done.

```
const raw = await MixinClient.makeCollectibleTransactionRaw({
  output,
  token,
  receivers,
  threshold: 1,
});

const createRes = await createCollectibleRequest(
  JWTtoken,
  "sign",
  raw,
);

// This url is used for users to scan to create a collectible request.
const url = `https://mixin.one/codes/${createRes.code_id}`;

// After the user paid, step 4 is done. We need to loop UTXO and send mainnet transaction.
```

```
// Loop UTXO and send TX
async loopPaymentState(output) {
  const JWTtoken; // replace with your method
  const userID;   // replace with your method
  while (true) {
    await this.getPaymentState(output, userID, JWTtoken);
    await new Promise((resolve) => setTimeout(resolve, 1000));
  }
},

async getPaymentState(output, userID, JWTtoken) {
  const outputs = await getSignedOutputs(JWTtoken, [userID]);
  if (outputs.length == 0) return;
  outputs.forEach(async (element) => {
    if (element.output_id === output.output_id) {
      await MixinClient.sendRawTransaction(element.signed_tx);
      return;
    }
  });
},

async function getSignedOutputs(token, userids) {
  let config = {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
  let resp = await axios.get(`https://api.mixin.one/collectibles/outputs?members=${hashMembers(userids)}&threshold=1&state=signed`, config)
  if (resp.data) {
    return resp.data.data
  }
}

const hashMembers = (ids) => {
  const key = ids.sort().join('');
  const sha = new JsSHA('SHA3-256', 'TEXT', { encoding: 'UTF8' });
  sha.update(key);
  return sha.getHash('HEX');
};
```

That's about it! If you have any questions, leave a comment down below.