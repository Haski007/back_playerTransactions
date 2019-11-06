# back_playerTransactions

Backend site for players transactions.

There are realised requests:

* AddUser
    - creating user to cache(map[uint64]*User)
* GetUser
    - get .json format status of user
* AddDeposit
    - create deposit for certain user
* Transaction
    - type: Bet - make transaction Bet, reduses user balance
    - type: Win - make transaction Win, increases user balance
