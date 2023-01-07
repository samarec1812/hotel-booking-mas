CREATE TABLE bank_account
(
    id       serial       not null unique,
    name     varchar(255) not null,
    password varchar(255) not null,
    secure_code int not null,
    balance BIGINT,
    CONSTRAINT balance_check CHECK (balance >= 0)
);

CREATE TABLE transactions
(
    id            SERIAL    NOT NULL unique,
    bank_acc_id   INT       NOT NULL,
    start_balance BIGINT    NOT NULL,
    end_balance   BIGINT    NOT NULL,
    amount        BIGINT    NOT NULL,
    description   VARCHAR(255)  NULL,
    date          TIMESTAMP NOT NULL,
    FOREIGN KEY (bank_acc_id) REFERENCES bank_account(id) on delete cascade,
    CONSTRAINT balance_check CHECK (start_balance >= 0 AND
                                    end_balance >= 0 AND
                                    amount >= 0)
);
