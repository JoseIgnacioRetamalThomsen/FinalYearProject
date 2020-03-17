import React from 'react'
import {Button, TextInput, View, NativeModules, Image, ScrollView, Icon, TouchableHighlight} from "react-native";
import styles from "../../../styles/Style";
import AsyncStorage from "@react-native-community/async-storage";
import {Card, CardTitle, CardContent, CardAction, CardButton, CardImage} from 'react-native-material-cards'
import {Body, CardItem, Text, Title, Root} from "native-base";
import CustomHeader from "../../CustomHeader";
import PhotoUpload from "react-native-photo-upload";

export default class CreatePlace extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            placeId: 0,
            cityId:'',
            image: '',
            url: '',
            placeName: '',
            city: '',
            country: '',
            lat: 1,
            lon: 1,
            description: '',
        }
    }

    componentDidMount() {
        const cityId = this.props.navigation.getParam('cityId', '')
        const city = this.props.navigation.getParam('city', '')
        const country = this.props.navigation.getParam('country', '')
        const email = this.props.navigation.getParam('email', '')
        const image = this.props.navigation.getParam('image', '')
        console.log("Passed values: in createPlace", city, country, email, image)
        this.setState({
            cityId,
            city,
            country,
            email,
            image
        })
    }

    createPlace() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]
                    NativeModules.ProfilesModule.createPlace(
                        token,
                        email,
                        this.state.placeName,
                        this.state.city,
                        this.state.country,
                        89,
                        90,
                        //parseFloat(this.state.lon),
                        this.state.description,
                        (err) => {
                            console.log(err)
                        },
                        (placeId) => {
                            this.setState({placeId: placeId})
                            console.log(placeId)
                            // console.log(this.state.placeName, this.state.description)
                            this.uploadPlacePhoto();
                            //this.props.navigation.navigate('DisplayPlaces')
                        })
                })
            })
        })
    }

    uploadPlacePhoto() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]
                    if (token != null) {
                        NativeModules.PhotosModule.uploadPlacePhoto(
                            token,
                            email,
                            parseInt(this.state.placeId),
                            this.state.image,
                            this.state.cityId,
                            (err) => {
                                console.log(err)
                            },
                            (url) => {
                                this.setState({url: url})
                            })
                    }
                })
            })
        })
    }

    render() {
        return (
            <Root>
                <View style={{flex: 1}}>
                    <CustomHeader title="Create place" isHome={false} navigation={this.props.navigation}/>
                    <View style={styles.container}>
                        {/*<GeoLoc parentCallback={this.callbackFunction} />*/}
                        <View style={styles.inputContainer}>
                            <TextInput
                                style={styles.inputs}
                                placeholder="Place Name"
                                underlineColorAndroid='transparent'
                                onChangeText={(placeName) => this.setState({placeName})}
                            >
                                {this.state.placeName}
                            </TextInput>
                        </View>

                        <View style={styles.inputContainer}>
                            <TextInput
                                style={styles.inputs}
                                placeholder="Description"
                                onChangeText={(description) => this.setState({description})}/>
                            <View style={styles.container}>
                            </View>
                        </View>

                        <PhotoUpload onPhotoSelect={image => {
                            if (image) {
                                this.setState({image: image})
                                //this.createPlace()
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

                    <View style={styles.container}>
                        <TouchableHighlight style={[styles.buttonContainer, styles.loginButton]}
                                            onPress={() => this.createPlace()}>
                            <Text style={styles.loginText}>Submit</Text>
                        </TouchableHighlight>
                    </View>
                </View>
            </Root>
        )
    }
}
