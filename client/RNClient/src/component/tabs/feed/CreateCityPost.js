import React, {Component} from "react";
import {Button, Text, View, Card, TextInput, TouchableHighlight, NativeModules, ScrollView} from "react-native";
import CustomHeader from "../../CustomHeader";
import {Root} from "native-base";

import styles from '../../../styles/Style'
import AsyncStorage from '@react-native-community/async-storage'
import LoadImage from '../../LoadImage'

export default class CreateCityPost extends Component {
    constructor(props) {
        super(props)
        this.state = {
            indexId: 0,
            creatorEmail: '',
            city: '',
            country: '',
            title: '',
            body: '',
            timeStamp: '',
            //new Date().getDate()
            likes: [],
            mongoId: '',
            img: '',
            lat: 0,
            lon: 0,
        }
    }

    // componentDidMount() {
    //     const indexId = this.props.navigation.getParam('indexId', '')
    //     const city = this.props.navigation.getParam('city', '')
    //     const country = this.props.navigation.getParam('country', '')
    //
    //     this.setState({
    //         indexId,
    //         city,
    //         country,
    //     })
    //     console.log('componentDidMount' + indexId, city)
    // }

    getCityPosts() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in getCity() " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.getCity(
                            value,
                            this.state.city,
                            this.state.country,
                            key,
                            this.state.description,
                            this.state.lat,
                            this.state.lon,

                            (err) => {
                                console.log("err in getCity " + err)
                            },
                            (valid, name, country, email, description, lat, lon, id) => {
                                this.setState({name: name})
                                this.setState({country: country})
                                // this.setState({email: email})
                                this.setState({description: description})
                                this.setState({lat: lat})
                                this.setState({lon: lon})
                                console.log("name id  is " + name, id)
                                console.log("successfully got a city!!!")
                                //this.props.navigation.navigate('DisplayCities')
                            })
                    }
                })
            })
        })
    }

    createCityPost() {
        console.log("indexId " + this.state.indexId)
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]

                    console.log("key/value in city " + key + " " + value)

                    if (value !== null) {
                        NativeModules.PostModule.createCityPost(
                            this.state.indexId,
                            key,
                            this.state.city,
                            this.state.country,
                            this.state.title,
                            this.state.body,
                            this.state.timeStamp,
                            (err) => {
                                console.log("err in createCityPost " + err)
                            },
                            (indexId) => {
                                this.setState({indexId: indexId})
                                console.log("indexId in createCityPost is " + indexId, this.state.indexId)
                                this.props.navigation.navigate('DisplayCityPosts')
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

                        <ScrollView style={{flex: 1}}>

                            <View style={styles.inputContainer}>
                                <TextInput
                                    style={styles.inputs}
                                    placeholder="City"
                                    underlineColorAndroid='transparent'
                                    onChangeText={(city) => this.setState({city})}/>


                            </View>

                            <View style={styles.inputContainer}>
                                <TextInput
                                    style={styles.inputs}
                                    placeholder="Country"
                                    onChangeText={(country) => this.setState({country})}/>


                            </View>


                            <View style={styles.inputContainer}>
                                <TextInput
                                    style={styles.inputs}
                                    placeholder="Description"
                                    onChangeText={(body) => this.setState({body})}/>
                            </View>


                            <TouchableHighlight style={[styles.buttonContainer, styles.loginButton]}
                                                onPress={() => this.createCityPost()}>
                                <Text style={styles.loginText}>Submit</Text>
                            </TouchableHighlight>

                        </ScrollView>
                    </View>

                </View>
            </Root>
        )
    }
}
