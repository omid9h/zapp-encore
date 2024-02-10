-- create "user_profiles" table
CREATE TABLE "user_profiles" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  "email" character varying(100) NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "is_deleted" boolean NULL DEFAULT false,
  PRIMARY KEY ("id")
);
