CREATE TABLE IF NOT EXISTS task_tags (
    task_tag_id SERIAL PRIMARY KEY ,
    task_id INT NOT NULL ,
    tag_id INT NOT NULL ,

    FOREIGN KEY (task_id) REFERENCES tasks (task_id),
    FOREIGN KEY (tag_id) REFERENCES tags (tag_id)
);