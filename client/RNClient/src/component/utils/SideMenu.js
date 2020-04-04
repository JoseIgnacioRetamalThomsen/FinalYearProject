import React, {Component} from 'react'
import {Animated, View, Image, SafeAreaView, ScrollView, NativeModules} from 'react-native'
import {Text, List, ListItem, Root, CardItem} from 'native-base'
import {IMAGE} from '../../constants/Image'
import GeoLoc from "./GeoLoc"
import AsyncStorage from '@react-native-community/async-storage'
import PhotoUpload from 'react-native-photo-upload'
import Style from '../../styles/Style'

class SideMenu extends Component {
    constructor(props) {
        super(props)
        this.state = {
            avatar_url: '',
            city: '',
            country: '',
        }
    }

    callbackFunction = (lat, lng, city, country) => {
        this.setState({city: city})
        this.setState({country: country})
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

    updatePhoto() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]
                    if (token !== null) {
                        NativeModules.PhotosModule.uploadProfilePhoto(
                            email,
                            token,
                            this.state.avatar_url,

                            (err) => {
                                console.log(err)
                            },
                            (url) => {
                                this.setState({avatar_url: url})
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
                    <Animated.View style={{height: 200, alignItems: 'center', justifyContent: 'center'}}>
                        <PhotoUpload onPhotoSelect={avatar => {
                            if (avatar) {
                                this.setState({avatar_url: avatar})
                                this.updatePhoto()
                            }
                        }}>
                            {this.displayPhoto()}
                        </PhotoUpload>
                        {/*<GeoLoc parentCallback={this.callbackFunction}/>*/}
                        <CardItem>
                            <Text > {this.state.city}, {this.state.country} </Text>
                        </CardItem>
                    </Animated.View>
                    <ScrollView>
                        <List>
                            <ListItem style={{
                                borderColor: '#0080ff', borderBottomColor: "0080ff",
                                borderBottomWidth: 1
                            }} />
                            <ListItem style={{
                                borderColor: '#0080ff', borderBottomColor: "0080ff",
                                borderBottomWidth: 1
                            }} onPress={() => this.props.navigation.navigate('Profile')}>
                                <Text>
                                    Profile
                                </Text>
                            </ListItem>
                            <ListItem style={{
                                borderColor: '#0080ff', borderBottomColor: "0080ff",
                                borderBottomWidth: 1
                            }} onPress={() => this.props.navigation.navigate('Settings')}>
                                <Text>
                                    Settings
                                </Text>
                            </ListItem>

                            <ListItem style={{
                                borderColor: '#0080ff', borderBottomColor: "0080ff",
                                borderBottomWidth: 1
                            }} >
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

    displayPhoto() {
        if (this.state.avatar_url === '') {
            return (<Image source={IMAGE.ICON_DEFAULT_PROFILE} style={Style.profilePhoto}/>)
        } else {
            return (<Image source={{uri: this.state.avatar_url}}
                           style={Style.profilePhoto}/>)
        }
    }
}

export default SideMenu
