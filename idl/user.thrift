namespace go userdemo

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct User {
    1:i64 user_id
    2:string user_name
    3:i64 follow_count
    4:i64 fans_count
    5:bool is_follow
    6:i64 create_time
}

struct RegisterRequest {
    1:string user_name
    2:string password
}

struct RegisterResponse {
    1:BaseResp base_resp
    2:i64 user_id
}

struct GetUserRequest {
    1:i64 user_id
}

struct GetUserResponse {
    1:BaseResp base_resp
    2:User user
}

struct FollowRequest {
    1:i64 user_id
    2:i64 follow_user_id
    3:bool is_follow
}

struct FollowResponse {
    1:BaseResp base_resp
}

struct GetFollowListRequest {
    1:i64 user_id
}

struct GetFollowListResponse {
    1:BaseResp base_resp
    2:list<User> user_list
}

service UserService {
    RegisterResponse Register(1:RegisterRequest req)
    RegisterResponse Login(1:RegisterRequest req)
    GetUserResponse GetUserById(1:GetUserRequest req)
    FollowResponse Follow(1:FollowRequest req)
    GetFollowListResponse GetFollowList(1:GetFollowListRequest req)
}

