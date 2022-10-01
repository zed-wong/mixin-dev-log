When deploying the original contract:
---------------------------------------------------------
      | Deploy Factory | Create Pay | Approve | Deposit |
---------------------------------------------------------
MVM   |       o        |     o (XIN)|    o    |    x    |
---------------------------------------------------------
Goerli|       o        |     o (UNI)|    o    |    o    |
---------------------------------------------------------

Goerli:
Factory 0x030ee4c2d8059249d5768c09b49c9f27114f152b
Pay 0x98661d613c6D522235c4Ed4163B757894d1d316D
Approve 10000000000000000	(Decimal 18)
Deposit 1000000000000000        (Succeed)

MVM:
Factory 0x2B36Bc8fb4fD51B0f893cE296d43ABC091d17727
Pay 0x40f8393e09293943c128745063fa9ce076000091
Approve 100000			(Decimal 8)
Deposit 100000			(ERROR:Gas 0.095)

Possible reason:
The Deposit function called the erc20.SafeTransferFrom while Erc20 only have TransferFrom

Solution:
Try to modify it to Token.TransferFrom

---

---------------------------------------------------------
      | Deploy Factory | Create Pay | Approve | Deposit |
---------------------------------------------------------
MVM   |       o        |     o (XIN)|    o    |    o    |
---------------------------------------------------------

Factory 0x4A172B7A8d1FE92EEd72755F1A523a455be9013f
Pay 0x55fdeec677729364d01c5dbeb79af5284c4e3b51
Approve 10000			(Decimal 8)
Deposit 10000			(Succeed)

Fixed!
