-- Drop foreign key constraints

-- Drop constraints in transactions table
ALTER TABLE "transactions" DROP CONSTRAINT IF EXISTS "transactions_user_id_fkey";

-- Drop constraints in sessions table
ALTER TABLE "sessions" DROP CONSTRAINT IF EXISTS "sessions_user_id_fkey";

-- Drop constraints in verifications table
ALTER TABLE "verifications" DROP CONSTRAINT IF EXISTS "verifications_user_id_fkey";

-- Drop constraints in payment_pins table
ALTER TABLE "payment_pins" DROP CONSTRAINT IF EXISTS "payment_pins_user_id_fkey";

-- Drop constraints in scheduled_payments table
ALTER TABLE "scheduled_payments" DROP CONSTRAINT IF EXISTS "scheduled_payments_user_id_fkey";
ALTER TABLE "scheduled_payments" DROP CONSTRAINT IF EXISTS "scheduled_payments_amount_fkey";

-- Drop foreign key constraints in accounts table
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_user_id_fkey";

ALTER TABLE "wallets" DROP CONSTRAINT IF EXISTS "wallets_user_id_fkey";

-- Drop tables
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "accounts";
DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "transactions";
DROP TABLE IF EXISTS "verifications";
DROP TABLE IF EXISTS "payment_pins";
DROP TABLE IF EXISTS "scheduled_payments";
DROP TABLE IF EXISTS "wallets";
