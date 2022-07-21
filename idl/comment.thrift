namespace go commentdemo

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct Comment {
    1:i64 comment_id
    2:i64 user_id
    3:string content
    4:string create_date
}

struct GiveCommentRequest {
    1:i64 user_id
    2:i64 video_id
    3:string content
}

struct GiveCommentResponse {
    1:BaseResp base_resp
}

struct DeleteCommentRequest {
    1:i64 user_id
    2:i64 comment_id
}

struct DeleteCommentResponse {
    1:BaseResp base_resp
}

struct GetVideoCommentRequest {
    1:i64 user_id
    2:i64 video_id
}

struct GetVideoCommentResponse {
    1:BaseResp base_resp
    2:list<Comment> comment_list
}

service CommentService {
    GiveCommentResponse GiveComment(1:GiveCommentRequest req)
    DeleteCommentResponse DeleteComment(1:DeleteCommentRequest req)
    GetVideoCommentResponse GetVideoComment(1:GetVideoCommentRequest req)
}
