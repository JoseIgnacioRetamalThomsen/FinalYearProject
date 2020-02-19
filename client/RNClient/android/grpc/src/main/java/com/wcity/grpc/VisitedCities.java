package com.wcity.grpc;

public class VisitedCities {
    private boolean valid;
    private String email;
    private City city;

    public VisitedCities(boolean valid, String email, City city) {
        this.valid = valid;
        this.email = email;
        this.city = city;
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

    public City getCity() {
        return city;
    }

    public void setCity(City city) {
        this.city = city;
    }
}
