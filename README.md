# selaginella

Cross-chain interoperability relayer

## How to develop 🤪

Clone project:

```shell
git clone https://github.com/bridge-alchemy/selaginella.git
cd selaginella
```

Modify the `[database]` configuration of `config.toml` file.

Compile project:

```shell
make build
```

Run:

```shell
./selaginella grpc
./selaginella exporter
```

## How to use 🤔

### 1. Deposit

```grpc
GRPC grpc://localhost:8080/deposit.DepositService/AddDeposit

{
"wallet_guid": "wallet_guid_123",
"user_guid": "user_guid_456",
"asset_guid": "asset_guid_789",
"transaction_guid": "transaction_guid_abc",
"from_address": "from_address_xyz",
"to_address": "to_address_def",
"amount": "100.0",
"tx_hash": "tx_hash_123",
"tx_fee": "1.0",
"status": 1
}
```
