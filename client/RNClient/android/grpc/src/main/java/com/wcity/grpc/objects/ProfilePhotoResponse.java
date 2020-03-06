package com.wcity.grpc.objects;

import java.util.List;

public class ProfilePhotoResponse {
    private String email;
    private boolean valid;
    private List<ProfilePhoto> photos;
    private ProfilePhoto photo;

    public ProfilePhotoResponse(String email, boolean valid, ProfilePhoto photo) {
        this.email = email;
        this.valid = valid;
        this.photo = photo;
    }

    public ProfilePhotoResponse(String email, boolean valid, List<ProfilePhoto> photos) {
        this.email = email;
        this.valid = valid;
        this.photos = photos;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
    }

    public List<ProfilePhoto> getPhotos() {
        return photos;
    }

    public void setPhotos(List<ProfilePhoto> photos) {
        this.photos = photos;
    }

    public ProfilePhoto getPhoto() {
        return photo;
    }

    public void setPhoto(ProfilePhoto photo) {
        this.photo = photo;
    }

}
