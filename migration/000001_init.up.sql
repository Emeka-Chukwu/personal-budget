CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar NOT NULL UNIQUE,
  "phone" varchar NOT NULL,
  "is_verified" boolean NOT NULL DEFAULT false,
  "password" varchar NOT NULL,
  "bvn" varchar, -- fix the colon to a comma
  "pin" varchar,
  "profile_url" varchar, -- add a comma here
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "is_suspended" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id")
);

CREATE TABLE "accounts" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "name" varchar,
  "number" varchar,
  "user_id" uuid,
  "balance" int,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "wallets" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "user_id" uuid,
  "balance" int,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "sessions" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "refresh_token" varchar,
  "is_blocked" boolean,
  "user_id" uuid,
  "user_agent" varchar,
  "client_ip" varchar,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "transactions" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "type" varchar,
  "status" varchar,
  "user_id" uuid,
  "reference" varchar,
  "amount" int,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "verifications" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "secret_code" varchar,
  "is_used" boolean,
  "user_id" uuid,
  "email" varchar,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "payment_pins" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "pin" varchar,
  "user_id" uuid,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "scheduled_payments" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "user_id" uuid,
  "amount" uuid,
  "paydate" timestamptz,
  "is_completed" boolean DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "verifications" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "payment_pins" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "scheduled_payments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "wallets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
