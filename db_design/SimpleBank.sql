CREATE TABLE "users" (
  "username" varchar PRIMARY KEY
  "hashed_password" varchar NOT NULL
  "full_name" varchar NOT NULL
  "email" varchar UNIQUE NOT NULL
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
  "created_at" timestamptz DEFAULT (now())
)

--records all bank accounts
CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY, --unique account ID, auto incremental primary key
  "owner" varchar NOT NULL, --name of account holder
  "balance" decimal NOT NULL, --amount of available money
  "currency" varchar NOT NULL, --type of currency
  "created_at" timestamptz NOT NULL DEFAULT (now()) --time when account is created
);

--records all changes to account balance
CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY, --unique entry ID, auto incremental primary key
  "account_id" bigint NOT NULL, --account where change is made
  "amount" decimal NOT NULL, --change of balance
  "created_at" timestamptz NOT NULL DEFAULT (now()) --time when entry is created
);

--records all transfers made between two accounts
--only allows transfers within WeBank, transfer from/to other banks may be implemented in future as needed
CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY, --unique transfer ID, auto incremental primary key
  "from_account_id" bigint NOT NULL, --account making the transfer
  "to_account_id" bigint NOT NULL, --account receiving the transfer
  "amount" decimal NOT NULL, --amount of money being moved
  "created_at" timestamptz NOT NULL DEFAULT (now()) --time when transfer is made
);

--may want to search for accounts by owner name
CREATE INDEX ON "accounts" ("owner");

CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");

--may want to list all entries of a specific accoount
CREATE INDEX ON "entries" ("account_id");

--may want to list all entries going in OR out of an account OR all transfers between two accounts
CREATE INDEX ON "transfers" ("from_account_id");
CREATE INDEX ON "transfers" ("to_account_id");
CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

--amount of balance change can be positive (deposits ...) or negative (withdrawal ...)
COMMENT ON COLUMN "entries"."amount" IS 'can be positive or negative';

--amount of a transfer can only be positive to make sense
COMMENT ON COLUMN "transfers"."amount" IS 'must be postive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

--ref to account where change is made (one to many from accounts to entries)
ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

-- ref to account making the transfer (one to many from accounts to transfers)
ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

--ref to account receiving the transfer (one to many from accounts to transfers)
ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
