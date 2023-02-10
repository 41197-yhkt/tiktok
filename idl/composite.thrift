// idl/composite.thrift
namespace go composite

// BaseResp
struct BaseResp {
    1: required i32 status_code         // 状态码，0-成功，其他值-失败
    2: optional string status_msg       // 状态描述
}

// User
struct User {
    1: required i64 id              // 用户 id
    2: required string name         // 昵称
    3: optional i64 follow_count    // 关注总数
    4: optional i64 follower_count  // 粉丝总数
    5: required bool is_follow      // true-已关注，false-未关注
}

// Video
struct Video {
    1: required i64 id                  // 视频 id
    2: required User author             // 作者
    3: required string play_url         // 播放地址
    4: required string cover_url        // 封面地址
    5: required i64 favorite_count      // 点赞总数
    6: required i64 comment_count       // 评论总数
    7: required bool is_favorite        // true-已点赞，false-未点赞
    8: required string title            // 标题
}

// Comment
struct Comment {
    1: required i64 id              // 评论 id
    2: required User user           // 评论用户信息 
    3: required string content      // 评论内容
    4: required string create_date  // 评论发布日期，格式 mm-dd
}

// Feed
struct BasicFeedRequest {
    1: required i64 user_id             // 用户 id
    2: optional i64 lastest_time        // 限制返回视频最新投稿时间戳，精确到秒，不填表示当前时间
}

struct BasicFeedResponse {
    1: required BaseResp base_resp
    3: list<Video> video_list           // 视频列表
    4: optional i64 next_time           // 发布最早时间，作为下次请求的lastest_time
}

// Favorite Action
struct BasicFavoriteActionRequest {
    1: required i64 video_id            // 视频 id
    2: required i64 user_id             // 点赞用户
    3: required i32 action_type         // 1-点赞，2-取消
}

struct BasicFavoriteActionResponse {
    1: required BaseResp base_resp 
}

// Favorite List
struct BasicFavoriteListRequest {
    1: required i64 user_id             // 用户 id
    2: required i64 query_id            // 查询用户 id
}

struct BasicFavoriteListResponse {
    1: required BaseResp base_resp
    2: list<Video> video_list           // 用户点赞视频列表
}

// Comment Action
struct BasicCommentActionRequest {
    1: required i64 user_id         // 用户 id
    2: required i64 video_id        // 视频 id
    3: required i32 action_type     // 1-发表评论，2-删除评论
    4: optional string comment_text // action_type = 1 时，用户填写评论内容
    5: optional i64 comment_id      // action_type = 2 时，要删除的评论id
}

struct BasicCommentActionResponse {
    1: required BaseResp base_resp   
    2: optional Comment comment     // 评论成功返回评论内容
}

// Comment List
struct BasicCommentListRequest {
    1: required i64 user_id
    2: required i64 video_id        // 视频 id
}

struct BasicCommentListResponse {
    1: required BaseResp base_resp       
    2: list<Comment> comment_list   // 评论列表
}


service CompositeService {
    // Favorite 
    BasicFavoriteActionResponse BasicFavoriteActionMethod(1: BasicFavoriteActionRequest req) (api.post="/douyin/favorite/action")
    BasicFavoriteListResponse BasicFavoriteListMethod(1: BasicFavoriteListRequest req) (api.get="/douyin/favorite/list")
    
    // Feed
    BasicFeedResponse BasicFeedMethod(1: BasicFeedRequest req) (api.get="/douyin/feed");

    // Comment
    BasicCommentActionResponse BasicCommentActionMethod(1: BasicCommentActionRequest req) (api.post="/douyin/comment/action")
    BasicCommentListResponse BasicCommentListMethod(1: BasicCommentListRequest req) (api.get="/douyin/comment/list")
}