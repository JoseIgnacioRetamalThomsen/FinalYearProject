package com.rnclient.grpc;

import com.wcity.grpc.HelloWorldClient;

import com.facebook.react.bridge.NativeModule;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;

import java.util.Map;
import java.util.HashMap;

public class HelloWorldModule extends ReactContextBaseJavaModule {

    private static ReactApplicationContext reactContext;

    private static final String DURATION_SHORT_KEY = "SHORT";
    private static final String DURATION_LONG_KEY = "LONG";

    HelloWorldModule(ReactApplicationContext context) {
        super(context);
        reactContext = context;
    }

    @Override
    public String getName() {
        return "HelloWorldModule";
    }

    @ReactMethod
    public void sayHello(String message) {
       HelloWorldClient client = new HelloWorldClient("104.40.206.141", 7776);

        client.greet(message);
    }
}
