import React from 'react'
import {Button, NativeModules, TextInput, View} from "react-native";
import CustomHeader from "./CustomHeader";
import styles from "../styles/Style";
import AsyncStorage from "@react-native-community/async-storage";

export default class City extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            name: '',
            country: '',
            description: '',
            lat: 0,
            lon: 0,
        }
    }

    addCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in city " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.createCity(
                            value,
                            this.state.name,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lon,
                            (err) => {
                                console.log("err in createCity " + err)
                            },
                            (name, country, email, description, lat, lon, id) => {

                                console.log("name, country, email, description, lat, lon id is " +
                                    name, country, email, description, lat, lon, id)
                                console.log("successfully created a city!!!")
                            })
                    }
                })
            })
        })
    }

    getCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in getCity() " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.getCity(
                            value,
                            this.state.name,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lon,

                            (err) => {
                                console.log("err in getCity " + err)
                            },
                            (name, country, email, description, lat, lon) => {
                                this.setState({name: name})
                                this.setState({country: country})
                               // this.setState({email: email})
                                this.setState({description: description})
                                this.setState({lat: lat})
                                this.setState({lon: lon})
                                console.log("name, country, email, description, lat, lon  is " + name, country, email,
                                    description, lat, lon)
                                console.log("successfully got a city!!!")
                                this.props.navigation.navigate('DisplayCity')
                            })
                    }
                })
            })
        })
    }

    render() {
        return (
            // <View style={{flex: 1}}>

                <View style={{flex: 2, justifyContent: 'center', alignItems: 'center'}}>

                    <TextInput
                        style={styles.inputs}
                        placeholder="Name"
                        onChangeText={(name) => this.setState({name})}/>

                    <TextInput
                        style={styles.inputs}
                        placeholder="Country"
                        onChangeText={(country) => this.setState({country})}/>
                    <TextInput
                        style={styles.inputs}
                        placeholder="Description"
                        onChangeText={(description) => this.setState({description})}/>
                    {/*<TextInput*/}
                    {/*    style={styles.inputs}*/}
                    {/*    keyboardType = 'numeric'*/}
                    {/*    placeholder="Lat"*/}
                    {/*    onChangeText={(lat) => this.setState({lat})}/>*/}
                    {/*<TextInput*/}
                    {/*    style={styles.inputs}*/}
                    {/*    keyboardType = 'numeric'*/}
                    {/*    placeholder="Lon"*/}
                    {/*    onChangeText={(lon) => this.setState({lon})}/>*/}
                    <Button title="Add a new city"
                            onPress={() => this.addCity()}/>
                </View>
            // </View>
        )

    }
}
