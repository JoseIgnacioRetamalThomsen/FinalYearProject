package com.wcity.grpc.objects;

public class PostUploadResponse {

    private boolean success;
    private PostPhoto photo;

    public PostUploadResponse(boolean success, PostPhoto photo) {
        this.success = success;
        this.photo = photo;
    }

    public boolean isSuccess() {
        return success;
    }

    public void setSuccess(boolean success) {
        this.success = success;
    }

    public PostPhoto getPhoto() {
        return photo;
    }

    public void setPhoto(PostPhoto photo) {
        this.photo = photo;
    }
}
