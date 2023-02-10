// idl/video.thrift
namespace go video

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

// video
struct video {
    1: required i64 id                  // 视频 id
    2: required User author             // 作者
    3: required string play_url         // 播放地址
    4: required string cover_url        // 封面地址
    5: required i64 favorite_count      // 点赞总数
    6: required i64 comment_count       // 评论总数
    7: required bool is_favorite        // true-已点赞，false-未点赞
    8: required string title            // 标题
}

// Publish Action
struct DouyinPublishActionRequest {
    1: required binary data         // 视频数据
    2: required string title        // 视频标题
    3: required i64  author         // 视频作者
    // 3: required string token     // 用户鉴权 token
}

struct DouyinPublishActionResponse {
    1: required BaseResp base_resp
}

// Publish List
struct DouyinPublishListRequest {
    1: required i64 user_id         // 用户 id
    // 2: required string token     // 用户鉴权 token
}

struct DouyinPublishListResponse {
    1: required BaseResp base_resp 
    2: list<video> video_list       // 用户发布的视频列表
}

// GetVideo
struct GetVideoRequest {
    1: required i64 user_id          // 查询用户 id，用于判断 is_favorite
    2: required i64 target_video_id  // 目标视频 id
}

struct GetVideoResponse {
    1: required video video           // 视频信息
}

// MGetVideo
struct MGetVideoRequest {
    1: required i64 user_id          // 查询用户 id，用于判断 is_favorite
    2: list<i64> target_videos_id    // 目标视频 id 列表
}

struct MGetVideoResponse {
    1: list<video> video_list         // 视频信息列表
}

service DouyinService {
    DouyinPublishActionResponse DouyinPublishActionMethod(1: DouyinPublishActionRequest req) (api.post="/douyin/publish/action");
    DouyinPublishListResponse DouyinPublishListMethod(1: DouyinPublishListRequest req) (api.get="/douyin/publish/list");
    GetVideoResponse DouyinGetVideoMethod(1: GetVideoRequest req) (api.get="/douyin/publish/GetVideo");
    MGetVideoResponse DouyinMGetVideoMethod(1: MGetVideoRequest req) (api.get="/douyin/publish/MGetVideo");
}