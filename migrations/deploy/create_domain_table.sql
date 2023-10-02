-- Deploy twitterLikeGo:create_domain_table to pg

BEGIN;

CREATE DOMAIN email_d AS TEXT
CHECK (VALUE ~* '^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$');

CREATE DOMAIN name_d AS TEXT
CHECK (VALUE ~* '^[A-Z]{2,}([ -][A-Z]{2,})?$');

COMMIT;
