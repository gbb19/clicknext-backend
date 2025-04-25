CREATE TABLE IF NOT EXISTS board_members (
    member_id SERIAL PRIMARY KEY ,
    user_id INT NOT NULL ,
    board_id INT NOT NULL ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (board_id) REFERENCES boards (board_id)
);