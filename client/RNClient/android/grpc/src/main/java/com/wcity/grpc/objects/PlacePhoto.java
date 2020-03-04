package com.wcity.grpc.objects;

public class PlacePhoto {
    private int id;
    private int placeId;
    private String url;
    private String timestamp;
    private boolean selected;

    public PlacePhoto(int id, int placeId, String url, String timestamp, boolean selected) {
        this.id = id;
        this.placeId = placeId;
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

    public int getPlaceId() {
        return placeId;
    }

    public void setPlaceId(int placeId) {
        this.placeId = placeId;
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
