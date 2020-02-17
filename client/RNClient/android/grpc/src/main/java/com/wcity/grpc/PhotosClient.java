package com.wcity.grpc;


import com.google.protobuf.ByteString;

import java.nio.charset.StandardCharsets;
import java.util.Base64;
import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.wcity.photo.PhotosServiceGrpc;

import io.grpc.wcity.photo.ProfilePhotoRequestP;
import io.grpc.wcity.photo.ProfilePhotoResponseP;
import io.grpc.wcity.photo.ProfileUploadRequest;
import io.grpc.wcity.photo.ProfileUploadResponse;

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
            e.getMessage();
        }
        return url;
    }

    public String uploadProfilePhoto(String email, String token, String image) {
        ProfileUploadRequest userData = ProfileUploadRequest.newBuilder().setEmail(email).setToken(token).setImage(ByteString.copyFrom(Base64.getMimeDecoder().decode(image.replaceFirst("^.*;base64,", "")))).build();

        ProfileUploadResponse response;
        String url = "";
        boolean isValid;
        try {
            response = stub.uploadProfilePhoto(userData);
            isValid = response.getValid();
            if (isValid == true)
                url = response.getUrl();
            else url = "";
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return url;
    }
}
