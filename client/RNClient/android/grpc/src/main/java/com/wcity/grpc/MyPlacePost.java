package com.wcity.grpc;

import java.util.List;

public class MyPlacePost {
    private int indexId;
    private String creatorEmail;
    private String cityName;
    private String countryName;
    private String placeName;
    private String title;
    private String body;
    private String timeStamp;
    private List<String> likes;
    private String mongoId;

    public MyPlacePost(int indexId, String creatorEmail, String cityName, String countryName,
                       String placeName, String title, String body, String timeStamp,
                       List<String> likes, String mongoId) {
        this.indexId = indexId;
        this.creatorEmail = creatorEmail;
        this.cityName = cityName;
        this.countryName = countryName;
        this.placeName = placeName;
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

    public String getCountryName() {
        return countryName;
    }

    public void setCountryName(String countryName) {
        this.countryName = countryName;
    }

    public String getPlaceName() {
        return placeName;
    }

    public void setPlaceName(String placeName) {
        this.placeName = placeName;
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
