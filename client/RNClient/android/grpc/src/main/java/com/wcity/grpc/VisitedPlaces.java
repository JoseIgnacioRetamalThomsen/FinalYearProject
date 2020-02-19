package com.wcity.grpc;

import java.util.ArrayList;


public class VisitedPlaces {
    ArrayList<Place> visitedPlaces;
    private boolean valid;
    private String email;

    public VisitedPlaces(boolean valid, String email, ArrayList<Place> visitedPlaces) {
        this.valid = valid;
        this.email = email;
        this.visitedPlaces = visitedPlaces;
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

    public ArrayList<Place> getVisitedPlaces() {
        return visitedPlaces;
    }

    public void setVisitedPlaces(ArrayList<Place> visitedPlaces) {
        this.visitedPlaces = visitedPlaces;
    }
}
