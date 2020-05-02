package com.wcity.grpc;

public class VisitPlace {
    private boolean valid;
    private String email;
    private String placeName;
    private String placeCity;
    private String placeCountry;

    public VisitPlace(boolean valid, String email, String placeName, String placeCity,
                      String placeCountry) {
        this.valid = valid;
        this.email = email;
        this.placeName = placeName;
        this.placeCity = placeCity;
        this.placeCountry = placeCountry;
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

    public String getPlaceName() {
        return placeName;
    }

    public void setPlaceName(String placeName) {
        this.placeName = placeName;
    }

    public String getPlaceCity() {
        return placeCity;
    }

    public void setPlaceCity(String placeCity) {
        this.placeCity = placeCity;
    }

    public String getPlaceCountry() {
        return placeCountry;
    }

    public void setPlaceCountry(String placeCountry) {
        this.placeCountry = placeCountry;
    }
}
