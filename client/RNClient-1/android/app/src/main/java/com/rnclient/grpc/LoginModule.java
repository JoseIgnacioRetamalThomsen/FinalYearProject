package com.rnclient.grpc;

import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.wcity.grpc.LoginClient;
import io.grpc.wcity.login.UserResponse;

import com.facebook.react.bridge.Callback;

public class LoginModule  extends ReactContextBaseJavaModule {

    private static ReactApplicationContext reactContext;

    private static final String DURATION_SHORT_KEY = "SHORT";
    private static final String DURATION_LONG_KEY = "LONG";
    private static final String IP_ADDRESS = "35.197.216.42";
    private static final int PORT_NUMBER = 50051;
    private static LoginClient client;

   public  LoginModule(ReactApplicationContext context) {
        super(context);
        reactContext = context;
        client = new LoginClient(IP_ADDRESS, PORT_NUMBER);
    }

    @Override
    public String getName() {
        return "LoginModule";
    }

     @ReactMethod
        public void createUser(String email,String password,Callback errorCallback,
        Callback successCallback) {
        LoginClient client = new LoginClient("35.197.216.42", 50051);
            String token = "";
            String response = client.createUser(email,password);
            try {
                if(response == null) {
                   token = "StatusRuntimeException";
                   // errorCallback.invoke(e.getMessage());
                     } else token = response;
                successCallback.invoke(token);
              } catch (Exception e) {
                errorCallback.invoke(e.getMessage());
              }
/*
                          if(token == "") {
                              errorCallback.invoke(token);
                           } else successCallback.invoke(token);*/
        }

        /* @ReactMethod
        public void updateUser(String email,String password,Callback errorCallback,
        Callback successCallback) {
            LoginClient client = new LoginClient("35.197.216.42", 50051);

            String msg;
             String res = client.updateUser(email,password);
            if (res==null) msg="null";
            else msg = res;

            try {
               successCallback.invoke(msg);
              } catch (Exception e) {
                errorCallback.invoke(e.getMessage());
              }
        }*/

        @ReactMethod
        public void loginUser(String email,String password,Callback errorCallback,
                 Callback successCallback) {
            LoginClient client = new LoginClient("35.197.216.42", 50051);
           /* String token = client.loginUser(email,password);
            if(token == null) {
                errorCallback.invoke("");
             } else successCallback.invoke(token);*/
             String msg;
             String res = client.loginUser(email,password);
             if (res==null) {
             msg="null";
              errorCallback.invoke(msg);
             }
             else msg = res;

             try {
                successCallback.invoke(msg);
               } catch (Exception e) {
                 errorCallback.invoke(msg);
               }
        }

        @ReactMethod
        public void checkToken(String token,String email,Callback errorCallback,
            Callback successCallback) {
            LoginClient client = new LoginClient("35.197.216.42", 50051);
            boolean result = client.checkToken(token, email);


            if(result == false) {
                errorCallback.invoke(false);
            } else successCallback.invoke(true);
        }


            @ReactMethod
            public void logout(String token, String email, Callback errorCallback,
            Callback successCallback) {
            boolean result = client.logout(token, email);

           if(result == false) {
               errorCallback.invoke(false);
           } else successCallback.invoke(true);
        }
}
