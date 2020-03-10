package com.wcity.grpc.objects;

import java.util.List;

public class PostPhotoResponse {
    private boolean valid;
    private String postId;
    private String userEmail;
    private String photos;

    public PostPhotoResponse(boolean valid, String postId, String userEmail,
                             String photos) {
        this.valid = valid;
        this.postId = postId;
        this.userEmail = userEmail;
        this.photos = photos;
    }

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
    }

    public String getPostId() {
        return postId;
    }

    public void setPostId(String postId) {
        this.postId = postId;
    }

    public String getUserEmail() {
        return userEmail;
    }

    public void setUserEmail(String userEmail) {
        this.userEmail = userEmail;
    }

    public String getPhotos() {
        return photos;
    }

    public void setPhotos(String photos) {
        this.photos = photos;
    }
}
