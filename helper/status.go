package helper

// 定义一些系统常用的错误码

const (
	// 普通错误码
	BindModelErr      = 20000
	NoneParamErr      = 20001
	LoginStatusErr    = 20002
	SaveStatusErr     = 20003
	DeleteStatusErr   = 20004
	ExistUsernameErr  = 20005
	ExistTelephoneErr = 20006
	PasswordErr       = 20007
	ArticlesPageErr   = 20008
	ArticleGetErr     = 20009
	ArticleAddErr     = 20010
	ArticleModifyErr  = 20011
	ArticleDelErr     = 20012
	ArticleLikeErr    = 20013
	ArticleUnLikeErr  = 20014
	ArticleStarErr    = 20015
	ArticleUnStarErr  = 20016
	ArticleCommentErr = 20017
	CommentAddErr     = 20018
	CommentModifyErr  = 20019
	CommentLikeErr    = 20020
	CommentUnLikeErr  = 20021
	FollowErr         = 20022
	UnFollowErr       = 20023
	LoginStatusOK     = 20200
	SaveStatusOK      = 20201
	DeleteStatusOK    = 20202
	ArticlesPageOk    = 20203
	ArticleOk         = 20204
	ArticleLikeOk     = 20205
	ArticleUnLikeOk   = 20206
	ArticleStarOk     = 20207
	ArticleUnStarOk   = 20208
	ArticleCommentOk  = 20209
	CommentAddOk      = 20210
	CommentModifyOk   = 20211
	CommentLikeOk     = 20212
	CommentUnLikeOk   = 20213
	FollowOk          = 20214
	UnFollowOk        = 20215
	SaveObjIsNil      = 20400
	DeleteObjIsNil    = 20401
	UpdateObjIsNil    = 20402

	// 正则相关错误码
	FixLessZeroErr    = 20500
	MaxLessZeroErr    = 20501
	MinThanMaxErr     = 20502
	MediumPasswordErr = 20503
	StrongPasswordErr = 20504
	ChineseNameErr    = 20505
	EnglishNameErr    = 20506
)

var statusText = map[int]string{
	// 普通错误信息
	BindModelErr:      "模型封装异常！",
	NoneParamErr:      "无有效参数",
	LoginStatusErr:    "用户名或密码错误!",
	LoginStatusOK:     "登陆成功！",
	SaveStatusOK:      "保存成功！",
	SaveStatusErr:     "保存失败！",
	SaveObjIsNil:      "保存的对象为空！",
	DeleteStatusOK:    "删除成功！",
	DeleteStatusErr:   "删除失败！",
	DeleteObjIsNil:    "删除的记录不存在！",
	UpdateObjIsNil:    "修改的记录不存在！",
	ExistUsernameErr:  "用户名已存在！",
	ExistTelephoneErr: "手机号已存在！",
	ArticlesPageErr:   "文章分页查找错误！",
	ArticleGetErr:     "获取单篇文章错误！",
	ArticleAddErr:     "添加单篇文章错误！",
	ArticleModifyErr:  "编辑单篇文章错误！",
	ArticleDelErr:     "删除单篇文章错误！",
	ArticleLikeErr:    "文章点赞失败！",
	ArticleUnLikeErr:  "文章取消点赞失败！",
	ArticleStarErr:    "文章收藏失败！",
	ArticleUnStarErr:  "文章取消收藏失败！",
	ArticleCommentErr: "获取文章评论失败！",
	CommentAddErr:     "评论添加失败",
	CommentModifyErr:  "评论修改失败！",
	CommentLikeErr:    "评论点赞失败！",
	CommentUnLikeErr:  "评论取消点赞失败！",
	FollowErr:         "关注失败！",
	UnFollowErr:       "取消关注失败！",
	ArticlesPageOk:    "文章分页查找成功！",
	ArticleOk:         "文章操作成功！",
	ArticleLikeOk:     "文章点赞成功！",
	ArticleUnLikeOk:   "文章取消点赞成功！",
	ArticleStarOk:     "文章收藏成功！",
	ArticleUnStarOk:   "文章取消收藏成功！",
	ArticleCommentOk:  "获取文章评论成功！",
	CommentAddOk:      "评论成功！",
	CommentModifyOk:   "修改评论成功！",
	CommentLikeOk:     "评论点赞成功！",
	CommentUnLikeOk:   "评论取消点赞成功！",
	FollowOk:          "关注成功！",
	UnFollowOk:        "取消关注成功！",

	// 正则相关错误信息
	FixLessZeroErr:    "验证规则错误，定长小于0",
	MaxLessZeroErr:    "验证规则错误，最大值小于0",
	MinThanMaxErr:     "验证规则错误，最大值小于最小值",
	MediumPasswordErr: "密码为%d-%d位字母、数字，字母数字必须同时存在",
	StrongPasswordErr: "密码为%d-%d位字母、数字和符号必须同时存在，符号存在开头和结尾且仅限!@#$%^*",
	ChineseNameErr:    "中文名为%d-%d位中文字符可包含'·'",
	EnglishNameErr:    "英文名为%d-%d英文字符可包含空格",
}

// 提取状态码对应的错误信息
func StatusText(code int) string {
	return statusText[code]
}
