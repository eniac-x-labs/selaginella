CREATE TABLE IF NOT EXISTS address (
    guid             VARCHAR PRIMARY KEY,
    wallet_guid      VARCHAR NOT NULL,
    user_guid        VARCHAR NOT NULL,
    wallet_type      SMALLINT NOT NULL,
    address          INTEGER NOT NULL,
    privateKey       VARCHAR NOT NULL,
 );
CREATE INDEX IF NOT EXISTS address_guid ON address(guid);
CREATE INDEX IF NOT EXISTS address_wallet_guid ON address(wallet_guid);
CREATE INDEX IF NOT EXISTS address_user_guid ON address(user_guid);
CREATE INDEX IF NOT EXISTS address_address ON address(address);

CREATE TABLE IF NOT EXISTS asset (
    guid             VARCHAR PRIMARY KEY,
    name             VARCHAR NOT NULL,
    contract         VARCHAR NOT NULL UNIQUE,
    decimal          UINT256 NOT NULL DEFAULT 18,
);
CREATE INDEX IF NOT EXISTS asset_guid ON asset(guid);
CREATE INDEX IF NOT EXISTS asset_contract ON asset(contract);

CREATE TABLE IF NOT EXISTS balance (
    guid             VARCHAR PRIMARY KEY,
    asset_guid       VARCHAR NOT NULL,
    address_guid     VARCHAR NOT NULL,
    wallet_type      SMALLINT NOT NULL,
    balance          VARCHAR NOT NULL,
);
CREATE INDEX IF NOT EXISTS balance_guid ON balance(guid);
CREATE INDEX IF NOT EXISTS balance_asset_guid ON balance(asset_guid);
CREATE INDEX IF NOT EXISTS balance_address_guid ON balance(address_guid);

CREATE TABLE IF NOT EXISTS block (
    hash        VARCHAR PRIMARY KEY,
    parent_hash VARCHAR NOT NULL UNIQUE,
    number      UINT256 NOT NULL UNIQUE,
    timestamp   INTEGER NOT NULL UNIQUE CHECK (timestamp > 0),
    rlp_bytes   VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS block_timestamp ON block(timestamp);
CREATE INDEX IF NOT EXISTS block_number ON block(number);

CREATE TABLE IF NOT EXISTS deposits (
    guid             VARCHAR PRIMARY KEY,
    wallet_guid      VARCHAR NOT NULL,
    user_guid        VARCHAR NOT NULL,
    asset_guid       VARCHAR NOT NULL,
    transaction_guid VARCHAR NOT NULL,
    from_address     VARCHAR NOT NULL,
    to_address       VARCHAR NOT NULL,
    amount           VARCHAR NOT NULL,
    tx_hash          VARCHAR NOT NULL,
    tx_fee           VARCHAR NOT NULL,
    status           SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS deposits_guid ON deposits(guid);
CREATE INDEX IF NOT EXISTS deposits_wallet_guid ON deposits(wallet_guid);
CREATE INDEX IF NOT EXISTS deposits_user_guid ON deposits(user_guid);
CREATE INDEX IF NOT EXISTS deposits_asset_guid ON deposits(asset_guid);
CREATE INDEX IF NOT EXISTS deposits_transaction_guid ON deposits(transaction_guid);
CREATE INDEX IF NOT EXISTS deposits_tx_hash ON deposits(tx_hash);

CREATE TABLE IF NOT EXISTS withdraws (
    guid             VARCHAR PRIMARY KEY,
    wallet_guid      VARCHAR NOT NULL,
    user_guid        VARCHAR NOT NULL,
    asset_guid       VARCHAR NOT NULL,
    transaction_guid VARCHAR NOT NULL,
    from_address     VARCHAR NOT NULL,
    to_address       VARCHAR NOT NULL,
    amount           VARCHAR NOT NULL,
    tx_hash          VARCHAR NOT NULL,
    raw_tx           VARCHAR NOT NULL,
    receive_tx_fee   VARCHAR NOT NULL,
    chain_tx_fee     VARCHAR NOT NULL,
    status           SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS withdraws_guid ON withdraws(guid);
CREATE INDEX IF NOT EXISTS withdraws_wallet_guid ON withdraws(wallet_guid);
CREATE INDEX IF NOT EXISTS withdraws_user_guid ON withdraws(user_guid);
CREATE INDEX IF NOT EXISTS withdraws_asset_guid ON withdraws(asset_guid);
CREATE INDEX IF NOT EXISTS withdraws_transaction_guid ON withdraws(transaction_guid);
CREATE INDEX IF NOT EXISTS withdraws_tx_hash ON withdraws(tx_hash);


CREATE TABLE IF NOT EXISTS transactions (
    guid             VARCHAR PRIMARY KEY,
    block_guid       VARCHAR NOT NULL,
    tx_hash          VARCHAR NOT NULL,
    amount           VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS transactions_guid ON transactions(guid);
CREATE INDEX IF NOT EXISTS transactions_block_guid ON transactions(block_guid);
CREATE INDEX IF NOT EXISTS transactions_tx_hash ON transactions(tx_hash);

