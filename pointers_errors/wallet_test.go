// @Title Jackcoin wallet
// @Description Wallet for the new jackcoin
// @Author Belaoura Yacoub
// @Update v0.0.1
package pointerserrors

import "testing"

// make a Jackcoin wallet
//keypoints
// -pointers
// -errors
// -new type from int
// -string interface
// -t.Fatal to break from test
// -var create an err with error.new
// -var define globale value to package

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Jackcoin(10))

		want := Jackcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Jackcoin(20)}

		err := wallet.Withdraw(Jackcoin(10))
		want := Jackcoin(10)

		assertNoBalance(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Jackcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Jackcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("wanted an error but didn't get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertNoBalance(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Jackcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
