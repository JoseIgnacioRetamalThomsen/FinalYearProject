package com.wcity.grpc;

public class VisitCity {
    private boolean valid;
    private String email;
    private String cityName;
    private String cityCountry;

    public VisitCity(boolean valid, String email, String cityName, String cityCountry) {
        this.valid = valid;
        this.email = email;
        this.cityName = cityName;
        this.cityCountry = cityCountry;
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

    public String getCityName() {
        return cityName;
    }

    public void setCityName(String cityName) {
        this.cityName = cityName;
    }

    public String getCityCountry() {
        return cityCountry;
    }

    public void setCityCountry(String cityCountry) {
        this.cityCountry = cityCountry;
    }
}
