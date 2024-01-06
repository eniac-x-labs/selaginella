package deposit

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
)

type DepositService struct {
	Db *sql.DB
}

func (s *DepositService) AddDeposit(ctx context.Context, req *AddDepositRequest) (*AddDepositResponse, error) {
	// Implement the logic of inserting or updating data into the PostgreSQL database here.
	// Use s.Db to execute database operations.

	insertSQL := `
    INSERT INTO deposits (
        guid,
        wallet_guid,
        user_guid,
        asset_guid,
        transaction_guid,
        from_address,
        to_address,
        amount,
        tx_hash,
        tx_fee,
        status
    )
    VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
    )
`
	log.Info("Inserting deposit record", "sql", insertSQL)

	guid := generateGUID()

	log.Debug("Generated GUID", "guid", guid)
	log.Debug("Wallet GUID", "wallet_guid", req.WalletGuid)
	log.Debug("User GUID", "user_guid", req.UserGuid)
	log.Debug("Asset GUID", "asset_guid", req.AssetGuid)
	log.Debug("Transaction GUID", "transaction_guid", req.TransactionGuid)
	log.Debug("From address", "from_address", req.FromAddress)
	log.Debug("To address", "to_address", req.ToAddress)
	log.Debug("Amount", "amount", req.Amount)
	log.Debug("Transaction hash", "tx_hash", req.TxHash)
	log.Debug("Transaction fee", "tx_fee", req.TxFee)
	log.Debug("Status", "status", req.Status)

	_, err := s.Db.ExecContext(ctx, insertSQL,
		guid,
		req.WalletGuid,
		req.UserGuid,
		req.AssetGuid,
		req.TransactionGuid,
		req.FromAddress,
		req.ToAddress,
		req.Amount,
		req.TxHash,
		req.TxFee,
		req.Status,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert deposit record: %v", err)
	}
	log.Info("Deposit record inserted", "guid", guid)

	return &AddDepositResponse{Guid: guid}, nil
}

func generateGUID() string {
	// 你需要根据实际情况生成唯一的 GUID
	return "some_generated_guid"
}

func (*DepositService) mustEmbedUnimplementedDepositServiceServer() {}
