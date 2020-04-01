import React, {Component} from 'react';
import {Animated, View, Image, SafeAreaView, ScrollView, NativeModules} from 'react-native';
import {Text, List, ListItem, Root, ActionSheet} from 'native-base';
import {IMAGE} from '../constants/Image'
import GeoLoc from "./GeoLoc";
import AsyncStorage from '@react-native-community/async-storage'
import PhotoUpload from 'react-native-photo-upload'

class SideMenu extends Component {
    constructor(props) {
        super(props);
    }
    componentDidMount() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.PhotosModule.getProfilePhoto(
                            email,
                            token,
                            (err) => {
                                console.log(err)
                            },
                            (url) => {
                                    this.setState({avatar_url: url})
                                    console.log("url", url)

                            })
                    }
                })
            })
        })
    }

     logout() {
        try {
            AsyncStorage.getAllKeys((err, keys) => {
                AsyncStorage.multiGet(keys, (err, stores) => {
                    stores.map((result, i, store) => {
                        let key = store[i][0];
                        let value = store[i][1]
                        NativeModules.LoginModule.logout(
                            value,
                            key,
                            (err) => {
                                this.props.navigation.navigate('auth')
                                console.log("Deleted email & token ", key, value)
                                 AsyncStorage.clear()

                            },
                             (isSuccess) => {
                                isSuccess ? this.props.navigation.navigate('auth') : this.props.navigation.navigate('app')
                                 AsyncStorage.clear()
                            })
                    })
                })
            })
        } catch (e) {
            console.log(e.message)
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
                            // console.log('Image base64 string: ', avatar)
                        }
                    }}>
                       <Image source={IMAGE.ICON_DEFAULT_PROFILE} style={{height: 120, width: 120, borderRadius: 60,  resizeMode:'cover'}}/>
                        </PhotoUpload>
                    <GeoLoc/>
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
