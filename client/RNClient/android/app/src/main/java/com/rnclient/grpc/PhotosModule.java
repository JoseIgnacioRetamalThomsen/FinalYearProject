package com.rnclient.grpc;

import com.facebook.react.bridge.Arguments;
import com.facebook.react.bridge.Callback;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.ReadableArray;
import com.facebook.react.bridge.ReadableMap;
import com.facebook.react.bridge.ReadableMapKeySetIterator;
import com.facebook.react.bridge.WritableMap;
import com.google.common.primitives.Ints;
import com.google.gson.Gson;
import com.wcity.grpc.clients.PhotosClient;
import com.wcity.grpc.objects.CityPhoto;
import com.wcity.grpc.objects.CityPhotoResponse;
import com.wcity.grpc.objects.PlacePhoto;
import com.wcity.grpc.objects.PlacePhotoResponse;
import com.wcity.grpc.objects.PostPhoto;
import com.wcity.grpc.objects.PostPhotoResponse;
import com.wcity.grpc.objects.ProfilePhoto;
import com.wcity.grpc.objects.ProfilePhotoResponse;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.Map;

import io.grpc.wcity.photoShared.PostType;


public class PhotosModule extends ReactContextBaseJavaModule {

    private static ReactApplicationContext reactContext;

    private static final String DURATION_SHORT_KEY = "SHORT";
    private static final String DURATION_LONG_KEY = "LONG";
    private static final String IP_ADDRESS = "35.197.242.214";
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

        if (response.isError()) {
            errorCallback.invoke(response.getError());
            return;
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
        Gson gson = new Gson();

        CityPhotoResponse response = client.getCityImage(token, email, cityId);

        if (response.error != null) {
            errorCallback.invoke(response.error);
            return;
        }
        try {
            successCallback.invoke(gson.toJson(response.getPhotos()));
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getPlacePhoto(String token, String email, int placeId, Callback errorCallback,
                              Callback successCallback) {
        Gson gson = new Gson();
        PlacePhotoResponse response = client.getPlacePhoto(token, email, placeId);
        try {
            successCallback.invoke(gson.toJson(response.getPhotos()));
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void uploadPlacePhoto(String token, String email, int placeId, String image, int placeCityId,
                                 Callback errorCallback, Callback successCallback) {
        PlacePhoto response = client.uploadPlacePhoto(token, email, placeId, image, placeCityId);
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
        try {
            successCallback.invoke(response.isValid(), response.getPostId(),
                    response.getUserEmail(), response.getPhotos());
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void uploadPostImage(String token, String userEmail, String postId, String image,
                                int type, int parentId, Callback errorCallback, Callback successCallback) {
        PostPhoto response = client.uploadPostImage(token, userEmail, postId, image, type, parentId);
        if (response.getError() != null) {
            errorCallback.invoke(response.getError());
            return;
        }
        try {
            successCallback.invoke(response.getUrl());
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getCitysPhoto(String token, String email, Callback errorCallback,
                              Callback successCallback) {
        List<CityPhoto> response = client.getCitysPhoto(token, email);
        WritableMap map = Arguments.createMap();
        for (CityPhoto e : response) {
            map.putString("" + e.getCityId(), e.getUrl());
        }
        try {
            successCallback.invoke(map);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getPlacesPerCityPhoto(String token, String email, int cityId, Callback errorCallback,
                                      Callback successCallback) {
        List<PlacePhoto> response = client.getPlacesPerCityPhoto(token, email, cityId);
        WritableMap map = Arguments.createMap();
        for (PlacePhoto e : response) {
            map.putString("" + e.getPlaceId(), e.getUrl());
        }
        try {
            successCallback.invoke(map);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getPostsPhotosIdP(String token, String email, int type, int parentId,
                                  Callback errorCallback, Callback successCallback) {
        List<PostPhoto> response = client.getPostsPhotosIdP(token, email, type, parentId);

        WritableMap map = Arguments.createMap();
        for (PostPhoto e : response) {
            map.putString("" + e.getPostId(), e.getUrl());
        }

        try {
            successCallback.invoke(map);
        } catch (Exception e) {

            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getVisitedCitysPhotos(String email, String token, ReadableArray cityId, Callback errorCallback,
                                      Callback successCallback) {
        List<Integer> list = new ArrayList<>();

      for(Object e: cityId.toArrayList() ){
          list.add(((Double)e).intValue());
      }

        List<CityPhoto> response = client.getVisitedCitysPhotos(email, token, list);
        WritableMap map = Arguments.createMap();
        for (CityPhoto e : response) {
            map.putString("" + e.getCityId(), e.getUrl());
        }
        try {
            successCallback.invoke(map);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void getVisitedPlacesPhotos(String email, String token, ReadableArray placeId, Callback errorCallback,
                                       Callback successCallback) {
        List<Integer> list = new ArrayList<>();

        for(Object e: placeId.toArrayList() ){
            list.add(((Double)e).intValue());
        }
        List<PlacePhoto> response = client.getVisitedPlacesPhotos(email, token, list);
        WritableMap map = Arguments.createMap();
        for (PlacePhoto e : response) {
            map.putString("" + e.getPlaceId(), e.getUrl());
        }
        try {
            successCallback.invoke(map);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }
}
