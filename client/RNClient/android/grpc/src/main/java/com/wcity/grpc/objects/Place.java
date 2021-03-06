package com.wcity.grpc.objects;

public class Place {
    private String name;
    private String city;
    private String country;
    private String creatorEmail;
    private String description;
    private float lat;
    private float lon;
    private int id;
    public String error;

    public Place() {
    }

    public Place(String name, String city, String country, String creatorEmail,
                 String description, float lat, float lon, int id) {

        this.name = name;
        this.city = city;
        this.country = country;
        this.creatorEmail = creatorEmail;
        this.description = description;
        this.lat = lat;
        this.lon = lon;
        this.id = id;
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

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }
}
