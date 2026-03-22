-- Create "users" table
CREATE TABLE "public"."users" (
  "user_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "full_name" text NOT NULL,
  "email_id" text NOT NULL,
  "password_hash" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("user_id")
);
-- Create "auth" table
CREATE TABLE "public"."auth" (
  "auth_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NULL,
  "email_id" text NOT NULL,
  "token" text NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "is_expired" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("auth_id"),
  CONSTRAINT "fk_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "mines" table
CREATE TABLE "public"."mines" (
  "mine_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "owner_id" uuid NOT NULL,
  "mine_name" character varying(255) NOT NULL,
  "purity_score" integer NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NOT NULL,
  "updated_at" timestamptz NULL,
  "updated_by" uuid NULL,
  PRIMARY KEY ("mine_id"),
  CONSTRAINT "fk_user" FOREIGN KEY ("owner_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "gems" table
CREATE TABLE "public"."gems" (
  "gem_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "mine_id" uuid NOT NULL,
  "title" character varying(255) NOT NULL,
  "file_url" text NOT NULL,
  "teaser_text" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NOT NULL,
  "updated_at" timestamptz NULL,
  "updated_by" uuid NULL,
  PRIMARY KEY ("gem_id"),
  CONSTRAINT "fk_mine" FOREIGN KEY ("mine_id") REFERENCES "public"."mines" ("mine_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "drops" table
CREATE TABLE "public"."drops" (
  "drop_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "gem_id" uuid NOT NULL,
  "drop_code" character varying(50) NOT NULL,
  "label" character varying(50) NOT NULL,
  "max_claims" integer NOT NULL,
  "current_claims" integer NOT NULL DEFAULT 0,
  "expires_at" timestamptz NULL,
  "access_level" character varying(50) NOT NULL DEFAULT 'public',
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" uuid NOT NULL,
  "updated_at" timestamptz NULL,
  "updated_by" uuid NULL,
  PRIMARY KEY ("drop_id"),
  CONSTRAINT "fk_gem" FOREIGN KEY ("gem_id") REFERENCES "public"."gems" ("gem_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "chk_access_level" CHECK ((access_level)::text = ANY ((ARRAY['public'::character varying, 'subscriber'::character varying])::text[]))
);
-- Create "extractions" table
CREATE TABLE "public"."extractions" (
  "extraction_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "gem_id" uuid NOT NULL,
  "drop_id" uuid NOT NULL,
  "tract_id" character varying(50) NULL,
  "extracted_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("extraction_id"),
  CONSTRAINT "extractions_user_id_gem_id_key" UNIQUE ("user_id", "gem_id"),
  CONSTRAINT "fk_drop_extraction" FOREIGN KEY ("drop_id") REFERENCES "public"."drops" ("drop_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_gem_extraction" FOREIGN KEY ("gem_id") REFERENCES "public"."gems" ("gem_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_prospector" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
