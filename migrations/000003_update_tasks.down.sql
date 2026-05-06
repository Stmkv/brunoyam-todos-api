ALTER TABLE tasks
    DROP CONSTRAINT fk_tasks_users;

ALTER TABLE tasks
    DROP COLUMN user_id;