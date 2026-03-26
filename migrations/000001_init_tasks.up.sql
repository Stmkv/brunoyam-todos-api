CREATE TABLE tasks
(
    tid         TEXT PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT,
    status      INT  NOT NULL
);