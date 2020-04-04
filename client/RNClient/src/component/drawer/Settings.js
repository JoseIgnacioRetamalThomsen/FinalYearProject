import React, {Component} from 'react';
import {Button, Image, NativeModules, TextInput, TouchableOpacity, View} from 'react-native';
import CustomHeader from '../headers/CustomHeader'
import AsyncStorage from "@react-native-community/async-storage";
import Style from "../../styles/Style";
import Card from "react-native-material-cards/Card";
import {Body, CardItem, Text} from "native-base";

class Settings extends Component {
    constructor(props) {
        super(props);
        this.state = {
            userId: -999,
            avatar_url: '',
            name: '',
            description: '',
        }
    }
    componentDidMount() {
        const userId = this.props.navigation.getParam('userId', '')
        this.setState({
            userId,
        })
    }
    updateUser() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.ProfilesModule.updateUser(
                            token,
                            email,
                            email,
                            this.state.name,
                            this.state.description,
                            this.state.userId,
                            (err) => {
                                console.log(err)
                            },
                            (email, name, description, userId) => {
                                this.setState({name: name})
                                this.setState({description: description})
                                //this.userId({userId: userId})
                                console.log("successful!!!" + this.state.name, this.state.description)
                                this.props.navigation.navigate("Profile")
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
                {/*<View style={Style.view}>*/}
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
                            <TouchableOpacity style={Style.btnPressStyle}
                                              onPress={() => this.updateUser()}>
                                <Text style={Style.txtStyle}>Save changes</Text>
                            </TouchableOpacity>
                        </CardItem>
                    </Card>
                {/*</View>*/}
            </View>
        )
    }
}

export default Settings
