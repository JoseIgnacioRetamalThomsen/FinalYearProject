package com.wcity.grpc.objects;

import java.util.List;

public class CityPhotoResponse {
    private  boolean valid;
    private int cityID;
    private List <CityPhoto> photos;
    private int active;

    public CityPhotoResponse(boolean valid, int cityID, List<CityPhoto> photos, int active) {
        this.valid = valid;
        this.cityID = cityID;
        this.photos = photos;
        this.active = active;
    }

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
    }

    public int getCityID() {
        return cityID;
    }

    public void setCityID(int cityID) {
        this.cityID = cityID;
    }

    public List<CityPhoto> getPhotos() {
        return photos;
    }

    public void setPhotos(List<CityPhoto> photos) {
        this.photos = photos;
    }

    public int getActive() {
        return active;
    }

    public void setActive(int active) {
        this.active = active;
    }
}
