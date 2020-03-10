package com.wcity.grpc.objects;

public class CityPhoto {
    private int id;
    private int cityId;
    private String url;
    private String timestamp;
    private boolean selected;
    private String error;

    public boolean isError(){
        return error !=  null;
    }

    public CityPhoto() {

    }

    public String getError() {
        return error;
    }

    public void setError(String error) {
        this.error = error;
    }

    public CityPhoto(int id, int cityId, String url, String timestamp, boolean selected) {
        this.id = id;
        this.cityId = cityId;
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

    public int getCityId() {
        return cityId;
    }

    public void setCityId(int cityId) {
        this.cityId = cityId;
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
