-- modify "user_profiles" table
ALTER TABLE "user_profiles" ADD COLUMN "full_name" character varying(255) NOT NULL DEFAULT '';
