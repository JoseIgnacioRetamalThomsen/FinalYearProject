package com.wcity.grpc;

public class User {
    boolean isValid;
    private String name;
    private String description;

    public User(boolean isValid, String name, String description) {
        this.isValid = isValid;
        this.name = name;
        this.description = description;
    }

    public boolean isValid() {
        return isValid;
    }

    public void setValid(boolean valid) {
        isValid = valid;
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
