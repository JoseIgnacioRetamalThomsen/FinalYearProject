package com.wcity.grpc.objects;

import java.util.List;

public class PostPhotoResponse {
    private boolean valid;
    private String postId;
    private String userEmail;
    private List<PostPhoto> photos;

    public PostPhotoResponse(boolean valid, String postId, String userEmail,
                             List<PostPhoto> photos) {
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

    public List<PostPhoto> getPhotos() {
        return photos;
    }

    public void setPhotos(List<PostPhoto> photos) {
        this.photos = photos;
    }
}
