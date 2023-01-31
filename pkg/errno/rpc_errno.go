package errno

var (
	ErrorGetListFail = NewErrNo(2001, "GetListFail")
	ErrorCreateFail  = NewErrNo(2002, "CreateFail")
	ErrorUpdateFail  = NewErrNo(2003, "UpdateFail")
	ErrorDeleteFail  = NewErrNo(2004, "DeleteFail")
	ErrorCountFail   = NewErrNo(2005, "Count Fail")

	UserAlreadyExistErr		= NewErrNo(3001, "User Already Exist")
	UserPwdErr				= NewErrNo(3002, "User Password Wrong")
	UserNotExist			= NewErrNo(3003, "User Not Exist")
	
	CompCommentErr			= NewErrNo(4001, "Comment Error")

	VideoPublishErr			= NewErrNo(5001, "Publish Errror")
	
)