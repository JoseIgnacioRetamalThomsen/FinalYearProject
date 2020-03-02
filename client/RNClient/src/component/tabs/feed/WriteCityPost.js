import React, {Component} from "react";
import {Button, Text, View, Card, TextInput, TouchableHighlight, NativeModules } from "react-native";
import CustomHeader from "../../CustomHeader";
import {Root} from "native-base";
import GeoLoc from "../../GeoLoc";
import styles from '../../../styles/Style'
import AsyncStorage from '@react-native-community/async-storage'
import LoadImage from '../../LoadImage'

export default class WriteCityPost extends Component {
    constructor(props) {
            super(props)
            this.state = {
                indexId: Date.now(),
                creatorEmail:'',
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
callbackFunction = ( city, country, lat, lon) => {
        this.setState({city: city})
        this.setState({country: country})
         this.setState({lat: lat})
          this.setState({lon: lon})
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
                                this.state.city,
                                this.state.country,
                                key,
                                this.state.body,
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
        addCityPost() {
         console.log("indexId "+this.state.indexId)
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
                                this.state.likes,
                                this.state.mongoId,
                                (err) => {
                                    console.log("err in createCityPost " + err)
                                },
                                (isValid, index) => {
                                    if(index != -1){
                                   // this.setState({indexId: indexId})
                                    console.log(" indexId in createCityPost is " + index)
                                    this.props.navigation.navigate('DisplayCityPosts')
                                    }else{
                                    console.log(" smth went wrong " + indexId)
                                    }
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
                    onChangeText={(body) => this.setState({body})}/>
                    <View style={styles.container}>

                    </View>
                     </View>

                    </View>

                    <View style={styles.container}>
                       <TouchableHighlight style={[styles.buttonContainer, styles.loginButton]}
                           onPress={() => this.addCityPost()}>
                         <Text style={styles.loginText}>Submit</Text>
                     </TouchableHighlight>
                      </View>

                      </View>
            </Root>
        )
    }
}
