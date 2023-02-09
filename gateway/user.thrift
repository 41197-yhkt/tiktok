// idl/user.thrift
namespace go user

// BaseResp
struct BaseResp {
    1: required i32 status_code     // 状态码，0-成功，其他值-失败
    2: optional string status_msg   // 状态描述
}

// User
struct User {
    1: required i64 id              // 用户 id
    2: required string name         // 昵称
    3: optional string username     // 用户名，设为 optional 的原因 DouyinUserResponse 不传
    4: optional string password     // 密码，设为 optional 的原因是 DouyinUserResponse 不传
    5: list<i64> follow_user_id_list    // 关注用户 id 列表， 关注总数 follow_count 由 follow_user_id_list 长度得出
    6: list<i64> follower_user_id_list  // 粉丝用户 id 列表， 粉丝总数 follower_count 由 follower_user_id_list 长度得出
    7: optional bool is_follow      // true-已关注，false-未关注
}

// User Information
struct DouyinUserRequest {
    1: required i64 user_id         // 用户 id
    // 2: required string token     // 用户鉴权 token
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
    // 发 token 是不是由 Gateway 来完成？
    // 3: required string token     // 用户鉴权 token
}

// User Login
struct DouyinUserLoginRequest {
    1: required string username     // 登录用户名
    2: required string password     // 登录密码
}

struct DouyinUserLoginResponse {
    1: required BaseResp base_resp
    2: required i64 user_id         // 用户 id
    // 3: required string token     // 用户鉴权 token
}

service DouyinService {
    DouyinUserRegisterResponse DouyinUserRegisterMethod(1: DouyinUserRegisterRequest req) (api.post="/douyin/user/register");
    DouyinUserLoginResponse DouyinUserLoginMethod(1: DouyinUserLoginRequest req) (api.post="/douyin/user/login");
    DouyinUserResponse DouyinUserMethod(1: DouyinUserRequest req) (api.get="/douyin/user");
}
