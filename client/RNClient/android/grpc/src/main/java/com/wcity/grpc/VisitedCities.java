package com.wcity.grpc;

import java.util.ArrayList;

public class VisitedCities {
    ArrayList<City> visitedCities;
    private boolean valid;
    private String email;


    public VisitedCities(boolean valid, String email, ArrayList<City> visitedCities) {
        this.valid = valid;
        this.email = email;
        this.visitedCities = visitedCities;
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

    public ArrayList<City> getVisitedCities() {
        return visitedCities;
    }

    public void setVisitedCities(ArrayList<City> visitedCities) {
        this.visitedCities = visitedCities;
    }
}
