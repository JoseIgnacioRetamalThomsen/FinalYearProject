import React from 'react'
import {Button, TextInput, View, NativeModules} from "react-native";
import CustomHeader from "./CustomHeader";
import styles from "../styles/Style";
import AsyncStorage from "@react-native-community/async-storage";

export default class Place extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            name: '',
            city: '',
            country: '',
            description: '',
            lat: '',
            lon: '',
        }
    }

    async addPlace() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in city " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.createPlace(
                            value,
                            this.state.name,
                            this.state.city,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lon,

                            (err) => {
                                console.log("err in createPlace " + err)
                            },
                            (name, country, email, description, lat, lon) => {
                                console.log("name, country, email, description, lat, lon  is " + name, country, email, description, lat, lon)
                                console.log("successfully created a place!!!")
                            })
                    }
                })
            })
        })
    }

    async getPlace() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in getPlace " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.getPlace(
                            value,
                            this.state.name,
                            this.state.city,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lon,

                            (err) => {
                                console.log("err in createPlace " + err)
                            },
                            (name, city, country, email, description, lat, lon) => {
                                console.log("name, city, country, email, description, lat, lon  is " + name, city, country, email, description, lat, lon)
                                console.log("successfully got a place!!!")
                            })
                    }
                })
            })
        })
    }

    render() {
        return (
            <View style={{flex: 1}}>


                <View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>

                    <TextInput
                        style={styles.inputs}
                        placeholder="Name"
                        onChangeText={(name) => this.setState({name})}/>
                    <TextInput
                        style={styles.inputs}
                        placeholder="City"
                        onChangeText={(city) => this.setState({city})}/>
                    <TextInput
                        style={styles.inputs}
                        placeholder="Country"
                        onChangeText={(country) => this.setState({country})}/>
                    <TextInput
                        style={styles.inputs}
                        placeholder="Description"
                        onChangeText={(description) => this.setState({description})}/>
                    <TextInput
                        style={styles.inputs}
                        placeholder="Lat"
                        onChangeText={(lat) => this.setState({lat})}/>
                    <TextInput
                        style={styles.inputs}
                        placeholder="Lon"
                        onChangeText={(lon) => this.setState({lon})}/>
                    <Button title="Add place"
                            onPress={() => this.addPlace()}/>
                </View>
            </View>
        )
    }
}
