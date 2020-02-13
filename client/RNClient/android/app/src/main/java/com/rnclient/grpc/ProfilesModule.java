package com.rnclient.grpc;

import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.wcity.grpc.ProfilesClient;


import com.facebook.react.bridge.Callback;
import com.wcity.grpc.User;

public class ProfilesModule extends ReactContextBaseJavaModule {

    private static ReactApplicationContext reactContext;

    private static final String DURATION_SHORT_KEY = "SHORT";
    private static final String DURATION_LONG_KEY = "LONG";
    private static final String IP_ADDRESS = "35.197.216.42";
    private static final int PORT_NUMBER = 60051;
    private static ProfilesClient client;

    public ProfilesModule(ReactApplicationContext context) {
        super(context);
        reactContext = context;
        client = new ProfilesClient(IP_ADDRESS, PORT_NUMBER);
    }

    @Override
    public String getName() {
        return "ProfilesModule";
    }

    @ReactMethod
    public void getUser(String token, String email, Callback errorCallback,
                        Callback successCallback) {
        User user = client.getUser(token, email);

        try {
           // if (user.isValid() == true)
                successCallback.invoke(user.getName(), user.getDescription());
            //else
        } catch (Exception e) {
            errorCallback.invoke(user);
        }
    }

    @ReactMethod
    public void updateUser(String token, String email, String name, String description, Callback errorCallback,
                           Callback successCallback) {
        User user = client.updateUser(token, email, name, description);
        try {
            // if (user.isValid() == true)
            successCallback.invoke(user.getName(), user.getDescription());
            //else
        } catch (Exception e) {
            errorCallback.invoke(user);
        }
    }
}
