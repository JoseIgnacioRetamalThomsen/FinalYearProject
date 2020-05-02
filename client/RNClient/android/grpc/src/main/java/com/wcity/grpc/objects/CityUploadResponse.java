package com.wcity.grpc.objects;

public class CityUploadResponse {
    private boolean success;
    private CityPhoto photo;

    public CityUploadResponse(boolean success, CityPhoto photo) {
        this.success = success;
        this.photo = photo;
    }

    public boolean isSuccess() {
        return success;
    }

    public void setSuccess(boolean success) {
        this.success = success;
    }

    public CityPhoto getPhoto() {
        return photo;
    }

    public void setPhoto(CityPhoto photo) {
        this.photo = photo;
    }
}
