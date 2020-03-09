package com.wcity.grpc.objects;

public class CreateCityRequest {
    private String token;
    private String email;
    private City city;

    public CreateCityRequest(String token, String email, City city) {
        this.token = token;
        this.email = email;
        this.city = city;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
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
