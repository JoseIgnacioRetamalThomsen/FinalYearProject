import React, {Component} from 'react';
import {View, Image, SafeAreaView, ScrollView, NativeModules} from 'react-native';
import {Text, List, ListItem} from 'native-base';
import {IMAGE} from '../constants/Image'
import GeoLoc from "./GeoLoc";
import AsyncStorage from "@react-native-community/async-storage";
import {logger} from 'react-native-logger'

class SideMenu extends Component {
    async logout() {
        try {
            AsyncStorage.getAllKeys((err, keys) => {
                AsyncStorage.multiGet(keys, (err, stores) => {
                    stores.map((result, i, store) => {
                        let key = store[i][0];
                        let value = store[i][1]
                        NativeModules.LoginModule.logout(
                            value,
                            key,
                            async (err) => {
                                this.props.navigation.navigate('auth')
                                await AsyncStorage.clear()
                                logger.log(err)
                                console.log(key, value)
                            },
                            async (isSuccess) => {
                                isSuccess ? this.props.navigation.navigate('auth') : this.props.navigation.navigate('app')
                                await AsyncStorage.clear()
                            })
                    })
                })
            })
        } catch (e) {
            logger.log(e.message)
        }
        this.props.navigation.navigate('auth')
    }

    render() {
        return (
            <SafeAreaView style={{flex: 1}}>
                <View style={{height: 150, alignItems: 'center', justifyContent: 'center'}}>
                    <Image source={IMAGE.ICON_DEFAULT_PROFILE} style={{height: 120, width: 120, borderRadius: 60}}/>
                    {/*<GeoLoc/>*/}
                </View>
                <ScrollView>
                    <List>
                        <ListItem onPress={() => this.props.navigation.navigate('Profile')}>
                            <Text>
                                Profile
                            </Text>
                        </ListItem>
                        <ListItem onPress={() => this.props.navigation.navigate('Settings')}>
                            <Text>
                                Settings
                            </Text>
                        </ListItem>
                        <ListItem noBorder>
                            <Text onPress={() => this.logout()}>
                                Log out
                            </Text>
                        </ListItem>
                    </List>
                </ScrollView>
            </SafeAreaView>
        )
    }
}

export default SideMenu
