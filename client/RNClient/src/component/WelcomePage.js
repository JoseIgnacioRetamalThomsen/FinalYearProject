import React, {Component} from 'react';
import SplashScreen from 'react-native-splash-screen'
import {
    NativeModules,
    View,
    Text,
} from 'react-native';
import AsyncStorage from '@react-native-community/async-storage'
import {logger} from 'react-native-logger'

export default class WelcomePage extends Component {
    componentDidMount() {
        this.getSavedToken().then(r => logger.log(r))
        SplashScreen.hide();
    }
    getSavedToken = async () => {
        try {
            AsyncStorage.getAllKeys((err, keys) => {
                AsyncStorage.multiGet(keys, (err, stores) => {
                    stores.map((result, i, store) => {
                        let key = store[i][0];
                        let value = store[i][1]
                        console.log("key/value in welcome page  " + key + " " + value)
                        if (value !== null) {
                            NativeModules.LoginModule.checkToken(
                                value,
                                key,
                                (err) => {
                                    logger.log(err)
                                    console.log("this.token, this.email " + key, value)
                                    this.props.navigation.navigate('auth')
                                },
                                (isSuccess) => {
                                    isSuccess ? this.props.navigation.navigate('app') : this.props.navigation.navigate('auth')
                                    logger.log(isSuccess)
                                }
                             )
                        } else {
                            this.props.navigation.navigate("auth")
                        }
                    });
                });
            });
        } catch (error) {
            logger.log(error)
        }
        this.props.navigation.navigate("auth")
    }

    render() {
        return (
            <View>
                <Text> Welcome Page</Text>
            </View>
        )
    }
}
