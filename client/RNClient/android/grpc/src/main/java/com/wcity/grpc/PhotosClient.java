package com.wcity.grpc;

import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.wcity.photo.PhotosServiceGrpc;

import io.grpc.wcity.photo.ProfilePhotoRequestP;
import io.grpc.wcity.photo.ProfilePhotoResponseP;

public class PhotosClient {

    private final ManagedChannel channel;
    private final PhotosServiceGrpc.PhotosServiceBlockingStub stub;

    public PhotosClient(String host, int port) {
        this.channel = ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .build();
        stub = PhotosServiceGrpc.newBlockingStub(channel);
    }

    public void shutdown() throws InterruptedException {
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
    }

    public String getProfilePhoto(String email, String token) {
        ProfilePhotoRequestP userData = ProfilePhotoRequestP.newBuilder().setEmail(email).setToken(token).build();
        ProfilePhotoResponseP response;
        String url = "";
        boolean isValid;
        try {
            response = stub.getProfilePhoto(userData);
            isValid = response.getValid();
            if (isValid == true)
                url = response.getUrl();
            else url = "";
        } catch (StatusRuntimeException e) {

        }
        return url;
    }
}
