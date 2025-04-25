CREATE TABLE IF NOT EXISTS tags (
    tag_id SERIAL PRIMARY KEY ,
    name VARCHAR(255) NOT NULL ,
    created_by INT NOT NULL ,

    FOREIGN KEY (created_by) REFERENCES users (user_id)
);