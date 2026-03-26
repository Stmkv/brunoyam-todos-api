CREATE TYPE status as ENUM('new', 'in_process', 'completed');

CREATE TABLE IF NOT EXISTS tasks
(
    tid         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title       TEXT NOT NULL,
    description TEXT,
    status      status NOT NULL DEFAULT 'new'
);