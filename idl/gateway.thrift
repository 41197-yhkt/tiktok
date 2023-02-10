// idl/gateway.thrift
namespace go douyin

// BaseResp
struct BaseResp {
    1: required i32 status_code     // 状态码，0-成功，其他值-失败
    2: optional string status_msg   // 状态描述
}

// User
struct User {
    1: required i64 id              // 用户 id
    2: required string name         // 昵称
    3: optional i64 follow_count    // 关注总数
    4: optional i64 follower_count  // 粉丝总数
    5: required bool is_follow      // true-已关注，false-未关注
}

// Vedio
struct Vedio {
    1: required i64 id              // 视频 id
    2: required User author         // 作者
    3: required string play_url     // 播放地址
    4: required string cover_url    // 封面地址
    5: required i64 favorite_count  // 点赞总数
    6: required i64 comment_count   // 评论总数
    7: required bool is_favorite    // true-已点赞，false-未点赞
    8: required string title        // 标题
}

// Comment
struct Comment {
    1: required i64 id              // 评论 id
    2: required User user           // 评论用户信息 
    3: required string content      // 评论内容
    4: required string create_date  // 评论发布日期，格式 mm-dd
}

// User Information
struct DouyinUserRequest {
    1: required i64 user_id         // 用户 id
    2: required string token        // 用户鉴权 token
}

struct DouyinUserResponse {
    1: required BaseResp base_resp   
    2: required User user           // 用户信息
}

// User Register
struct DouyinUserRegisterRequest {
    1: required string username     // 用户名，最长32个字符
    2: required string password     // 密码，最长32个字符
}

struct DouyinUserRegisterResponse {
    1: required BaseResp base_resp  
    2: required i64 user_id         // 用户 id
    3: required string token        // 用户鉴权 token
}

// User Login
struct DouyinUserLoginRequest {
    1: required string username     // 登录用户名
    2: required string password     // 登录密码
}

struct DouyinUserLoginResponse {
    1: required BaseResp base_resp
    2: required i64 user_id         // 用户 id
    3: required string token        // 用户鉴权 token
}

// Feed
struct DouyinFeedRequest {
    1: optional i64 lastest_time    // 限制返回视频最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token        // 登录用户设置
}

struct DouyinFeedResponse {
    1: required BaseResp base_resp
    3: list<Vedio> vedio_list       // 视频列表
    4: optional i64 next_time       // 发布最早时间，作为下次请求的lastest_time
}

// Publish Action
struct DouyinPublishActionRequest {
    1: required string token        // 用户鉴权 token
    2: required binary data         // 视频数据
    3: required string title        // 视频标题
}

struct DouyinPublishActionResponse {
    1: required BaseResp base_resp
}

// Publish List
struct DouyinPublishListRequest {
    1: required i64 user_id         // 用户 id
    2: required string token        // 用户鉴权 token
}

struct DouyinPublishListResponse {
    1: required BaseResp base_resp 
    2: list<Vedio> vedio_list       // 用户发布的视频列表
}

// Favorite Action
struct DouyinFavoriteActionRequest {
    1: required string token        // 用户鉴权 token
    2: required i64 vedio_id        // 视频 id
    3: required i32 action_type     // 1-点赞，2-取消
}

struct DouyinFavoriteActionResponse {
    1: required BaseResp base_resp 
}

// Favorite List
struct DouyinFavoriteListRequest {
    1: required i64 user_id         // 用户 id
    2: required string token        // 用户鉴权 token
}

struct DouyinFavoriteListResponse {
    1: required BaseResp base_resp
    2: list<Vedio> vedio_list       // 用户点赞视频列表
}

// Comment Action
struct DouyinCommentActionRequest {
    1: required string token        // 用户鉴权 token
    2: required i64 vedio_id        // 视频 id
    3: required i32 action_type     // 1-发表评论，2-删除评论
    4: optional string comment_text // action_type = 1 时，用户填写评论内容
    5: optional i64 comment_id      // action_type = 2 时，要删除的评论id
}

struct DouyinCommentActionResponse {
    1: required BaseResp base_resp   
    2: optional Comment comment     // 评论成功返回评论内容
}

// Comment List
struct DouyinCommentListRequest {
    1: required string token        // 用户鉴权 token
    2: required i64 vedio_id        // 视频 id
}

struct DouyinCommentListResponse {
    1: required BaseResp base_resp       
    2: list<Comment> comment_list   // 评论列表
}

// Relation Action
struct DouyinRelationActionRequest {
    1: required string token        // 用户鉴权 token
    2: required i64 to_user_id      // 对方用户 id
    3: required i32 action_type     // 1-关注，2-取消关注
}

struct DouyinRelationActionResponse {
    1: required BaseResp base_resp
}

// Relation Follow List
struct DouyinRelationFollowListRequest {
    1: required i64 user_id         // 用户 id
    2: required string token        // 用户鉴权 token
}

struct DouyinRelationFollowListResponse {
    1: required BaseResp base_resp
    2: list<User> user_list         // 用户列表
}

// Relation Follower List
struct DouyinRelationFollowerListRequest {
    1: required i64 user_id         // 用户 id
    2: required string token        // 用户鉴权 token
}

struct DouyinRelationFollowerListResponse {
    1: required BaseResp base_resp           
    2: list<User> user_list         // 用户列表
}

// Relation Friend List 
struct DouyinRelationFriendListRequest {
    1: required i64 user_id         // 用户 id
    2: required string token        // 用户鉴权 token
}

struct DouyinRelationFriendListResponse {
    1: required BaseResp base_resp
    2: list<User> user_list         // 用户列表
}

service DouyinService {
    // 基础接口
    DouyinUserRegisterResponse DouyinUserRegisterMethod(1: DouyinUserRegisterRequest req) (api.post="/douyin/user/register");
    DouyinUserLoginResponse DouyinUserLoginMethod(1: DouyinUserLoginRequest req) (api.post="/douyin/user/login");
    DouyinUserResponse DouyinUserMethod(1: DouyinUserRequest req) (api.get="/douyin/user");
    DouyinFeedResponse DouyinFeedMethod(1: DouyinFeedRequest req) (api.get="/douyin/feed");
    DouyinPublishActionResponse DouyinPublishActionMethod(1: DouyinPublishActionRequest req) (api.post="/douyin/publish/action")
    DouyinPublishListResponse DouyinPublishListMethod(1: DouyinPublishListRequest req) (api.get="/douyin/publish/list")

    // 互动接口
    DouyinFavoriteActionResponse DouyinFavoriteActionMethod(1: DouyinFavoriteActionRequest req) (api.post="/douyin/favorite/action")
    DouyinFavoriteListResponse DouyinFavoriteListMethod(1: DouyinFavoriteListRequest req) (api.get="/douyin/favorite/list")
    DouyinCommentActionResponse DouyinCommentActionMethod(1: DouyinCommentActionRequest req) (api.post="/douyin/comment/action")
    DouyinCommentListResponse DouyinCommentListMethod(1: DouyinCommentListRequest req) (api.get="/douyin/comment/list")

    // 社交接口
    DouyinRelationActionResponse DouyinRelationActionMethod(1: DouyinRelationActionRequest req) (api.post="/douyin/relation/action")
    DouyinRelationFollowListResponse DouyinRelationFollowListMethod(1: DouyinRelationFollowListRequest req) (api.get="/douyin/relation/follow/list")
    DouyinRelationFollowerListResponse DouyinRelationFollowerListMethod(1: DouyinRelationFollowerListRequest req) (api.get="/douyin/relation/follower/list")
    DouyinRelationFriendListResponse DouyinRelationFriendListMethod(1: DouyinRelationFriendListRequest req) (api.get="/douyin/relation/friend/list")
}
