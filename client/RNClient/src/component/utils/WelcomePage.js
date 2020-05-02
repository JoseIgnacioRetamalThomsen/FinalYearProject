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
        this.getSavedToken().then(r => logger.log(r)),
        SplashScreen.hide(),
        this.props.navigation.navigate('auth')
    }
    getSavedToken = async () => {
        try {
            AsyncStorage.getAllKeys((err, keys) => {
                AsyncStorage.multiGet(keys, (err, stores) => {
                    stores.map((result, i, store) => {
                        let email = store[i][0];
                        let token = store[i][1]
                        console.log("email/token in welcome page  " + email + " " + token)
                        if (token !== null) {
                            NativeModules.LoginModule.checkToken(
                                token,
                                email,
                                (err) => {
                                    logger.log(err)
                                    SplashScreen.hide();
                                    this.props.navigation.navigate('auth')
                                },
                                (isSuccess) => {
                                    isSuccess ? this.props.navigation.navigate('app') : this.props.navigation.navigate('auth')
                                    SplashScreen.hide();
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
            </View>
        )
    }
}
