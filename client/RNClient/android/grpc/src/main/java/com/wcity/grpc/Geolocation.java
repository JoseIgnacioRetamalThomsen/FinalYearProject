package com.wcity.grpc;

public class Geolocation {
    private float longitude;
    private float latitude;

    public Geolocation(float longitude, float latitude) {
        this.longitude = longitude;
        this.latitude = latitude;
    }

    public float getLongitude() {
        return longitude;
    }

    public void setLongitude(float longitude) {
        this.longitude = longitude;
    }

    public float getLatitude() {
        return latitude;
    }

    public void setLatitude(float latitude) {
        this.latitude = latitude;
    }
    public void setLocation(float latitude, float longitude) {
        this.latitude = latitude;
        this.longitude = longitude;
    }
}
