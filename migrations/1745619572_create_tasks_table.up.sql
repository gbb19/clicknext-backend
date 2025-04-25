CREATE TABLE IF NOT EXISTS tasks
(
    task_id    SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    position   INT          NOT NULL,
    due_date   TIMESTAMP,
    start_date TIMESTAMP    NOT NULL,
    column_id  INT NOT NULL ,
    created_by INT NOT NULL ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (created_by) REFERENCES users (user_id),
    FOREIGN KEY (column_id) REFERENCES columns (column_id)

);
