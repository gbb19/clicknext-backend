CREATE TABLE IF NOT EXISTS columns (
    column_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    position INT NOT NULL,
    color VARCHAR(255) DEFAULT '#5D5D5D' NOT NULL,
    created_by INT NOT NULL,
    board_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by) REFERENCES users (user_id),
    FOREIGN KEY (board_id) REFERENCES boards (board_id)
);