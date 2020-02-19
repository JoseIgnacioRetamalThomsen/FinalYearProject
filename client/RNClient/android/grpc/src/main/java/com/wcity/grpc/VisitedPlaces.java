package com.wcity.grpc;

public class VisitedPlaces {
    private boolean valid;
    private String email;
    private Place place;

    public VisitedPlaces(boolean valid, String email, Place place) {
        this.valid = valid;
        this.email = email;
        this.place = place;
    }

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public Place getPlace() {
        return place;
    }

    public void setPlace(Place place) {
        this.place = place;
    }
}
