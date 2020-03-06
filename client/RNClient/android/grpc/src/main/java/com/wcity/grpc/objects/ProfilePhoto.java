package com.wcity.grpc.objects;

public class ProfilePhoto {
    private int id;
    private String userEmail;
    private String url;
    private String timestamp;
    private boolean selected;
    public String error;

    public ProfilePhoto() {

    }

    public ProfilePhoto(int id, String userEmail, String url, String timestamp, boolean selected) {
        this.id = id;
        this.userEmail = userEmail;
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

    public String getUserEmail() {
        return userEmail;
    }

    public void setUserEmail(String userEmail) {
        this.userEmail = userEmail;
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
