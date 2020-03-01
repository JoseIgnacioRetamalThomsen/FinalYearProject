package com.wcity.grpc;

import java.util.ArrayList;
import java.util.List;

public class PlacePostResponse {
    private boolean isValid;
    private  int index;
    private List<MyPlacePost> myPlacePosts;

    public PlacePostResponse(boolean isValid, int index, List<MyPlacePost> myPlacePosts) {
        this.isValid = isValid;
        this.index = index;
        this.myPlacePosts = myPlacePosts;
    }

    public boolean isValid() {
        return isValid;
    }

    public void setValid(boolean valid) {
        isValid = valid;
    }

    public int getIndex() {
        return index;
    }

    public void setIndex(int index) {
        this.index = index;
    }

    public List<MyPlacePost> getMyPlacePosts() {
        return myPlacePosts;
    }

    public void setMyPlacePosts(ArrayList<MyPlacePost> myPlacePosts) {
        this.myPlacePosts = myPlacePosts;
    }
}
