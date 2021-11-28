## How to calculate a vault's current ratio?

Pando Leaf API doesn't provide current ratio data, but provides a lot of other information. It took me a while to figure out how to calculate it.



1. HTTP GET `leaf-api.pando.im/api/vats/:vaultid` to get `ink`,`art`,`collateral_id`

   (replace :vaultid with the UUID of vault you want to query)

   - `ink` is the number of locked collateral
   - `art` is the number of normalized debt
   - `collateral_id ` is the UUID to identify what kind of collateral this vault has

2. HTTP GET `https://leaf-api.pando.im/api/cats` to get `gem`,`mat`,`duty`

   - `gem` is the UUID of locked collateral
   - `mat` is the liquidation ratio of this collateral
   - `duty` is the stability fee (aka interest rate)

3. HTTP GET `https://leaf-api.pando.im/api/vats/%s/events` to get `created_at`,`dart`,`debt`

   - `created_at ` is the time when this event created
   - `dart` 
   - `debt` is the change of total debt.



From step 1,  we can get `art` which is called "normalized debt", but not the current debt. With time goes by, the interest accumulates, so the real amount of debt would always be higher than `art`. This `art` is kinda like a debt number without interest.

To get the real amount of debt, I will have to know how much the interest have accumulated. The only possible way I know yet is to get the latest events of the vault. With an event's dart != 0, `debt/dart*art` will be the total debt at the moment of the event.

But it's still not the end. Since the event can be happened in many days ago, the number we got from `debt/dart*art` would still be slightly less than the real one. We will have to calculate the interest generated from event's created_at to today. 

```go
func compound(in,days,rate float64) float64{
                for i:=0; i<int(days); i++{
                        in = in * rate/365 + in
                }
                return in
        }
}
debt = compound(event_debt, event_days, interest_rate)
```

`event_debt` is `debt/dart*art`

`event_days` is `today-created_at`

`interest_rate` is duty - 1

The result will be the current debt. Aha! The hard part is done.

Current ratio = `ink * current / debt`

---

A bug occurs afterward. The response of step 3 is an array, and I used array[0] as the data source for calculation. This is bad. When some different action happens (like `vatDeposit` and `vatWithdraw`) , the `dart` and `debt` will be 0. So the calculation fails. Current ratio will be shown as 0.

To avoid this, don't just use array[0], find the latest `dart` and `debt` that is not equal to 0 instead, use them for calculation. Then add or subtract `dink` with `ink` base on `action` (`VatDeposit == +`  , `VatWithDraw == -`)

