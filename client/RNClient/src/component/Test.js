import React from 'react'
import {Button, TextInput, View, NativeModules} from "react-native";
import styles from "../styles/Style";
import AsyncStorage from "@react-native-community/async-storage";
import TestLocationReducer from "./TestLocationReducer";
import MapInput from "./MapInput";

export default class Test extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            email: '',
            name: '',
            city: '',
            country: '',
            description:'',
            lat: 0,
            lng: 0,

        }
    }

    visitPlace() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in city " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.visitPlace(
                            value,
                            key,
                            this.state.name,
                            this.state.city,
                            this.state.country,

                            (err) => {
                                console.log("err in visitPlace " + err)
                            },

                            (isValid, email, name, city, country) => {
                                this.setState({name: name})
                                this.setState({city: city})
                                this.setState({country: country})

                                console.log("name, city, country in visitPlace is " + name, city, country)
                                console.log("successfully visited a place!!!")
                            })
                    }
                })
            })
        })
    }

    visitCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in visitCity " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.visitCity(
                            value,
                            key,
                            this.state.name,
                            this.state.country,

                            (err) => {
                                console.log("err in visitCity " + err)
                            },
                            (isValid, email, name, country) => {
                                this.setState({name: name})
                                this.setState({country: country})

                                console.log("name, country in visitCity is " + name, country)
                                console.log("successfully visited a city!!!")
                            })
                    }
                })
            })
        })
    }

    getVisitedCities() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in getVisitedCities " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.getVisitedCities(
                            value,
                            key,

                            (err) => {
                                console.log("err in getVisitedCities " + err)
                            },
                            (isValid, email, valid, name, country, creatorEmail, description, lng, lat, id) => {
                                this.setState({name: name})
                                this.setState({country: country})

                                console.log("name, country in getVisitedCities is " + name, country)
                                console.log("successfully getVisitedCities!!!")
                            })
                    }
                })
            })
        })
    }

    getVisitedPlaces() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in getVisitedCities " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.getVisitedPlaces(
                            value,
                            key,
                            (err) => {
                                console.log("err in getVisitedPlaces " + err)
                            },

                            (isValid, name, city, country, email, description, lat, lng, id) => {
                                this.setState({name: name})
                                this.setState({country: country})

                                console.log("name, country in getVisitedPlaces is " + name, country)
                                console.log("successfully getVisitedPlaces!!!")
                            })
                    }
                })
            })
        })
    }

    getCityPlaces() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in getVisitedCities " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.getCityPlaces(
                            value,
                            this.state.name,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lng,

                            key,
                            (err) => {
                                console.log("err in getCityPlaces " + err)
                            },
                            (isValid, name, city, country, email, description, lat, lng, id) => {
                                this.setState({name: name})
                                this.setState({country: country})

                                console.log("name, country in getCityPlaces is " + name, country)
                                console.log("successfully getCityPlaces!!!")
                            })
                    }
                })
            })
        })
    }

    render() {
        return (
            <View>
                {/*<Button title="Visit this place"*/}
                {/*        onPress={() => this.visitPlace()}/>*/}

                {/*<Button title="Visit this city"*/}
                {/*        onPress={() => this.visitCity()}/>*/}

                {/*<Button title="getVisitedCities"*/}
                {/*        onPress={() => this.getVisitedCities()}/>*/}

                {/*<Button title="getVisitedPlaces"*/}
                {/*        onPress={() => this.getVisitedPlaces()}/>*/}

                {/*<Button title="getCityPlaces"*/}
                {/*        onPress={() => this.getCityPlaces()}/>*/}
                <TestLocationReducer/>
            </View>
        )
    }
}
