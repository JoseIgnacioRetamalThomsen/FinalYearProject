import React from "react";
import {Button, Image, View} from "react-native";
import PhotoUpload from "react-native-photo-upload";

export default class Test extends React.Component {

    render() {
        return (
            <View>
                <PhotoUpload onPhotoSelect={image => {
                    if (image) {
                        this.setState({image: image})
                        this.uploadCityPhoto()
                    }
                }
                }>
                    <Image source={{image: this.state.image}}
                           style={{
                               height: 120,
                               width: 120,
                               borderRadius: 60,
                               borderColor: 'black',
                               borderWidth: 5,
                               flex: 0,
                               resizeMode: 'cover'
                           }}/>
                </PhotoUpload>
            </View>
        )
    }
}
