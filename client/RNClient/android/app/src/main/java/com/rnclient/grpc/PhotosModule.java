package com.rnclient.grpc;

import com.facebook.react.bridge.Callback;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.google.protobuf.ByteString;
import com.wcity.grpc.clients.PhotosClient;
import com.wcity.grpc.objects.CityPhoto;
import com.wcity.grpc.objects.PlacePhoto;
import com.wcity.grpc.objects.PostPhoto;
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
    public void getProfilePhoto(String email, String token, Callback errorCallback, Callback successCallback) {

        ProfilePhotoResponse response = client.getProfilePhoto(email, token);

        try {
            if (response == null) {
                successCallback.invoke();
            } else successCallback.invoke(response);
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }

    @ReactMethod
    public void uploadProfilePhoto(String email, String token, String image, Callback errorCallback, Callback successCallback) {

        ProfilePhoto response = client.uploadProfilePhoto(email, token, image);
        try {
            if (response.error != null) {
                errorCallback.invoke(response.error);
            } else successCallback.invoke(response.getUrl());
        } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
        }
    }
}
