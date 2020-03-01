package com.wcity.grpc;


import java.util.List;

public class CityPostResponse {
    private boolean isValid;
    private  int index;
    private List<MyCityPost> myCityPosts;

    public CityPostResponse(boolean isValid, int index, List<MyCityPost> myCityPosts) {
        this.isValid = isValid;
        this.index = index;
        this.myCityPosts = myCityPosts;
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

    public List<MyCityPost> getMyCityPosts() {
        return myCityPosts;
    }

    public void setMyCityPosts(List<MyCityPost> myCityPosts) {
        this.myCityPosts = myCityPosts;
    }
}
