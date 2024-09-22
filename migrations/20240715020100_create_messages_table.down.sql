DROP INDEX IF EXISTS messages_per_target_per_team;

--bun:split

DROP TABLE IF EXISTS messages;

--bun:split

DROP TYPE IF EXISTS message_target;
