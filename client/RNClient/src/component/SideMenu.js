import React, {Component} from 'react';
import {Animated, View, Image, SafeAreaView, ScrollView, NativeModules} from 'react-native';
import {Text, List, ListItem, Root, ActionSheet} from 'native-base';
import {IMAGE} from '../constants/Image'
import GeoLoc from "./GeoLoc";
import AsyncStorage from '@react-native-community/async-storage'
import PhotoUpload from 'react-native-photo-upload'

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
            <Root>
            <SafeAreaView style={{flex: 1}}>
                <Animated.View style={{height: 150, alignItems: 'center', justifyContent: 'center'}}>
                    <PhotoUpload onPhotoSelect={avatar => {
                        if (avatar) {
                            console.log('Image base64 string: ', avatar)
                        }
                    }}>
                       <Image source={IMAGE.ICON_DEFAULT_PROFILE} style={{height: 120, width: 120, borderRadius: 60,  resizeMode:'cover'}}/>
                        </PhotoUpload>
                    {/*<GeoLoc/>*/}
                </Animated.View>
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
            </Root>
        )
    }
}

export default SideMenu
