import React from 'react'
import {Button, TextInput, View, NativeModules, Image, ScrollView, Icon, TouchableHighlight} from "react-native";
import styles from "../../../styles/Style";
import AsyncStorage from "@react-native-community/async-storage";
import {Card, CardTitle, CardContent, CardAction, CardButton, CardImage} from 'react-native-material-cards'
import {Body, CardItem, Text, Title, Root} from "native-base";
import ActionButton from "react-native-action-button";
import CustomHeader from "../../headers/CustomHeader";
import PhotoUpload from "react-native-photo-upload";
import Style from '../../../styles/Style'
import {IMAGE} from "../../../constants/Image";

export default class CreateCity extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            image: '',
            description:'',
            cities: [
                {
                    cityId: 0,
                    token: '',
                    name: '',
                    country: '',
                    email: '',
                    description: '',
                    lat: 60,
                    lon: 80,
                    image: '',
                    url: '',
                },

            ],
            isUpdated: true
        }
    }

    callbackFunction = (lat, lng, fcity, country) => {
        this.setState({lat: lat})
        this.setState({lng: lng})
        this.setState({city: fcity})
        this.setState({country: country})
        console.log('!!!!', this.state.fcity, this.state.country)
    }

    componentDidMount() {
        const city = this.props.navigation.getParam('city', '')
        const country = this.props.navigation.getParam('country', '')
        this.setState({
            city,
            country,
        })
    }

    createCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    NativeModules.ProfilesModule.createCity(
                        value,
                        key,
                        this.state.city,
                        this.state.country,
                        key,
                        89,
                        90,
                        //parseFloat(this.state.lon),
                        this.state.description,
                        (err) => {
                            console.log(err)
                        },
                        (cityId) => {
                            this.setState({cityId: cityId})

                            this.props.navigation.navigate('CityDetail', {
                                cityId: this.state.cityId,
                                name: this.state.city,
                                indexId: this.state.indexId,
                                country: this.state.country,
                                description: this.state.description,
                            })
                        })
                    this.uploadCityPhoto()
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

                    if (value !== null) {
                        NativeModules.PhotosModule.uploadCityPhoto(
                            key,
                            value,
                            parseInt(this.state.cityId),
                            this.state.image,
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

    onClick() {
        console.log("mmmm", this.state.description, this.state.image)
        if (this.state.description === undefined || this.state.image === '') {
            alert("Please upload photo and provide description")
        } else {
            this.createCity()
        }
    }

    displayPhoto() {
        if (this.state.image === '') {
            return (<Image source={IMAGE.UPLOAD_IMG}
                           style={Style.uploadPhoto}/>)
        } else {
            return (<Image source={{uri: this.state.image}}
                           style={Style.uploadPhoto}/>)
        }
    }

    render() {
        return (
            <Root>
                <View style={{flex: 1}}>
                    <CustomHeader title="Create city" isHome={false} navigation={this.props.navigation}/>

                    <Card style={styles.createContainer} >
                        <PhotoUpload  onPhotoSelect={image => {
                            if (image) {
                                this.setState({image: image})
                            }
                        }
                        }>
                            {this.displayPhoto()}
                        </PhotoUpload>

                        <CardItem style={styles.createInputContainer}>
                            <Text style={styles.inputs}>
                                {this.state.city}
                            </Text>
                        </CardItem>

                        <CardItem style={styles.createInputContainer}>
                            <Text style={styles.inputs}>
                                {this.state.country}
                            </Text>
                        </CardItem>

                        <CardItem style={styles.descInputContainer}>
                            <TextInput
                                style={styles.inputs}
                                placeholder="Enter description"
                                onChangeText={(description) => this.setState({description})}/>
                        </CardItem>

                        <TouchableHighlight
                                            style={[styles.buttonContainer, styles.loginButton]}
                                            onPress={() => this.onClick()}>
                            <Text style={styles.loginText}>Submit</Text>
                        </TouchableHighlight>
                    </Card>
                </View>
            </Root>
        )
    }
}
