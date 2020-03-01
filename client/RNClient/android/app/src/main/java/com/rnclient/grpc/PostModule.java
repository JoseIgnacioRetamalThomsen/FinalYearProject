package com.rnclient.grpc;

import com.facebook.react.bridge.Callback;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.ReadableArray;
import com.wcity.grpc.CityPostResponse;
import com.wcity.grpc.PlacePostResponse;
import com.wcity.grpc.PostClient;

import java.util.ArrayList;
import java.util.List;


public class PostModule extends ReactContextBaseJavaModule {

    private static ReactApplicationContext reactContext;

    private static final String DURATION_SHORT_KEY = "SHORT";
    private static final String DURATION_LONG_KEY = "LONG";
    private static final String IP_ADDRESS = "35.197.216.42";
    private static final int PORT_NUMBER = 60051;
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
                               String title, String body, String timeStamp, ReadableArray likes, String mongoId, Callback errorCallback,
                               Callback successCallback) {
        int index;

        int response = client.createCityPost(indexId, creatorEmail, cityName, cityCountry, title, body, timeStamp, likes.toArrayList(), mongoId);
        try {
            if (response == 0) {
                index = -1;
            } else index = response;
            successCallback.invoke(index);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }
    @ReactMethod
    public void createPlacePost(int indexId, String creatorEmail, String cityName, String countryName,
                                String placeName, String title, String body, String timeStamp, List <String> likes, String mongoId, Callback errorCallback,
                               Callback successCallback) {
        int index;
        int response = client.createPlacePost(indexId, creatorEmail, cityName, countryName, placeName, title, body, timeStamp, likes, mongoId);
        try {
            if (response == 0) {
                index = -1;
            } else index = response;
            successCallback.invoke(index);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getPlacePosts(int indexId, Callback errorCallback, Callback successCallback) {

        PlacePostResponse response = client.getPlacePosts(indexId);
        try {
            if (response == null) {
                successCallback.invoke(null);
            } else {
                successCallback.invoke(response);
            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }
    @ReactMethod
    public void getCityPosts(int indexId, Callback errorCallback, Callback successCallback) {

        CityPostResponse response = client.getCityPosts(indexId);
        try {
            if (response == null) {
                successCallback.invoke(null);
            } else {
                successCallback.invoke(response);
            }
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }
}
