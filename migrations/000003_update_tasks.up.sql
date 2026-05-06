ALTER TABLE tasks
    ADD COLUMN user_id UUID;

ALTER TABLE tasks
    ADD CONSTRAINT fk_tasks_users
        FOREIGN KEY (user_id)
            REFERENCES users(uid)
            ON DELETE CASCADE;