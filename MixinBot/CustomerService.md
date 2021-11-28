





```sequence
user->Bot: xxx
user->Bot: message id 1
Note left of Bot: save messageid2:(userid,convid,messageid1)
Bot->me : xxx
Bot->me : message id 2
me-> Bot: yyy
me-> Bot: message id 3 (quote 2)
Note right of Bot: query message(quoteid 2)
Bot->user: yyy
Bot->user: message id 4 (quote 1)
```

