-- Revert twitterLikeGo:create_user_table from pg

BEGIN;

DROP FUNCTION "insert_user_json"(json);

DROP TABLE "users";

COMMIT;
