# MTG Development Guide (MTG 开发指南)

MTG, as known as Mixin Trusted Group, is a technic built for developers to develop decentralized applications based on Mixin Network. The UTXO model of Mixin network doesn't support that very well, that's why MTG borns. It simplifies it and making it possible to develop decentralized applications in an easier and more secure way.

Formerly, Lyric wrote some [articles](https://quill.im/764392/9085a9e1-5624-4d16-a502-2e75487162ac) to introduce what MTG is in plain words. Those articles are nice but not developer oriented, you won't know where and how to start developing MTG applications after reading. The offcial [documentation](https://developers.mixin.one/docs/mainnet/guide/mtg-guide) is developer oriented, but it only explains some basic concepts, you would still have no idea how to build a MTG app from scratch. It's not worthless, at least you can have a basic understanding of what you will need to handle, but that's all.

Until now, my general impression of MTG is, a behavior lock and multisig wallet for multiple mixin bots.

Let's say there're 3 Mixin bots, the developer deployed the MTG application to each bot, and these 3 bots run the application at the same time. The codes running on 3 bots are the same. So these 3 bots will have the exact same reaction for certain inputs. 

Compared to the classical single bot applications, the MTG applications are more secure and decentralized. A stable MTG application requires multiple team's corporation, and since the codes running on each party are the same, the application will always have the same behavior.

For example, a hacker tries to exploit the application, if the bot runs alone, the hacker can just hack the bot's server, steal the keystone of the bot, or replace the codes with his own codes and steal the money. This could done silently and without anyone noticing. But that's not so easy for MTG applications. To attack the same way, the hacker will have to exploit most of the MTG members. And even if he got all servers, he will still need to write more complicated codes to steal the money. That's way more work than hacking a single bot application, it also leaves more time for the developer to rescue.

This also prevents the development team from doing something evil. If the team did choose MTG members in a truly decentralized way, it's almost impossible for them to steal the money by themselves. Beacuse to do so they have to get approval from most of the MTG members, otherwise it will never work.

So, how to develop the application from the strach? There're three key points.

1. Processing Multisig UTXO
2. Creating Multisig transactions
3. Creating Multisig payment links

Expect these three points, the rest parts would be quite same like single bot applications.


## Development

Here is a basic example in the [MTG SDK](https://github.com/MixinNetwork/trusted-group/blob/105bc828e5d5f9527dfe61f523482200eb61374d/mtg/README.md). This example illustrates how you run MTG applications with workers. Since this article is for beginners, here is the step by step illustration:

Firstly, you define a worker:
```
type RefundWorker struct {
        grp mtg.Group
}
```
Then, create the initialization function of the worker:
```
func NewRefundWorker(grp mtg.Group) *RefundWorker {
        return &RefundWorker{
                grp: grp,
        }
}
```
After that, implement two methods:
```
func (rw *RefundWorker) ProcessOutput(ctx context.Context, out *mtg.Output) {
        // set receiver to sender of the output
        receivers := []string{out.Sender}

        // generate a trace ID
        traceId := mixin.UniqueConversationID(out.UTXOID, "refund")

        // build transaction and write to MTG storage
        err := rw.grp.BuildTransaction(ctx, out.AssetID, receivers, int(1), out.Amount.String(), "refund", traceId, "")
        if err != nil {
                panic(err)
        }
}
func (rw *RefundWorker) ProcessCollectibleOutput(ctx context.Context, out *mtg.CollectibleOutput) {
        return
}
```
These two methods decides what the MTG will do when there's an output found. `ProcessOutput` is for normal transactions, `ProcessCollectibleOutput` is for NFT transactions. This is where the core logics implemented.

And to run the MTG, you can do all the initialization in your project root's 'main.go'. 
```
        // Init context
        ctx := context.Background()

        // Read configuration
        // https://github.com/MixinNetwork/trusted-group/blob/master/mvm/boot.go#L36

        // Init database
        db, err := store.OpenBadger(ctx, bp)
        if err != nil {
                return err
        }
        defer db.Close()

        // Init MTG Group
        group, err := mtg.BuildGroup(ctx, db, conf.MTG)
        if err != nil {
                return err
        }

        // Init refundWorker
        refundWorker := workers.NewRefundWorker(*group)
        group.AddWorker(refundWorker)
        group.Run(ctx)
```

The code part of running a MVP is done, then we will need to prepare the config file.
```config.toml
groupsize=2                                 // The number of MTG
[mtg.app]                                   // From bot keystore
client-id=""
session-id=""
private-key=""
pin-token=""
pin=""
[mtg.genesis]
members=[                                   // The client id of each MTG members
  "83726d6d-59da-454e-8ea6-5632f7ea5260",
  "e7dba065-5141-4160-bc28-4e8aedc8266a"
]
threshold=1                                 // Multisig threshold (minium members required to move the money)
timestamp=0
```
Save it under project root. Make a copy, and replace the keystore with another bot's.

To run the MTG, build it, run with configuration file and database file provided. This could be different based on the implemention. In my case:
`./main boot -c config/config1.toml -d ~/.mtg/data2`
`./main boot -c config/config2.toml -d ~/.mtg/data2`

Until now, the MTG is successfully running. To test it, we will generate a payment link that transfer money to MTG and see if it will be refunded.

For CLI: `mixin-cli -f CONFIG.json transfer --receivers "UUID1,UUID2" --threshold 1 --asset 965e5c6e-434c-3fa9-b780-c50f43cd955c --amount 1 --qrcode --trace $(uuidgen)`
Install mixin-cli (https://github.com/fox-one/mixin-cli), replace CONFIG.json with your bot's keystore, replace UUID1,UUID2 with your MTG members. If you haven't install `uuidgen`, you can replace it with new uuid.

For messenger: https://github.com/fox-one/mixin-sdk-go/blob/master/payment.go

For bot: https://github.com/fox-one/mixin-sdk-go/blob/master/transfer.go

You will get a qrcode after running the command for CLI. Once you paid for it, you should get some log message on your screen and get refunded soon.


### mtg.BuildTransaction(GroupID) meaning?
