package com.rnclient.grpc;

import com.facebook.react.bridge.Callback;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.wcity.grpc.clients.PhotosClient;
import com.wcity.grpc.objects.CityPhoto;
import com.wcity.grpc.objects.CityPhotoResponse;
import com.wcity.grpc.objects.PlacePhoto;
import com.wcity.grpc.objects.PlacePhotoResponse;
import com.wcity.grpc.objects.PostPhoto;
import com.wcity.grpc.objects.PostPhotoResponse;
import com.wcity.grpc.objects.ProfilePhoto;
import com.wcity.grpc.objects.ProfilePhotoResponse;

public class PhotosModule extends ReactContextBaseJavaModule {

    private static ReactApplicationContext reactContext;

    private static final String DURATION_SHORT_KEY = "SHORT";
    private static final String DURATION_LONG_KEY = "LONG";
    private static final String IP_ADDRESS = "35.197.216.42";
    private static final int PORT_NUMBER = 30051;
    private static PhotosClient client;

    public PhotosModule(ReactApplicationContext context) {
        super(context);
        reactContext = context;
        client = new PhotosClient(IP_ADDRESS, PORT_NUMBER);
    }

    @Override
    public String getName() {
        return "PhotosModule";
    }

    @ReactMethod
    public void getProfilePhoto(String email, String token, Callback errorCallback,
                                Callback successCallback) {

        ProfilePhotoResponse response = client.getProfilePhoto(email, token);

        try {
            successCallback.invoke(response.getPhoto().getUrl());
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void uploadProfilePhoto(String email, String token, String image,
                                   Callback errorCallback, Callback successCallback) {

        ProfilePhoto response = client.uploadProfilePhoto(email, token, image);
        try {
            successCallback.invoke(response.getUrl());
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void uploadCityPhoto(String token, String email, int cityId, String image,
                                Callback errorCallback, Callback successCallback) {
        CityPhoto response = client.uploadCityPhoto(token, email, cityId, image);

        if (response.isError()){
            errorCallback.invoke(response.getError()); return;
        }
        try {
            successCallback.invoke(response.getUrl());
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getCityImage(String token, String email, int cityId,
                             Callback errorCallback, Callback successCallback) {
        CityPhotoResponse response = client.getCityImage(email, token, cityId);
        int id = 0;
        int cId = 0;
        String url = "";
        String timestamp = "";
        boolean selected = false;

        try {
            for (int i = 0; i < response.getPhotos().size(); i++) {
//                 id = response.getPhotos().get(i).getId();
//                 cId = response.getPhotos().get(i).getCityId();
                url = response.getPhotos().get(i).getUrl();
//                 timestamp = response.getPhotos().get(i).getTimestamp();
//                 selected = response.getPhotos().get(i).isSelected();
            }
//            successCallback.invoke(id, cId, url, timestamp, selected);
            successCallback.invoke(url);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getPlacePhoto(String token, String email, int placeId, Callback errorCallback,
                              Callback successCallback) {
        PlacePhotoResponse response = client.getPlacePhoto(token, email, placeId);
        int id = 0;
        int pId = 0;
        String url = "";
        String timestamp = "";
        boolean selected = false;

        try {
            for (int i = 0; i < response.getPhotos().size(); i++) {
//                 id = response.getPhotos().get(i).getId();
//                 pId = response.getPhotos().get(i).getCityId();
                url = response.getPhotos().get(i).getUrl();
//                 timestamp = response.getPhotos().get(i).getTimestamp();
//                 selected = response.getPhotos().get(i).isSelected();
            }
//            successCallback.invoke(id, cId, url, timestamp, selected);
            successCallback.invoke(url);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void uploadPlacePhoto(String token, String email, int placeId, String image, Callback errorCallback,
                                 Callback successCallback) {
        PlacePhoto response = client.uploadPlacePhoto(token, email, placeId, image);
        try {
            successCallback.invoke(response.getUrl());
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getPostImage(String token, String userEmail, String postId, Callback errorCallback,
                             Callback successCallback) {
        PostPhotoResponse response = client.getPostImage(token, userEmail, postId);
        int id = 0;
        int pId = 0;
        String url = "";
        String timestamp = "";
        boolean selected = false;

        try {
            for (int i = 0; i < response.getPhotos().size(); i++) {
//                 id = response.getPhotos().get(i).getId();
//                 pId = response.getPhotos().get(i).getPostId();
                url = response.getPhotos().get(i).getUrl();
//                 timestamp = response.getPhotos().get(i).getTimestamp();
//                 selected = response.getPhotos().get(i).isSelected();
            }
//            successCallback.invoke(id, cId, url, timestamp, selected);
            successCallback.invoke(url);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void uploadPostImage(String token, String postId, String userEmail, String image, Callback errorCallback,
                                Callback successCallback) {
        PostPhoto response = client.uploadPostImage(token, postId, userEmail, image);
        try {
            successCallback.invoke(response.getUrl());
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

}
