package com.wcity.grpc;

import io.grpc.wcity.profiles.GeolocationP;

public class Place {
    private boolean valid;
    private String name;
    private String city;
    private String country;
    private String creatorEmail;
    private String description;
    private GeolocationP location;

    public Place(boolean valid, String name, String city, String country, String creatorEmail, String description, GeolocationP location) {
        this.valid = valid;
        this.name = name;
        this.city = city;
        this.country = country;
        this.creatorEmail = creatorEmail;
        this.description = description;
        this.location = location;
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

    public String getCity() {
        return city;
    }

    public void setCity(String city) {
        this.city = city;
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

    public GeolocationP getLocation() {
        return location;
    }

    public void setLocation(GeolocationP location) {
        this.location = location;
    }
}
