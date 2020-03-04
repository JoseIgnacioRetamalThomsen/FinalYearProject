package com.wcity.grpc.objects;

public class PostPhoto {
    private int id;
    private String postId;
    private String url;
    private String timestamp;
    private boolean selected;

    public PostPhoto(int id, String postId, String url, String timestamp, boolean selected) {
        this.id = id;
        this.postId = postId;
        this.url = url;
        this.timestamp = timestamp;
        this.selected = selected;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getPostId() {
        return postId;
    }

    public void setPostId(String postId) {
        this.postId = postId;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public String getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(String timestamp) {
        this.timestamp = timestamp;
    }

    public boolean isSelected() {
        return selected;
    }

    public void setSelected(boolean selected) {
        this.selected = selected;
    }
}
