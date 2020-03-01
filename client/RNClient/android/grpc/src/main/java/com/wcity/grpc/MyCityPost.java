package com.wcity.grpc;

import java.util.List;

public class MyCityPost {
    private int indexId;
    private String creatorEmail;
    private String cityName;
    private String cityCountry;
    private String title;
    private String body;
    private String timeStamp;
    private List<String> likes;
    private String mongoId;

    public MyCityPost(int indexId, String creatorEmail, String cityName, String cityCountry,
                      String title, String body, String timeStamp, List<String> likes, String mongoId) {
        this.indexId = indexId;
        this.creatorEmail = creatorEmail;
        this.cityName = cityName;
        this.cityCountry = cityCountry;
        this.title = title;
        this.body = body;
        this.timeStamp = timeStamp;
        this.likes = likes;
        this.mongoId = mongoId;
    }

    public int getIndexId() {
        return indexId;
    }

    public void setIndexId(int indexId) {
        this.indexId = indexId;
    }

    public String getCreatorEmail() {
        return creatorEmail;
    }

    public void setCreatorEmail(String creatorEmail) {
        this.creatorEmail = creatorEmail;
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

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getBody() {
        return body;
    }

    public void setBody(String body) {
        this.body = body;
    }

    public String getTimeStamp() {
        return timeStamp;
    }

    public void setTimeStamp(String timeStamp) {
        this.timeStamp = timeStamp;
    }

    public List<String> getLikes() {
        return likes;
    }

    public void setLikes(List<String> likes) {
        this.likes = likes;
    }

    public String getMongoId() {
        return mongoId;
    }

    public void setMongoId(String mongoId) {
        this.mongoId = mongoId;
    }
}
