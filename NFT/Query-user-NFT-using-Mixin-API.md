Mixin has provided [an API](https://developers.mixin.one/docs/api/multisigs/outputs) for querying NFTs. But it wasn't really clear for a newbie to use. I will try to explain it.

# [Read Signature Outputs](https://developers.mixin.one/docs/api/multisigs/outputs)

1. ### Authorization:Authorized

​		The Bearer token should be signed by the owner of the NFTs. Which means if you wanna query some user's NFT list, you should use the user's token to query instead of your bot's token.

2. ### Parameters

​		There are five parameters, three of them are optional. Only "members" and "threshold" is required.

​		When you want to get someone's NFT list, those parameters should be:

```go
# Example of golang
UserID := "44d9717d-8cae-4004-98a1-f9ad544dcfb1"
members := []string{UserID}
threshold := uint8(1)
```

​		When using fox-one's [mixin-sdk-go](https://github.com/fox-one/mixin-sdk-go), the whole thing that prints all NFTs of a user would be like:

```go
# Example of golang
UserID := "44d9717d-8cae-4004-98a1-f9ad544dcfb1"
members := []string{UserID}
threshold := uint8(1)
offset, _ := time.Parse(time.RFC3339Nano,"2018-02-12T12:12:12.999999999Z")
limits := 500

nfts, err := rw.client.ReadCollectibleOutputs(ctx, members, threshold, offset, limits)
if err != nil {
        log.Println(err)
}
for _, i := range nfts{
        log.Printf("%+v", i)
}
```



Hope it will somehow help you.