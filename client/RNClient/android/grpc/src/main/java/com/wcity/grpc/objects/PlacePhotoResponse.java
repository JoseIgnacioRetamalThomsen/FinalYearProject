package com.wcity.grpc.objects;

import java.util.List;

public class PlacePhotoResponse {
    private boolean valid;
    private int placeId;
    private List<PlacePhoto> photos;
    private boolean active;

    public PlacePhotoResponse(boolean valid, int placeId, List<PlacePhoto> photos, boolean active) {
        this.valid = valid;
        this.placeId = placeId;
        this.photos = photos;
        this.active = active;
    }

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
    }

    public int getPlaceId() {
        return placeId;
    }

    public void setPlaceId(int placeId) {
        this.placeId = placeId;
    }

    public List<PlacePhoto> getPhotos() {
        return photos;
    }

    public void setPhotos(List<PlacePhoto> photos) {
        this.photos = photos;
    }

    public boolean isActive() {
        return active;
    }

    public void setActive(boolean active) {
        this.active = active;
    }
}
