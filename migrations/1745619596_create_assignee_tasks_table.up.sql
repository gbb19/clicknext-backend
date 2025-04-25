CREATE TABLE IF NOT EXISTS assignee_tasks(
    assignee_tasks_id SERIAL PRIMARY KEY ,
    assignee_id INT NOT NULL ,
    task_id INT NOT NULL ,

    FOREIGN KEY (assignee_id) REFERENCES users (user_id),
    FOREIGN KEY (task_id) REFERENCES tasks (task_id)
);
