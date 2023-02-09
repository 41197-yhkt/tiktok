package errno

var (
	ErrorGetListFail = NewErrNo(2001, "GetListFail")
	ErrorCreateFail  = NewErrNo(2002, "CreateFail")
	ErrorUpdateFail  = NewErrNo(2003, "UpdateFail")
	ErrorDeleteFail  = NewErrNo(2004, "DeleteFail")
	ErrorCountFail   = NewErrNo(2005, "Count Fail")

	UserAlreadyExistErr           = NewErrNo(3001, "User Already Exist")
	UserPwdErr                    = NewErrNo(3002, "User Password Wrong")
	UserNotExist                  = NewErrNo(3003, "User Not Exist")
	UserFollowRelationExistErr    = NewErrNo(3004, "FollowRelation Already Exist")
	UserFollowRelationNotExistErr = NewErrNo(3005, "FollowRelation Not Exist")

	CompCommentErr     = NewErrNo(4001, "Comment Error")
	CommentNotExistErr = NewErrNo(4002, "Comment Not Exist")

	VideoPublishErr  = NewErrNo(5001, "Publish Errror")
	VideoNotExistErr = NewErrNo(5002, "Video Not Exist")
)
