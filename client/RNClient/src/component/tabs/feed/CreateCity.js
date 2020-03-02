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

export default class CreateCity extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            cities: [
                {
                    cityId: 0,
                    token: '',
                    name: '',
                    country: '',
                    email: '',
                    description: '',
                    lat: 55,
                    lon: 65,
                    img:'',
                },

            ]

        }
    }

    // sendData = () => {
    //     this.props.parentCallback(this.state.city, this.state.cityId)
    //     console.log("sendData"+this.state.city, this.state.cityId)
    // }

    callbackFunction = ( city, country, lat, lon) => {
        this.setState({city: city})
        this.setState({country: country})
        this.setState({lat: lat})
        this.setState({lon: lon})
        console.log(lat, lon)
    }

    createCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in city " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.createCity(
                            value,
                            this.state.city,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lon,
                            (err) => {
                                console.log("err in createCity " + err)
                            },
                            (valid, city, country, email, description, lat, lon, id) => {

                                console.log("valid, city, country, email, description, lat, lon id is " +
                                    valid, city, country, email, description, lat, lon, id)
                                console.log("successfully created a city!!!")
                               // this.sendData()
                                this.props.navigation.navigate('DisplayCities')
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
                    <CustomHeader title="Write a city post " isHome={false} navigation={this.props.navigation}/>
                    <View style={styles.container}>
                        <GeoLoc parentCallback={this.callbackFunction} />
                        <View style={styles.inputContainer}>
                            <Text
                                style={styles.inputs}
                                placeholder="CreateCity"
                                underlineColorAndroid='transparent'
                            >
                                {this.state.city}
                            </Text>
                        </View>

                        <View style={styles.inputContainer}>
                            <Text
                                style={styles.inputs}
                                placeholder="Country"
                            >
                                {this.state.country}
                            </Text>
                        </View>
                        <View style={styles.inputContainer}>
                            <TextInput
                                style={styles.inputs}
                                placeholder="Description"
                                onChangeText={(description) => this.setState({description})}/>
                            <View style={styles.container}>

                            </View>
                        </View>

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
