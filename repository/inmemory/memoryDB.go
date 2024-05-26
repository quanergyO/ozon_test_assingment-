package inmemory

type Comment struct {
	UserId   int
	PostId   int
	ParentId int
	Content  string
}

type Post struct {
	UserId         int
	Title          string
	Content        string
	CommentsEnable bool
}

type MemoryDB struct {
	Comments map[int]Comment
	Posts    map[int]Post
}
