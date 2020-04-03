import React, {Component} from 'react';
import {Button, Image, NativeModules, TextInput, View} from 'react-native';
import CustomHeader from '../headers/CustomHeader'
import AsyncStorage from "@react-native-community/async-storage";
import Style from "../../styles/Style";
import Card from "react-native-material-cards/Card";
import {Body, CardItem, Text} from "native-base";

class Settings extends Component {
    constructor(props) {
        super(props);
        this.state = {
            userId: 0,
            avatar_url: '',
            name: '',
            description: '',
        }
    }

    onClickListener() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let key = store[i][0];
                    let value = store[i][1]
                    console.log("key/value in settings " + key + " " + value)

                    if (value !== null) {
                        NativeModules.ProfilesModule.updateUser(
                            value,
                            key,
                            key,
                            this.state.name,
                            this.state.description,
                            this.state.userId,
                            (err) => {
                                console.log("error In settings " + err)
                            },
                            (email, name, description, userId) => {
                                this.setState({name: name})
                                this.setState({description: description})
                                this.userId({description: userId})
                                console.log("successful!!!" + this.state.name, this.state.description)
                            })
                    }
                })
            })
        })
    }

    render() {
        return (
            <View style={{flex: 1}}>
                <CustomHeader title="Settings" navigation={this.props.navigation}/>
                <View style={Style.view}>
                    <Card style={Style.cardContainer}>
                        <CardItem>
                            <TextInput
                                style={Style.text}
                                placeholder="Name"
                                onChangeText={(name) => this.setState({name})}/>
                        </CardItem>

                        <CardItem>
                            <TextInput
                                style={Style.text}
                                placeholder="Description"
                                onChangeText={(description) => this.setState({description})}/>
                        </CardItem>
                        <CardItem>
                            <Button style={Style.btnPressStyle} title="Save changes"
                                    onPress={() => this.onClickListener()}/>
                        </CardItem>
                    </Card>
                </View>
            </View>
        )
    }
}

export default Settings
