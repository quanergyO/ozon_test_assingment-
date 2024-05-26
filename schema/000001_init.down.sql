DROP INDEX IF EXISTS idx_comment_parentId;
DROP INDEX IF EXISTS idx_comment_postId;
DROP INDEX IF EXISTS idx_comment_userId;
DROP INDEX IF EXISTS idx_post_userId;

DROP TABLE IF EXISTS Comment;

DROP TABLE IF EXISTS Post;