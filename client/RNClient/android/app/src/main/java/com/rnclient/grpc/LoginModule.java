//LoginModule
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
    private static final int PORT_NUMBER = 50051;
    private static final String IP_ADDRESS = "104.40.206.141";
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
        LoginClient client = new LoginClient("104.40.206.141", 50051);
            String token;
            String res = client.createUser(email,password);

            try {
                if(res == null) {
                   token = "User has been already created";
                    errorCallback.invoke(token);
                     } else token = res;
                successCallback.invoke(token);
              } catch (Exception e) {
                errorCallback.invoke(e.getMessage());
              }
        }
        /*@ReactMethod
            public void updateUser(String email,String password,Callback errorCallback,
            Callback successCallback) {
                LoginClient client = new LoginClient("104.40.206.141", 50051);

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
            LoginClient client = new LoginClient("104.40.206.141", 50051);
                String token;
                String res = client.loginUser(email,password);

                try {
                    if(res == null) {
                       token = "User is already created";
                        errorCallback.invoke(token);
                         } else token = res;
                    successCallback.invoke(token);
                  } catch (Exception e) {
                    errorCallback.invoke(e.getMessage());
                  }
            }
/*
                    @ReactMethod
                    public void checkToken(String token,String email,Callback errorCallback,
                    Callback successCallback) {
                    LoginClient client = new LoginClient("104.40.206.141", 50051);
                    boolean result = client.checkToken(token, email);
                    }

                    @ReactMethod
                    public void logout(String token,String email,Callback errorCallback,
                    Callback successCallback) {
                    LoginClient client = new LoginClient("104.40.206.141", 50051);
                    boolean result = client.logout(token, email);
                    }*/
}
