## 如何计算一个金库的抵押率？

Pando Leaf API不提供抵押率数据，但提供很多其他信息。我花了一些时间才弄清楚如何计算它。



1. HTTP GET `leaf-api.pando.im/api/vats/:vaultid`来获取`ink`,`art`,`collateral_id`。

   (将:vaultid替换为你要查询的金库的UUID)

   - `ink`是被锁定的抵押品的数量
   - `art`是规范化的债务数量
   - `collateral_id`是UUID，用于识别这个金库有哪些抵押品

2. HTTP GET `https://leaf-api.pando.im/api/cats`获取`gem`,`mat`,`duty`。

   - `gem`是锁定抵押品的UUID
   - `mat` 是这个抵押品的清算率
   - `duty`是稳定费(又称利率)

3. HTTP GET `https://leaf-api.pando.im/api/vats/%s/events`获取`created_at`,`dart`,`debt`。

   - `created_at`是这个事件发生时的时间
   - `dart `
   - `debt` 是总债务的变化



从第1步，我们可以得到`art`，它被称为 "正常化债务"，但不是当前的债务。随着时间的推移，利息的积累，所以实际的债务额总是高于`art`。这个 `art` 有点像不包含利息的债务。

为了得到真正的债务额，我必须知道利息累积了多少。我知道的唯一可能的方法是获得金库的最新事件。在一个事件的`dart！=0`的情况下，`debt/dart*art`将是事件发生时的总债务。

但这仍然不是终点。由于事件可能发生在很多天前，我们从`debt/dart*art`得到的数字仍然比真实的数字要少一点。我们将不得不计算从事件的created_at到今天产生的利息。

```
func compound(in,days,rate float64) float64{
                for i:=0; i<int(days); i++{
                        in = in * rate/365 + in
                }
                return in
        }
}
debt = compound(event_debt, event_days, interest_rate)
```

`event_debt`是`debt/dart*art`

`event_days`是`today-created_at`

`interest_rate`是 duty - 1

结果将是当前的债务。啊哈! 困难的部分已经完成了。

抵押率 = `ink * current / debt`。

---

之后发生了一个错误。第3步的响应是一个数组，而我使用数组[0]作为计算的数据源。这很糟糕。当一些不同的动作发生时（如`VatDeposit`和`VatWithdraw`），`dart`和``debt`将为0。 所以计算失败。抵押率将被显示为0。

为了避免这种情况，不要只使用数组[0]，找到最新的不等于0的`dart` 和`debt`，用它们来计算。然后根据`action`（`VatDeposit == +` , `VatWithDraw == -`）将`dink`与`ink`加减。

