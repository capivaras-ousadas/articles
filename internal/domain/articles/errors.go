package articles

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrArticleNotFound = Error{
		Code:    404,
		Message: "article not found",
	}
	ErrArticleTitleRequired = Error{
		Code:    400,
		Message: "article title is required",
	}
	ErrArticleAuthorRequired = Error{
		Code:    400,
		Message: "article author is required",
	}
	ErrArticleAlreadyExists = Error{
		Code:    401,
		Message: "article already exists",
	}
	ErrArticleIDNotProvided = Error{
		Code:    400,
		Message: "article id is required",
	}
)
