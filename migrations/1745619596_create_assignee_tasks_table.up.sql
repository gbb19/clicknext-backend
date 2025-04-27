CREATE TABLE IF NOT EXISTS assignee_tasks(
    assignee_task_id SERIAL PRIMARY KEY ,
    assignee_id INT NOT NULL ,
    task_id INT NOT NULL ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (assignee_id) REFERENCES users (user_id),
    FOREIGN KEY (task_id) REFERENCES tasks (task_id)
);
