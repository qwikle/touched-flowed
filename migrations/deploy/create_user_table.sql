-- Deploy twitterLikeGo:create_user_table to pg

BEGIN;

CREATE TABLE "users" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    "first_name" name_d NOT NULL,
    "last_name" name_d NOT NULL,
    "email" email_d UNIQUE NOT NULL,
    "password" TEXT NOT NULL
    );


CREATE FUNCTION "insert_user_json" (json JSON) RETURNS BIGINT AS $$
DECLARE
    "user_id" BIGINT;
BEGIN
    INSERT INTO "users" ("first_name", "last_name", "email", "password")
    VALUES (json->>'first_name', json->>'last_name', json->>'email', json->>'password')
    RETURNING "id" INTO "user_id";
    RETURN "user_id";
END;
$$ LANGUAGE plpgsql;

COMMIT;
