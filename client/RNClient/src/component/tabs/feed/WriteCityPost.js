import React, {Component} from "react";
import {Button, Text, View, Card, TextInput, TouchableHighlight, NativeModules } from "react-native";
import CustomHeader from "../../CustomHeader";
import {Root} from "native-base";
import GeoLoc from "../../GeoLoc";
import { Rating, AirbnbRating } from 'react-native-ratings';
import styles from '../../../styles/Style'
import AsyncStorage from '@react-native-community/async-storage'
import LoadImage from '../../LoadImage'

export default class WriteCityPost extends Component {
    constructor(props) {
            super(props)
            this.state = {
                indexId: 777,
                creatorEmail:'',
                city: '',
                country: '',
                title: '',
                body: '',
                timeStamp: '',
                //new Date().getDate()
                likes: [],
                mongoId: '',
                img: ''
            }
        }
callbackFunction = ( city, country) => {
        this.setState({city: city})
        this.setState({country: country})
    }
        addCityPost() {

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
                                (isValid, indexId) => {
                                    if(indexId != -1){
                                    this.setState({indexId: indexId})
                                    console.log(" indexId in createCityPost is " + indexId)
                                    this.props.navigation.navigate('CityPosts')
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
                      placeholder="City"
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
