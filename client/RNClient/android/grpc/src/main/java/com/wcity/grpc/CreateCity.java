package com.wcity.grpc;

public class CreateCity {
    private boolean valid;
    private String name;
    private String country;
    private String creatorEmail;
    private String description;
    private float lat;
    private float lon;
    private int id;


    public CreateCity(boolean valid, String name, String country, String creatorEmail, String description, float lat, float lon, int id) {
        this.valid = valid;
        this.name = name;
        this.country = country;
        this.creatorEmail = creatorEmail;
        this.description = description;
        this.lat = lat;
        this.lon = lon;
        this.id = id;
    }

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
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
        return id;
    }

    public void setId(int id) {
        this.id = id;
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
}
