import React from 'react'
import {Button, TextInput, View, NativeModules, Image, ScrollView, Icon, TouchableHighlight} from "react-native";
import styles from "../../../styles/Style";
import AsyncStorage from "@react-native-community/async-storage";
import {Card, CardTitle, CardContent, CardAction, CardButton, CardImage} from 'react-native-material-cards'
import {Body, CardItem, Text, Title, Root} from "native-base";
import ActionButton from "react-native-action-button";
import CustomHeader from "../../CustomHeader";
import MapInput from "../../MapInput";
import LoadImage from "../../LoadImage";
import GeoLoc from "../../GeoLoc";
import PhotoUpload from "react-native-photo-upload";

export default class CreateCity extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            cities: [
                {
                    cityId: "32",
                    token: '',
                    name: '',
                    country: '',
                    email: '',
                    description: '',
                    lat: 60,
                    lon: 80,
                    image:'',
                    url: '',
                },

            ]

        }
    }

    // sendData = () => {
    //     this.props.parentCallback(this.state.city, this.state.cityId)
    //     console.log("sendData"+this.state.city, this.state.cityId)
    // }

    // callbackFunction = ( city, country, lat, lon) => {
    //     this.setState({city: city})
    //     this.setState({country: country})
    //     this.setState({lat: lat})
    //     this.setState({lon: lon})
    //     console.log(lat, lon)
    // }

    createCity() {
        console.log("Clicked")
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    NativeModules.ProfilesModule.createCity(
                        value,
                        key,
                        this.state.name,
                        this.state.country,
                        key,
                        89,
                        90,
                        //parseFloat(this.state.lon),
                        this.state.description,
                        (err) => {
                            console.log("err in createCity " + err)
                        },
                        (cityId) => {
                            this.setState({cityId:cityId})
                            this.uploadCityPhoto();
                            console.log("valid, city, country, email, description, lat, lon id is " + cityId)
                            console.log("successfully created a city!!!")
                            // this.sendData()
                            this.props.navigation.navigate('DisplayCities')
                        })
                })
            })
        })
    }

    uploadCityPhoto() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in createcity " + key + " " + value)
                    //console.log("this.state.image " + this.state.image)
                    if (value !== null) {
                        NativeModules.PhotosModule.uploadCityPhoto(
                            key,
                            value,
                            parseInt(this.state.cityId),
                            this.state.image,
                            (err) => {
                                console.log("In uploadPhoto " + err)
                            },
                            (url) => {
                                this.setState({url: url})
                                console.log("avatar_url  is " + this.state.url, url, this.state.image)
                                console.log("successful upload!!!")
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
                    <CustomHeader title="Create city" isHome={false} navigation={this.props.navigation}/>
                    <View style={styles.container}>
                        {/*<GeoLoc parentCallback={this.callbackFunction} />*/}
                        <View style={styles.inputContainer}>
                            <TextInput
                                style={styles.inputs}
                                placeholder="CreateCity"
                                underlineColorAndroid='transparent'
                                onChangeText={(name) => this.setState({name})}
                            >
                                {this.state.city}
                            </TextInput>
                        </View>

                        <View style={styles.inputContainer}>
                            <TextInput
                                style={styles.inputs}
                                placeholder="Country"
                                onChangeText={(country) => this.setState({country})}
                            >
                                {this.state.country}
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
                                //this.uploadCityPhoto()
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
                                            onPress={() => this.createCity()}>
                            <Text style={styles.loginText}>Submit</Text>
                        </TouchableHighlight>
                    </View>
                </View>
            </Root>
        )
    }
}
