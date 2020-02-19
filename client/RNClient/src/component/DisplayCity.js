import React, {Component} from 'react';
import {Button, NativeModules, View} from 'react-native';
import AsyncStorage from "@react-native-community/async-storage";
import {Text} from "native-base";

export default class DisplayCity extends Component {
    constructor(props) {
        super(props);
        this.state = {
            name: '',
            country: '',
            email: '',
            description: '',
            lat: 0,
            lon: 0,
            id: 0,
        }
    }

    componentDidMount() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in displayCity " + key + " " + value)

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
                                console.log("error In ProfilesModule.getCity " + err)
                            },
                            (name, country, email, description, lat, lon, id) => {
                                this.setState({name: name})
                                this.setState({country: country})
                                this.setState({email: email})
                                this.setState({description: description})
                                this.setState({lat: lat})
                                this.setState({lon: lon})
                                this.setState({id: id})
                                console.log("successful values in getCity!!!" + this.state.name, this.state.country, this.state.description)
                            })

                    }
                })
            })
        })
    }
    updateCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in updatecity " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.updateCity(
                            value,
                            this.state.name,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lon,
                            (err) => {
                                console.log("err in updateCity " + err)
                            },
                            (name, country, email, description, lat, lon, id) => {
                                this.setState({name: name})
                                this.setState({country: country})
                                // this.setState({email: email})
                                this.setState({description: description})
                                this.setState({lat: lat})
                                this.setState({lon: lon})
                                console.log("name, country, email, description, lat, lon id in updateCity is " + name, country, email, description, lat, lon, id)
                                console.log("successfully updated a city!!!")
                            })
                    }
                })
            })
        })
    }
    render() {

        return (

            <View style={{flex: 1}}>
                <View>
                    <Text>Name {this.state.name} </Text>
                </View>
                <View>
                    <Text>Country {this.state.country} </Text>
                </View>
                <View>
                    <Text>email {this.state.email} </Text>
                </View>
                <View>
                    <Text>Description {this.state.description} </Text>
                </View>
                <View>
                    <Text>lat {this.state.lat} </Text>
                </View>
                <View>
                    <Text>lon {this.state.lon} </Text>
                </View>
                <View>
                    <Text>id {this.state.id} </Text>
                </View>
                <Button title="Update an existing city"
                        onPress={() => this.updateCity()}/>
            </View>
        );
    }
}
