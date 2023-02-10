// idl/user.thrift
namespace go user

// User
struct User {
    1: required i64 id              // 用户 id
    2: required string name         // 昵称
    3: optional i64 follow_count    // 关注总数
    4: optional i64 follower_count  // 粉丝总数
    5: required bool is_follow      // true-已关注，false-未关注
}

// BaseResp
struct BaseResp {
    1: required i32 status_code     // 状态码，0-成功，其他值-失败
    2: optional string status_msg   // 状态描述
}

// User Register Request
struct UserRegisterRequest {
    1: required string username     // 用户名，最长32个字符
    2: required string password     // 密码，最长32个字符
}

// User Register Response
struct UserRegisterResponse {
    1: required BaseResp base_resp
    2: required i64 user_id         // 用户 id
    3: required string token        // 用户鉴权 token
}

// User Login Request
struct UserLoginRequest {
    1: required string username     // 登录用户名
    2: required string password     // 登录密码
}

// User Login Response
struct UserLoginResponse {
    1: required BaseResp base_resp
    2: required i64 user_id         // 用户 id
    3: required string token        // 用户鉴权 token
}

// User Info Request
struct UserInfoRequest {
    1: required i64 user_id    (go.tag = 'query:"user_id"')     // 用户 id
    2: required string token   (go.tag = 'query:"token"')     // 用户鉴权 token
}

// User Info Response
struct UserInfoResponse {
    1: required BaseResp base_resp
    2: required User user           // 用户信息
}

// User Follow Request, Update User Relation
struct UserFollowRequest {
    1: required i64 follow_from;
    2: required i64 follow_to;
}

// User Follow Response, Update User Relation
struct UserFollowResponse {
    1: required BaseResp base_resp;
}

// User Follow Request, Update User Relation
struct UserUnfollowRequest {
    1: required i64 follow_from;
    2: required i64 follow_to;
}

// User Follow Response, Update User Relation
struct UserUnfollowResponse {
    1: required BaseResp base_resp;
}

struct FollowListRequest {
    1: required i64 user_id
    2: required string token        // 用户鉴权 token
}

struct FollowListResponse {
    1: required BaseResp base_resp
    2: required list<User> user_list
}

struct FollowerListRequest {
    1: required i64 user_id
    2: required string token        // 用户鉴权 token
}

struct FollowerListResponse {
    1: required BaseResp base_resp
    2: required list<User> user_list
}

struct FriendListRequest {
    1: required i64 user_id
    2: required string token        // 用户鉴权 token
}

struct FriendListResponse {
    1: required BaseResp base_resp
    2: required list<User> user_list
}

struct IsFriendRequest{
    1:required  i64 user_id
    2:required  string token
    3:required  i64  to_user_id
}

struct IsFriendResponse{
    1:BaseResp base_resp
    2:required  bool is_friend
}

// GetUser
struct CompGetUserRequest {
    1: required i64 user_id         // 用户 id
    2: required i64 target_user_id  // 目标用户 id
}

struct CompGetUserResponse {
    1: required BaseResp base_resp
    2: required User user           // 用户信息
}

// MGetUser
struct CompMGetUserRequest {
    1: required i64 user_id         // 用户 id
    2: list<i64> target_users_id    // 目标用户 id 列表
}

struct CompMGetUserResponse {
    1: required BaseResp base_resp
    2: list<User> user_list         // 用户信息列表
}

service UserService {
    // 用户注册
    UserRegisterResponse UserRegister(1: UserRegisterRequest req)
    // 用户登陆
    UserLoginResponse UserLogin(1: UserLoginRequest req)
    // 获取用户信息
    UserInfoResponse UserInfo(1: UserInfoRequest req)
    // 用户关注，用户点击关注时，维护用户的信息
    UserFollowResponse UserFollow(1: UserFollowRequest req)
    // 用户取消关注，用户点击取消关注时，维护用户的信息
    UserUnfollowResponse UserUnfollow(1: UserUnfollowRequest req)
    // 查询用户关注列表
    FollowListResponse GetFollowList(1:FollowListRequest req)
    // 查询用户粉丝列表
    FollowerListResponse GetFollowerList(1:FollowerListRequest req)
    // 查询用户好友列表
    FriendListResponse GetFriendList(1:FriendListRequest req)
    // 判断是否是好友
    IsFriendResponse IsFriend(1:IsFriendRequest req)
    // 获取target_user和当前user的关系，并封装target_user返回
    CompGetUserResponse CompGetUser(1:CompGetUserRequest req)
    // 获取target_users和当前user的关系，并封装target_users返回
    CompMGetUserResponse CompMGetUser(1:CompMGetUserRequest req)
}