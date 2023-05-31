BEGIN;
CREATE TABLE IF NOT EXISTS public.operations (
	id text NOT NULL,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	deleted_at timestamp with time zone,
	"type" text,
	"cost" bigint,
	PRIMARY KEY(id)
);
COMMIT;

BEGIN;
CREATE TABLE IF NOT EXISTS public.records (
	id text,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	deleted_at timestamp with time zone,
	operation_id text NOT NULL,
	user_id text NOT NULL,
	amount bigint,
	user_balance bigint,
	operation_response text,
	PRIMARY KEY(id)
);
COMMIT;

BEGIN;
CREATE TABLE IF NOT EXISTS public.users (
	id text NOT NULL,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	deleted_at timestamp with time zone,
	email text,
	"password" text,
	balance bigint,
	PRIMARY KEY(id)
);
COMMIT;

ALTER TABLE IF EXISTS public.records
	ADD CONSTRAINT fk_users_records
	FOREIGN KEY (user_id)
	REFERENCES public.users (id);

ALTER TABLE IF EXISTS public.records
	ADD CONSTRAINT fk_operations_records
	FOREIGN KEY (operation_id)
	REFERENCES public.operations (id);

