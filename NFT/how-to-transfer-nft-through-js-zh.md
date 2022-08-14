# 如何使用javascript转账NFT(详细教程)

基本上，转移NFT需要5个步骤。

1. OAuth用户获得用户的[JWT token](https://developers.mixin.one/docs/api/oauth)

2. 用用户的JWT token请求[Mixin API](https://developers.mixin.one/docs/api/collectibles/outputs#get-collectiblesoutputs) `GET /collectibles/outputs`，以获得用户的outputs(UTXO)。`state`参数必须为`unspent`。

3. 请求 [Mixin API](https://developers.mixin.one/docs/api/collectibles/outputs#get-collectiblestokensuuid) `GET /collectibles/tokens/:token_id` , 将`:token_id`替换为步骤2中返回的token_id。

4. 请求 [Mixin API](https://developers.mixin.one/docs/api/collectibles/request#post-collectiblesrequests) `POST /collectibles/requests` 来创建一个交易请求。参数`action`是`sign`。而参数`raw`需要耗费一些力气来生成。

例如，在[`mixin-note-sdk`](https://github.com/liuzemei/bot-api-nodejs-client)中提供了一个函数来生成`raw`。这个[函数](https://github.com/liuzemei/bot-api-nodejs-client/blob/main/src/client/collectibles.ts#L68)叫做`makeCollectibleTransactionRaw`。我们可以像`makeCollectibleTransactionRaw( { output, token, [recipient_user_id], threshold:1 } )`这样调用它。

- `output`来自于步骤2 
- `token'来自于步骤3 
- `[recipient_user_id]`是NFT接收者的用户ID。`[]`是必须有的，因为它是一个数组。
- `threshold:1`表示NFT是发给一个人的。

调用这个函数后，我们将得到`raw`。然后我们可以创建一个转账（多签）请求。

这种生成`raw'的方式可能不是最好的方式。但这并不重要，因为它确实有效。

5. 使用用户的JWT token 轮询[Mixin API](https://developers.mixin.one/docs/api/collectibles/outputs#get-collectiblesoutputs) `GET /collectibles/outputs`，`state`参数为`signed`。一旦找到output/UTXO。发送一个主网交易，参数为第四步中的`raw`，以完成转账。


---

下面是用于完成整个过程的javascript代码。

代码假设步骤1,2,3已经完成。

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

// 这个url是用来给用户扫描以创建一个NFT的请求。
const url = `https://mixin.one/codes/${createRes.code_id}`。

// 用户付款后，第4步就完成了。我们需要轮询UTXO并发送主网交易。
```

```
// 轮询UTXO并发送TX
async loopPaymentState(output) {
  const JWTtoken; // 替换为你的值
  const userID;   // 替换为你的值
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

整体过程大概就是这样，如果你有遇到任何问题，请在下方留言。
