import React, { Component } from 'react';
import { Header, Left, Body, Right, Button, Icon, Title } from 'native-base';
import {NativeModules, Text} from "react-native";
import AsyncStorage from "@react-native-community/async-storage";

class CustomHeader extends Component {
    render() {
        let { title, isHome } = this.props;
        return (
            <Header style ={{backgroundColor:'#007AFF'}}>
                <Left>
                    {
                        isHome ?
                            <Button transparent onPress={() => this.props.navigation.openDrawer()}>
                                <Icon name='menu' />
                            </Button> :
                            <Button transparent onPress={() => this.props.navigation.goBack()}>
                                <Icon name='arrow-back' />
                            </Button>
                    }
                </Left>
                <Body>
                    <Title>{title}</Title>
                </Body>
                <Right>
                    <Button hasText transparent onPress= {() => this.visitCity()}>
                    </Button>
                </Right>
            </Header>
        );
    }

    visitCity() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.ProfilesModule.visitCity(
                            token,
                            email,
                            this.state.cityId,
                            (err) => {
                                console.log(err)
                            },

                            (timestamp) => {
                               // this.setState({timestamp: timestamp})
                                console.log("timestamp is ", timestamp)

                            })
                    }
                })
            })
        })
    }
}
export default CustomHeader
