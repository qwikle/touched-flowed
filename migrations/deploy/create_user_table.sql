-- Deploy twitterLikeGo:create_user_table to pg

BEGIN;

CREATE TABLE "users" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    "first_name" name_d NOT NULL,
    "last_name" name_d NOT NULL,
    "email" email_d UNIQUE NOT NULL,
    "password" TEXT NOT NULL
    );

COMMIT;
