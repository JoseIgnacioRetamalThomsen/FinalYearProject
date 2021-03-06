package com.rnclient.grpc;

import com.facebook.react.bridge.Callback;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.ReadableArray;
import com.google.gson.Gson;
import com.wcity.grpc.CityPostResponse;
import com.wcity.grpc.PlacePostResponse;
import com.wcity.grpc.clients.PostClient;


public class PostModule extends ReactContextBaseJavaModule {

    private static ReactApplicationContext reactContext;

    private static final String DURATION_SHORT_KEY = "SHORT";
    private static final String DURATION_LONG_KEY = "LONG";
    private static final String IP_ADDRESS = "35.197.242.214";
    private static final int PORT_NUMBER = 10051;
    private static PostClient client;

    public PostModule(ReactApplicationContext context) {
        super(context);
        reactContext = context;
        client = new PostClient(IP_ADDRESS, PORT_NUMBER);
    }

    @Override
    public String getName() {
        return "PostModule";
    }

    @ReactMethod
    public void createCityPost(int indexId, String creatorEmail, String cityName, String cityCountry,
                               String title, String body, Callback errorCallback,
                               Callback successCallback) {
        String response;
        try {
            response = client.createCityPost(indexId, creatorEmail, cityName, cityCountry, title, body);
            successCallback.invoke(response);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void createPlacePost(int indexId, String creatorEmail, String cityName, String countryName,
                                String placeName, String title, String body, Callback errorCallback,
                                Callback successCallback) {

        String response;
        try {
            response = client.createPlacePost(indexId, creatorEmail, cityName, countryName,
                    placeName, title, body);
            successCallback.invoke(response);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getPlacePosts(int indexId, Callback errorCallback, Callback successCallback) {
        Gson gson = new Gson();
        PlacePostResponse response = client.getPlacePosts(indexId);
        try {
            successCallback.invoke(gson.toJson(response.getMyPlacePosts()));
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getCityPosts(int indexId, Callback errorCallback, Callback successCallback) {

        Gson gson = new Gson();
        CityPostResponse response = client.getCityPosts(indexId);
        try {
            successCallback.invoke(gson.toJson(response.getMyCityPosts()));
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void updateCityPost(String mongoId, String title, String body, Callback errorCallback,
                               Callback successCallback) {
        boolean isValid;
        try {
            isValid = client.updateCityPost(mongoId, title, body);
            successCallback.invoke(isValid);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void updatePlacePost(String mongoId, String title, String body, Callback errorCallback,
                                Callback successCallback) {
        boolean isValid;

        try {
            isValid = client.updatePlacePost(mongoId, title, body);
            successCallback.invoke(isValid);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }
}
