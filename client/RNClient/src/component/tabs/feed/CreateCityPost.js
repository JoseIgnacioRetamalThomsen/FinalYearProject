import React, {Component} from "react";
import {Button, Text, View, Card, TextInput, TouchableHighlight, NativeModules, ScrollView, Image} from "react-native";
import CustomHeader from "../../CustomHeader";
import {Root} from "native-base";

import styles from '../../../styles/Style'
import AsyncStorage from '@react-native-community/async-storage'
import LoadImage from '../../LoadImage'
import PhotoUpload from "react-native-photo-upload";

export default class CreateCityPost extends Component {
    constructor(props) {
        super(props)
        this.state = {
            cityId: 10,
            postId: '',
            city: '',
            country: '',
            title: '',
            body: '',
            image: '',

           // creatorEmail: '',

            //timeStamp: '',
            // //new Date().getDate()
            // likes: [],
            // mongoId: '',
            // img: '',
            // lat: 0,
            // lon: 0,
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

    createCityPost() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let value = store[i][1]

                    if (value != null) {
                        NativeModules.PostModule.createCityPost(
                            this.state.cityId,
                            email,
                            this.state.city,
                            this.state.country,
                            this.state.title,
                            this.state.body,
                            (err) => {
                                console.log(err)
                            },
                            (postId) => {
                                this.setState({postId: postId})
                                console.log("h in createCityPost is " + postId)
                               // this.props.navigation.navigate('DisplayCityPosts')
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
                                    placeholder="Title"
                                    underlineColorAndroid='transparent'
                                    onChangeText={(title) => this.setState({title})}/>

                            </View>

                            <View style={styles.inputContainer}>
                                <TextInput
                                    style={styles.inputs}
                                    placeholder="Description"
                                    onChangeText={(body) => this.setState({body})}/>
                            </View>


                            <PhotoUpload onPhotoSelect={image => {
                                if (image) {
                                    this.setState({image: image})
                                }
                            }
                            }>
                                <Image source={{image: this.state.image}}
                                       style={{
                                           height: 120,
                                           width: 120,
                                           borderRadius: 60,
                                           borderColor: 'black',
                                           borderWidth: 5,
                                           flex: 0,
                                           resizeMode: 'cover'
                                       }}/>
                            </PhotoUpload>


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
