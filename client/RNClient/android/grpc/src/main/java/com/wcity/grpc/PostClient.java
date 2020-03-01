package com.wcity.grpc;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.wcity.postservice.CityPost;
import io.grpc.wcity.postservice.CityPostsResponse;
import io.grpc.wcity.postservice.CreatePostResponse;
import io.grpc.wcity.postservice.PlacePostsResponse;
import io.grpc.wcity.postservice.PostsRequest;
import io.grpc.wcity.postservice.PostsServiceGrpc;

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
                              String title, String body, String timeStamp, ArrayList likes, String mongoId) {

        CityPost cityPost = CityPost.newBuilder()
                .setIndexId(indexId)
                .setCreatorEmail(creatorEmail)
                .setCityName(cityName)
                .setCityCountry(cityCountry)
                .setTitle(title)
                .setBody(body)
                .setTimeStamp(timeStamp)
                //.setLikes(0, null)
                .setMongoId(mongoId)
                .build();
        CreatePostResponse response;
        int index = 0;
        try {
            response = stub.createCityPost(cityPost);
           // if(response.getValied() == true)
                index = response.getIndexId();
           // else index = 0;
        } catch (StatusRuntimeException e) {
            e.getMessage();
        }
        return index;
    }


    public int createPlacePost(int indexId, String creatorEmail, String cityName, String countryName,
                               String placeName, String title, String body, String timeStamp, List<String> likes, String mongoId) {

        io.grpc.wcity.postservice.PlacePost placePost = io.grpc.wcity.postservice.PlacePost.newBuilder()
                .setIndexId(indexId)
                .setCreatorEmail(creatorEmail)
                .setCityName(cityName)
                .setCountryName(countryName)
                .setPlaceName(placeName)
                .setTitle(title)
                .setBody(body)
                .setTimeStamp(timeStamp)
                .setLikes(0, null)
                .setMongoId(mongoId)
                .build();
        CreatePostResponse response;
        int index = 0;
        try {
            response = stub.createPlacePost(placePost);
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

}
