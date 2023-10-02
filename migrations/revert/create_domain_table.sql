-- Revert twitterLikeGo:create_domain_table from pg

BEGIN;

DROP DOMAIN email_d, name_d;

COMMIT;
