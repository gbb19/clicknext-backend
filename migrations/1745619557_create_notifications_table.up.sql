CREATE TYPE notify_type AS ENUM ('task', 'board');

CREATE TABLE IF NOT EXISTS notifications (
    notify_id SERIAL PRIMARY KEY ,
    message VARCHAR(255) NOT NULL ,
    is_read BOOLEAN DEFAULT FALSE,
    type notify_type NOT NULL ,
    user_id INT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (user_id)
);