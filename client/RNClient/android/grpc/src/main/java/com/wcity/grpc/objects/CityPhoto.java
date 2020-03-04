package com.wcity.grpc.objects;

public class CityPhoto {
    private int id;
    private String cityId;
    private String url;
    private String timestamp;
    private boolean selected;

    public CityPhoto(int id, String cityId, String url, String timestamp, boolean selected) {
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

    public String getCityId() {
        return cityId;
    }

    public void setCityId(String cityId) {
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
