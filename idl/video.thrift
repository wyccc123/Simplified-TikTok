namespace go videodemo

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct Video {
    1:i64 video_id
    2:i64 user_id
    3:string video_url
    4:string video_cover_url
    5:i64 favorite_count
    6:i64 comment_count
    7:string video_title
    8:bool is_favorite
    9:i64 create_time
}

struct PublishRequest {
    1:i64 user_id
    2:string video_url
    3:string video_title
}

struct PublishResponse {
    1:BaseResp base_resp
}

struct DeleteVideoRequest {
    1:i64 video_id
    2:i64 user_id
}

struct DeleteVideoResponse {
    1:BaseResp base_resp
}

struct GetVideoListRequst {
    1:i64 user_id
    2:i64 latest_time   //视频列表里最新的视频
}

struct GetVideoListResponse {
    1:BaseResp base_resp
    2:list<Video> video_list
    3:i64 next_time     //本次返回的视频列表里最早的视频，作为下次请求的latest_time
}

struct GetUserPublishedRequest {
    1:i64 user_id
}

struct GetUserPublishedResponse {
    1:BaseResp base_resp
    2:list<Video> video_list
}

struct FavoriteVideoRequest {
    1:i64 user_id
    2:i64 video_id
    3:bool is_favorite  //是否已经点赞
}

struct FavoriteVideoResponse {
    1:BaseResp base_resp
}

struct GetUserFavoriteRequest {
    1:i64 user_id
}

struct GetUserFavoriteResponse {
    1:BaseResp base_resp
    2:list<Video> video_list
}

service VideoService {
    PublishResponse PublishVideo(1:PublishRequest req)
    DeleteVideoResponse DeleteVideo(1:DeleteVideoRequest req)
    GetVideoListResponse GetVideoList(1:GetVideoListRequst req)
    GetUserPublishedResponse GetUserPublished(1:GetUserPublishedRequest req)
    FavoriteVideoResponse FavoriteVideo(1:FavoriteVideoRequest req)
    GetUserFavoriteResponse GetUserFavorite(1:GetUserFavoriteRequest req)
}