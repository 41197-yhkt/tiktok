// idl/composite.thrift
namespace go composite

// BaseResp
struct BaseResp {
    1: required i32 status_code         // 状态码，0-成功，其他值-失败
    2: optional string status_msg       // 状态描述
}

// User 
struct User {
    1: required i64 id                  // 用户 id
    2: list<Video> favorite_video_list  // 用户点赞视频列表
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

// Feed
struct BasicFeedRequest {
    1: optional i64 lastest_time        // 限制返回视频最新投稿时间戳，精确到秒，不填表示当前时间
}

struct BasicFeedResponse {
    1: required BaseResp base_resp
    3: list<Video> video_list           // 视频列表
    4: optional i64 next_time           // 发布最早时间，作为下次请求的lastest_time
}

// Favarite Action
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
}

struct BasicFavoriteListResponse {
    1: required BaseResp base_resp
    2: list<Video> video_list           // 用户点赞视频列表
}


service CompositeService {
    // 互动接口
    BasicFavoriteActionResponse BasicFavoriteActionMethod(1: BasicFavoriteActionRequest req) (api.post="/douyin/favorite/action")
    BasicFavoriteListResponse BasicFavoriteListMethod(1: BasicFavoriteListRequest req) (api.get="/douyin/favorite/list")
    BasicFeedResponse BasicFeedMethod(1: BasicFeedRequest req) (api.get="/douyin/feed");
}
