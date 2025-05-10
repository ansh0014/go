package account
import (
"sync"	// makes it thread- safe 
)

// Define the Account type here.

func Open(amount int64) *Account {
	if amount <0 {
	return nil
return &Account{balance: amount}
}
        }

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed   
        {
	return 0,false	
	}
	return a.balance,true
	panic("Please implement the Balance function")
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock(){
		defer a.mu.Unlock()
	if a.closed|| a.balance+amount<0 
                {
			return 0, false
                }
a.balance+=amount
return a.balance, true
			
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.close
        {
		return 0, false
	}
	a.closed = true
	payout:=a.balance
	return payout, true
	
}
