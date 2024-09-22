CREATE TYPE message_target AS ENUM ('user', 'company');

--bun:split

CREATE TABLE messages (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    author_id         VARCHAR(255) NOT NULL,
    team_id           VARCHAR(255) NOT NULL,
    public_identifier VARCHAR(255) NOT NULL,
    target            message_target NOT NULL,
    
    content           TEXT NOT NULL,
    created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

--bun:split

CREATE INDEX messages_per_target_per_team ON messages (team_id, target, public_identifier);
