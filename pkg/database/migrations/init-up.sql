BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
	"id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('u-', uuid_generate_v4())),
	"name" VARCHAR(255) NOT NULL,
	"type" VARCHAR(255) DEFAULT 'BUSINESS',
	"status" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

CREATE TABLE IF NOT EXISTS user_balances(
	"id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('ub-', uuid_generate_v4())),
	"user_id" VARCHAR(255) NOT NULL REFERENCES users(id),
	"currency" VARCHAR(255) NOT NULL,
	"amount" NUMERIC(35,15) NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

CREATE TABLE IF NOT EXISTS transfers(
	"id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('tf-', uuid_generate_v4())),
	"user_id" VARCHAR(255) NOT NULL REFERENCES users(id),
    "destination_account" VARCHAR(255) NOT NULL,
    "bank_code" VARCHAR(255) NOT NULL,
	"currency" VARCHAR(255) NOT NULL,
	"amount" NUMERIC(35,15) NOT NULL,
	"notes" VARCHAR(255),
	"reference_id" VARCHAR(255) NOT NULL,
    "status" VARCHAR(255) DEFAULT 'ACCEPTED',
    "metadata" JSONB DEFAULT '{}',
	"bank_transfer_id" VARCHAR(255),
	"created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

CREATE UNIQUE INDEX "transfers_userID_refID"
    on transfers (user_id, reference_id);

INSERT INTO public.users (id, 
						name, 
						type, 
						status, 
						created_at, 
						updated_at, 
						deleted_at) 
	VALUES ('u-515ffa15-43af-4773-b0d5-e80a53613a4b', 
			'Fahmy', 
			'PERSONAL', 
			'ACTIVE', 
			'2024-05-16 04:54:56.164854', 
			'2024-05-16 04:54:56.164854', 
			'2024-05-16 04:54:56.164854');

INSERT INTO user_balances (id, 
						user_id, 
						currency, 
						amount, 
						created_at, 
						updated_at, 
						deleted_at) 
	VALUES ('ub-796e1d56-4465-46ae-8d3b-653117f8ed84', 
			'u-515ffa15-43af-4773-b0d5-e80a53613a4b', 
			'IDR', 
			1000000, 
			'2024-05-16 04:55:07.151989', 
			'2024-05-16 12:10:36.175787', 
			'2024-05-16 04:55:07.151989');

COMMIT;