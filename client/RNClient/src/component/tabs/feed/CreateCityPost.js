import React from 'react'
import {Button, TextInput, View, NativeModules, Image} from "react-native";
import styles from "../../../styles/Style";
import AsyncStorage from "@react-native-community/async-storage";
import { Card, CardTitle, CardContent, CardAction, CardButton, CardImage } from 'react-native-material-cards'
import {Body, CardItem, Text, Title} from "native-base";

export default class CreateCityPost extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            indexId: 0,
            cityName: '',
            cityCountry: '',
            title: '',
            body: '',
            timeStamp: '',
            likes: [],
            mongoId: '',
        }
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
                            this.state.cityName,
                            this.state.cityCountry,
                            this.state.title,
                            this.state.body,
                            this.state.timeStamp,
                            this.state.likes,
                            this.state.mongoId,
                            (err) => {
                                console.log("err in createPlace " + err)
                            },
                            (indexId) => {
                                this.setState({indexId: indexId})

                                console.log(" indexId in createCityPost is " + indexId)
                            })
                    }
                })
            })
        })
    }

    render() {
        return (
            <View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>

                <Card>
                    <CardImage
                        source={{uri: 'http://placehold.it/480x270'}}
                        title="Above all i am here"
                    />
                    <CardTitle
                        title="This is a title"
                        subtitle="This is subtitle"
                    />
                    <CardContent text="Your device will reboot in few seconds once successful, be patient meanwhile" />
                    <CardAction
                        separator={true}
                        inColumn={false}>
                        <CardButton
                            onPress={() => {}}
                            title="Push"
                            color="blue"
                        />
                        <CardButton
                            onPress={() => {}}
                            title="Later"
                            color="blue"
                        />
                    </CardAction>
                </Card>

            </View>

        )
    }
}
