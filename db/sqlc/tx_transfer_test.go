package db

import (
	"testing"
)

func TestTxTransfer(t *testing.T) {

	// account1 := createAccountTest(t)
	// account2 := createAccountTest(t)

	// n := 5
	// amount := int64(10)
	// errs := make(chan error)
	// results := make(chan TransferTxResult)

	// var wg sync.WaitGroup
	// wg.Add(n) // ✅ to'g'ridan-to'g'ri n marta `Add` qilish

	// for i := 0; i < n; i++ {
	// 	go func(i int) {
	// 		defer wg.Done() // ✅ har goroutine `Done()` qiladi

	// 		ctx := context.Background()

	// 		result, err := testStore.TranferTx(ctx, TransferTxParams{
	// 			FromAccountID: account1.ID,
	// 			ToAccountID:   account2.ID,
	// 			Amount:        amount,
	// 		})

	// 		errs <- err
	// 		results <- result
	// 	}(i) // ✅ i ni parametr sifatida uzatyapmiz
	// }

	// // ✅ test tugashini kutish
	// go func() {
	// 	wg.Wait()
	// 	close(errs)
	// 	close(results)
	// }()

	// existed := make(map[int]bool)

	// //check
	// for i := 1; i <= n; i++ {
	// 	err := <-errs
	// 	require.NoError(t, err)
	// 	result := <-results
	// 	require.NotEmpty(t, result)

	// 	//transfer check
	// 	transfer := result.Transfer
	// 	require.NotEmpty(t, transfer)
	// 	require.NotZero(t, transfer.ID)
	// 	require.Equal(t, transfer.FromAccountID, account1.ID)
	// 	require.Equal(t, transfer.ToAccountID, account2.ID)
	// 	require.Equal(t, transfer.Amount, amount)
	// 	require.NotZero(t, transfer.CreatedAt)

	// 	_, err = testStore.GetTransfer(context.Background(), transfer.ID)
	// 	require.NoError(t, err)

	// 	//entry from
	// 	entryFrom := result.FromEntry
	// 	require.NotEmpty(t, entryFrom)
	// 	require.NotZero(t, entryFrom.ID)
	// 	require.NotZero(t, entryFrom.CreatedAt)
	// 	require.Equal(t, entryFrom.AccountID, account1.ID)
	// 	require.Equal(t, entryFrom.Amount, -amount)

	// 	_, err = testStore.GetEntry(context.Background(), entryFrom.ID)
	// 	require.NoError(t, err)

	// 	//entry to
	// 	entryTo := result.ToEntry
	// 	require.NotEmpty(t, entryTo)
	// 	require.NotZero(t, entryTo.ID)
	// 	require.NotZero(t, entryTo.CreatedAt)
	// 	require.Equal(t, entryTo.AccountID, account2.ID)
	// 	require.Equal(t, entryTo.Amount, amount)

	// 	_, err = testStore.GetEntry(context.Background(), entryTo.ID)
	// 	require.NoError(t, err)

	// 	//account
	// 	accountFrom := result.FromAccount
	// 	require.NotEmpty(t, accountFrom)
	// 	require.Equal(t, accountFrom.ID, account1.ID)

	// 	accountTo := result.ToAccount
	// 	require.NotEmpty(t, accountTo)
	// 	require.Equal(t, accountTo.ID, account2.ID)

	// 	diff1 := account1.Balance - accountFrom.Balance
	// 	diff2 := accountTo.Balance - account2.Balance

	// 	require.Equal(t, diff1, diff2)
	// 	require.True(t, diff1 > 0)
	// 	require.True(t, diff1%amount == 0)

	// 	k := int(diff1 / amount)
	// 	require.True(t, k >= 1 && k <= n)
	// 	require.NotContains(t, existed, k)
	// 	existed[k] = true
	// }

	// updatedAccount1, err := testStore.GetAccount(context.Background(), account1.ID)
	// require.NoError(t, err)

	// updatedAccount2, err := testStore.GetAccount(context.Background(), account2.ID)
	// require.NoError(t, err)

	// require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	// require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)

}
