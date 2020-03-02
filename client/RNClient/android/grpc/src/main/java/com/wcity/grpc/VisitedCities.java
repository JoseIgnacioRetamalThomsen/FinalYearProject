package com.wcity.grpc;

import java.util.ArrayList;

public class VisitedCities {
    ArrayList<CreateCity> visitedCities;
    private boolean valid;
    private String email;


    public VisitedCities(boolean valid, String email, ArrayList<CreateCity> visitedCities) {
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

    public ArrayList<CreateCity> getVisitedCities() {
        return visitedCities;
    }

    public void setVisitedCities(ArrayList<CreateCity> visitedCities) {
        this.visitedCities = visitedCities;
    }
}
