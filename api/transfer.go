package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	db "github.com/berdiyoroff/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type tranferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(c *gin.Context) {
	var req tranferRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(c, req.FromAccountID, req.Currency) {
		return
	}

	if !server.validAccount(c, req.ToAccountID, req.Currency) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TranferTx(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(c *gin.Context, accountId int64, currency string) bool {

	account, err := server.store.GetAccount(c, accountId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return false
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return false
		}
	}

	if account.Currency != currency {
		err = fmt.Errorf("account: [%d] currency mismach: %s vs %s", account.ID, account.Currency, currency)
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
