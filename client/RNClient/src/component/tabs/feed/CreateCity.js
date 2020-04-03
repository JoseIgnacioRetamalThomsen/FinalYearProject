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

export default class CreateCity extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            isDisabled: true,
            image: '../../../img/add_image.png',
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
        console.log("ppp", this.state.city, this.state.country)
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

    render() {
        return (
            <Root>
                <View style={{flex: 1}}>
                    <CustomHeader title="Create city" isHome={false} navigation={this.props.navigation}/>
                    <View style={styles.createContainer}>
                        <View style={{flex: 1, padding: 30}}>

                            <PhotoUpload onPhotoSelect={image => {
                                if (image) {
                                    this.setState({image: image})
                                    this.setState({isDisabled: false})
                                }
                            }
                            }>
                                <Image source={{image: this.state.image}}
                                       style={Style.profilePhoto}/>
                            </PhotoUpload>
                        </View>



                        <View style={styles.createInputContainer}>
                            <Text style={styles.inputs} >
                                {this.state.city}
                            </Text>
                        </View>

                        <View style={styles.createInputContainer}>
                            <Text style={styles.inputs}>
                                {this.state.country}
                            </Text>
                        </View>
                        <View style={styles.descInputContainer}>
                            <TextInput
                                style={styles.inputs}
                                placeholder="Description"
                                onChangeText={(description) => this.setState({description})}/>
                            <View style={styles.container}>
                            </View>
                        </View>

                        <View styles={{flex: 1}}>
                            <TouchableHighlight disabled={this.state.isDisabled} style={[styles.buttonContainer, styles.loginButton]}
                                                onPress={() => this.createCity()}>
                                <Text style={styles.loginText}>Submit</Text>
                            </TouchableHighlight>
                        </View>
                    </View>
                </View>
            </Root>
        )
    }
}
