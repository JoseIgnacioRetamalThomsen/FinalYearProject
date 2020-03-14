package com.wcity.grpc.clients;

import com.wcity.grpc.CityPostResponse;
import com.wcity.grpc.MyCityPost;
import com.wcity.grpc.MyPlacePost;
import com.wcity.grpc.PlacePostResponse;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.wcity.postservice.CityPost;
import io.grpc.wcity.postservice.CityPostsResponse;
import io.grpc.wcity.postservice.CreatePostResponse;
import io.grpc.wcity.postservice.PlacePost;
import io.grpc.wcity.postservice.PlacePostsResponse;
import io.grpc.wcity.postservice.PostsRequest;
import io.grpc.wcity.postservice.PostsServiceGrpc;
import io.grpc.wcity.postservice.UpdatePostRequest;
import io.grpc.wcity.postservice.UpdatePostResponse;

public class PostClient {

    private final ManagedChannel channel;
    private final PostsServiceGrpc.PostsServiceBlockingStub stub;

    public PostClient(String host, int port) {
        this.channel = ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .build();
        stub = PostsServiceGrpc.newBlockingStub(channel);
    }

    public void shutdown() throws InterruptedException {
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
    }


    public int createCityPost(int indexId, String creatorEmail, String cityName, String cityCountry,
                              String title, String body) {

        CityPost cityPost = CityPost.newBuilder()
                .setIndexId(indexId)
                .setCreatorEmail(creatorEmail)
                .setCityName(cityName)
                .setCityCountry(cityCountry)
                .setTitle(title)
                .setBody(body)
                .build();
        CreatePostResponse response;
        int index;
        try {
            response = stub.createCityPost(cityPost);
            //  if(response.getValied() == true)
            index = response.getIndexId();
            // else index = -1;
        } catch (StatusRuntimeException e) {
            e.getMessage();
            index = -999;
        }
        return index;
    }


    public int createPlacePost(int indexId, String creatorEmail, String cityName, String countryName,
                               String placeName, String title, String body) {

        PlacePost request = PlacePost.newBuilder()
                .setIndexId(indexId)
                .setCreatorEmail(creatorEmail)
                .setCityName(cityName)
                .setCountryName(countryName)
                .setPlaceName(placeName)
                .setTitle(title)
                .setBody(body)
                .build();
        CreatePostResponse response;
        int index = 0;
        try {
            response = stub.createPlacePost(request);
            index = response.getIndexId();
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return index;
    }

    public PlacePostResponse getPlacePosts(int indexId) {

        PostsRequest postsRequest = PostsRequest.newBuilder()
                .setIndexId(indexId)
                .build();

        PlacePostsResponse response;
        PlacePostResponse placePostResponse = null;

        try {
            response = stub.getPlacePosts(postsRequest);

            ArrayList<MyPlacePost> myPlacePostList = new ArrayList<>();
            for (io.grpc.wcity.postservice.PlacePost post : response.getPostsList()) {
                myPlacePostList.add(new MyPlacePost(post.getIndexId(), post.getCreatorEmail(),
                        post.getCityName(), post.getCountryName(), post.getPlaceName(),
                        post.getTitle(), post.getBody(), post.getTimeStamp(), post.getLikesList(),
                        post.getMongoId()));
            }
            placePostResponse = new PlacePostResponse(response.getValid(), response.getIndexId(), myPlacePostList);

        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return placePostResponse;
    }

    public CityPostResponse getCityPosts(int indexId) {

        PostsRequest postsRequest = PostsRequest.newBuilder()
                .setIndexId(indexId)
                .build();

        CityPostsResponse response;
        CityPostResponse cityPostResponse = null;

        try {
            response = stub.getCityPosts(postsRequest);

            List<MyCityPost> cityPostList = new ArrayList<>();
            for (io.grpc.wcity.postservice.CityPost post : response.getPostsList()) {
                cityPostList.add(new MyCityPost(post.getIndexId(), post.getCreatorEmail(),
                        post.getCityName(), post.getCityCountry(), post.getTitle(),
                        post.getBody(), post.getTimeStamp(), post.getLikesList(),
                        post.getMongoId()));
            }
            cityPostResponse = new CityPostResponse(response.getValid(), response.getIndexId(), cityPostList);

        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return cityPostResponse;
    }


    public boolean updateCityPost(String mongoId, String title, String body) {
        UpdatePostRequest request = UpdatePostRequest.newBuilder()
                .setMongoId(mongoId)
                .setTitle(title)
                .setBody(body)
                .build();

        UpdatePostResponse response;
        boolean isValid;
        try {
            response = stub.updateCityPost(request);
            isValid = response.getValid();
        } catch (StatusRuntimeException e) {
            e.getMessage();
            isValid = false;
        }
        return isValid;
    }


    public boolean updatePlacePost(String mongoId, String title, String body) {
        UpdatePostRequest request = UpdatePostRequest.newBuilder()
                .setMongoId(mongoId)
                .setTitle(title)
                .setBody(body)
                .build();

        UpdatePostResponse response;
        boolean isValid;
        try {
            response = stub.updateCityPost(request);
            isValid = response.getValid();
        } catch (StatusRuntimeException e) {
            e.getMessage();
            isValid = false;
        }
        return isValid;
    }
}
