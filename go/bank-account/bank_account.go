package account

import "sync"

type Account struct {
    sync.Mutex
    balance int64
    isClosed bool
}

func Open(initialDeposit int64) *Account {
    if initialDeposit < 0 {
        return nil
    }

    return &Account{
        balance: initialDeposit,
        isClosed: false,
    }
}

func (a *Account) Close() (payout int64, ok bool) {
    a.Lock()
    defer a.Unlock()

    if a.isClosed {
        return 0, false
    }

    a.isClosed = true

    return a.balance, true
}

func (a *Account) Balance() (balance int64, ok bool) {
    return a.balance, !a.isClosed
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
    if a.isClosed {
        return 0, false
    }

    a.Lock()
    defer a.Unlock()

    if a.balance + amount < 0 {
        return 0, false
    }

    a.balance += amount

    return a.balance, true
}
