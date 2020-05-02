package com.wcity.grpc.objects;

public class City {
    private String name;
    private String country;
    private String creatorEmail;
    private float lat;
    private float lon;
    private String description;
    private int cityId;
    public String error;

    public City(String name, String country, String creatorEmail, float lat, float lon,
                String description, int cityId) {
        this.name = name;
        this.country = country;
        this.creatorEmail = creatorEmail;
        this.lat = lat;
        this.lon = lon;
        this.description = description;
        this.cityId = cityId;
    }

    public City() {

    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getCountry() {
        return country;
    }

    public void setCountry(String country) {
        this.country = country;
    }

    public String getCreatorEmail() {
        return creatorEmail;
    }

    public void setCreatorEmail(String creatorEmail) {
        this.creatorEmail = creatorEmail;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public int getId() {
        return cityId;
    }

    public void setId(int cityId) {
        this.cityId = cityId;
    }

    public float getLat() {
        return lat;
    }

    public void setLat(float lat) {
        this.lat = lat;
    }

    public float getLon() {
        return lon;
    }

    public void setLon(float lon) {
        this.lon = lon;
    }

    public int getCityId() {
        return cityId;
    }

    public void setCityId(int cityId) {
        this.cityId = cityId;
    }
}
