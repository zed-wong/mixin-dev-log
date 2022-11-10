This is a function for calling mvm contracts. Replace uppercase words with your own.

``` javascript
async function callContract() {
  let abi = JSON.parse(fs.readFileSync('PATH TO YOUR CONTRACT ABI.JSON'));
  let address = "REPLACE WITH YOUR ADDRESS";
  let provider = new ethers.Wallet("fd9477620edb11e46679122475d61c56d8bfb753fe68ca5565bc1f752c5f0eeb", new ethers.providers.StaticJsonRpcProvider("https://quorum-mayfly-testnet.mixin.zone"));
  let contract = new ethers.Contract(address, abi, provider);
  let current = await contract.REPLACE_WITH_YOUR_FUNCTION_NAME();
  console.log(current)
}
```
