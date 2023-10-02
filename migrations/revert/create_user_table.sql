-- Revert twitterLikeGo:create_user_table from pg

BEGIN;

DROP TABLE "users";

COMMIT;
