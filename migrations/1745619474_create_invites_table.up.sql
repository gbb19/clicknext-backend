CREATE TYPE invite_status AS ENUM ('pending', 'accepted', 'declined');

CREATE TABLE IF NOT EXISTS invites (
    invite_id SERIAL PRIMARY KEY ,
    status invite_status DEFAULT 'pending',
    inviter_id INT NOT NULL ,
    invitee_id INT NOT NULL ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (inviter_id) REFERENCES users (user_id),
    FOREIGN KEY (invitee_id) REFERENCES users (user_id)
);