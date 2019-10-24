package com.rnclient.grpc;

import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.wcity.grpc.HelloWorldClient;
import com.wcity.grpc.LoginClient;
import io.grpc.wcity.login.UserResponse;

import com.facebook.react.bridge.Callback;

public class LoginModule  extends ReactContextBaseJavaModule {

    private static ReactApplicationContext reactContext;

    private static final String DURATION_SHORT_KEY = "SHORT";
    private static final String DURATION_LONG_KEY = "LONG";

   public  LoginModule(ReactApplicationContext context) {
        super(context);
        reactContext = context;
    }

    @Override
    public String getName() {
        return "LoginModule";
    }

    @ReactMethod
    public void check(String email,String password,Callback errorCallback,
    Callback successCallback) {
        LoginClient client = new LoginClient("192.168.43.221", 50051);
       
        String msg ="not null";
         String res = client.Check(email,password);
        if(res==null) msg="null";
        else msg = res;
     
        try {
           // measureLayout(tag, ancestorTag, mMeasureBuffer);
          
            successCallback.invoke(msg);
          } catch (Exception e) {
            errorCallback.invoke(e.getMessage());
          }
    }
}
