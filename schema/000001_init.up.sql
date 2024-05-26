CREATE TABLE IF NOT EXISTS Post (
    id SERIAL PRIMARY KEY,
    userId INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    commentsEnabled BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS Comment (
    id SERIAL PRIMARY KEY,
    userId INTEGER NOT NULL,
    postId INTEGER NOT NULL REFERENCES Post(id) ON DELETE CASCADE,
    parentId INTEGER,
    content TEXT NOT NULL
);

CREATE INDEX idx_post_userId ON Post(userId);
CREATE INDEX idx_comment_userId ON Comment(userId);
CREATE INDEX idx_comment_postId ON Comment(postId);
CREATE INDEX idx_comment_parentId ON Comment(parentId);
