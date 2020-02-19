import React, {Component} from 'react';
import {Button, NativeModules, TextInput, View} from 'react-native';
import AsyncStorage from "@react-native-community/async-storage";
import {Text} from "native-base";

export default class DisplayPlace extends Component {
    constructor(props) {
        super(props);
        this.state = {
            avatar_url: '',
            name: '',
            email:'',
            city: '',
            country: '',
            description: '',
            lat: 0,
            lon: 0,
        }
    }

    componentDidMount() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in displayPLace " + key + " " + value)

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
                                console.log("error In getPlace " + err)
                            },
                            (name, city, country, email, description) => {
                                this.setState({name: name})
                                this.setState({city: city})
                                this.setState({country: country})
                                this.setState({email: email})
                                this.setState({description: description})
                                console.log("successful in getPlace values " + this.state.name, this.state.description)
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

                    <View style={{flex: 1}}>
                        <View>
                            <Text>Name {this.state.name} </Text>
                        </View>
                        <View>
                            <Text>City {this.state.city} </Text>
                        </View>
                        <View>
                            <Text>Country {this.state.country} </Text>
                        </View>
                        <View>
                            <Text>Description {this.state.description} </Text>
                        </View>
                    </View>
                </View>
            </View>
        );
    }
}
