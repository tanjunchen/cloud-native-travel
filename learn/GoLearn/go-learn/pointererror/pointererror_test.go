package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	fmt.Println("address of balance in test is", &wallet.balance)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

func TestWallet2(t *testing.T) {

	wallet := WalletBitcoin{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestWallet4(t *testing.T) {
	wallet := WalletBitcoin{}
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := Bitcoin(20)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestWallet5(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := WalletBitcoin{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := WalletBitcoin{balance: Bitcoin(20)}

		wallet.Withdraw(10)

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestWalletRefactor(t *testing.T) {
	assertBalance := func(t *testing.T, walletBitcoin WalletBitcoin, want Bitcoin) {
		got := walletBitcoin.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		walletBitcoin := WalletBitcoin{}
		walletBitcoin.Deposit(Bitcoin(10))
		assertBalance(t, walletBitcoin, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		walletBitcoin := WalletBitcoin{balance: Bitcoin(20)}
		walletBitcoin.Withdraw(Bitcoin(10))
		assertBalance(t, walletBitcoin, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := WalletBitcoin{startingBalance}
		err := wallet.WithdrawCheck(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		fmt.Println(err)
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}

func TestWalletRefactor2(t *testing.T) {

	assertBalance := func(t *testing.T, walletBitcoin WalletBitcoin, want Bitcoin) {
		got := walletBitcoin.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		walletBitcoin := WalletBitcoin{}
		walletBitcoin.Deposit(Bitcoin(10))
		assertBalance(t, walletBitcoin, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		walletBitcoin := WalletBitcoin{balance: Bitcoin(20)}
		walletBitcoin.Withdraw(Bitcoin(10))
		assertBalance(t, walletBitcoin, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := WalletBitcoin{startingBalance}
		err := wallet.WithdrawCheck(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}



func TestWalletFinal(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := WalletBitcoin{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := WalletBitcoin{Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := WalletBitcoin{Bitcoin(20)}
		err := wallet.WithdrawCheck(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, InsufficientFundsError)
	})
}

func assertBalance(t *testing.T, walletBitcoin WalletBitcoin, want Bitcoin) {
	got := walletBitcoin.Balance()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}
