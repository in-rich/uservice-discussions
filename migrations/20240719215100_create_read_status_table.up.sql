CREATE TABLE read_statuses (
    id                     UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    user_id                VARCHAR(255) NOT NULL,
    team_id                VARCHAR(255) NOT NULL,
    public_identifier      VARCHAR(255) NOT NULL,
    target                 message_target NOT NULL,
    latest_read_message_id UUID NOT NULL,

    CONSTRAINT unique_status_per_discussion UNIQUE (user_id, team_id, public_identifier, target)
);

--bun:split

CREATE INDEX read_status_author_id_index ON read_statuses (user_id, team_id, public_identifier, target);
