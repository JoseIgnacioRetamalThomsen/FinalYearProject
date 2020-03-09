package com.wcity.grpc.objects;

public class User {
    private String email;
    private String name;
    private String description;
    private int userId;

    public User(String email, String name, String description, int userId) {
        this.email = email;
        this.name = name;
        this.description = description;
        this.userId = userId;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public int getUserId() {
        return userId;
    }

    public void setUserId(int userId) {
        this.userId = userId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }
}
