CREATE TABLE IF NOT EXISTS accounts (
    account_id BIGINT PRIMARY KEY,
    balance NUMERIC(19, 4) NOT NULL,               -- Current balance of the account (e.g., 1250.7500)
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), -- Timestamp when the account was created
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), -- Timestamp when the account was last updated
    deleted_at TIMESTAMPTZ                         -- Timestamp for soft deletion (NULL if not deleted)
);
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_accounts_updated_at
BEFORE UPDATE ON accounts
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();


CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY,             
    source_account_id BIGINT NOT NULL,                       -- The account from which funds are debited
    destination_account_id BIGINT NOT NULL,                  -- The account to which funds are credited
    amount NUMERIC(19, 4) NOT NULL CHECK (amount > 0),       -- The amount of the transaction, must be positive
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),           -- Timestamp when the transaction was initiated
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),           -- Timestamp when the transaction was last updated
    deleted_at TIMESTAMPTZ,                                  -- Timestamp for soft deletion (NULL if not deleted)

    -- Foreign key constraints (now referencing BIGINT IDs)
    CONSTRAINT fk_source_account
        FOREIGN KEY (source_account_id)
        REFERENCES accounts (account_id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_destination_account
        FOREIGN KEY (destination_account_id)
        REFERENCES accounts (account_id)
        ON DELETE RESTRICT,

    -- Optional: Ensure source and destination accounts are different for transfers
    CONSTRAINT chk_different_accounts
        CHECK (source_account_id <> destination_account_id)
);
-- speeding up query with indexes
CREATE INDEX idx_transactions_source_account_id ON transactions (source_account_id);
CREATE INDEX idx_transactions_destination_account_id ON transactions (destination_account_id);

CREATE OR REPLACE TRIGGER update_transactions_updated_at
BEFORE UPDATE ON transactions
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();