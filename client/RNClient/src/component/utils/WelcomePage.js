import React, {Component} from 'react';
import SplashScreen from 'react-native-splash-screen'
import {
    NativeModules,
    View,
} from 'react-native';
import AsyncStorage from '@react-native-community/async-storage'


export default class WelcomePage extends Component {
     componentDidMount() {
         this.getSavedToken().then(token => {
            if (token === undefined || token === null || token === "") {
                SplashScreen.hide()
                this.props.navigation.navigate('auth')

            }
            console.log("r", token)
        })


    }
    getSavedToken = async () => {
        try {
            AsyncStorage.getAllKeys((err, keys) => {
                AsyncStorage.multiGet(keys, (err, stores) => {
                    stores.map((result, i, store) => {
                        let email = store[i][0];
                        let token = store[i][1]
                        console.log("email/token in welcome page  " + email + " " + token)
                        if (token !== null || token !== undefined || token !== '') {
                            NativeModules.LoginModule.checkToken(
                                token,
                                email,
                                (err) => {
                                    console.log("e", err)
                                    SplashScreen.hide();
                                    this.props.navigation.navigate('auth')
                                },
                                (isSuccess) => {
                                    SplashScreen.hide()
                                    isSuccess ? this.props.navigation.navigate('app') : this.props.navigation.navigate('auth')

                                }
                             )
                        } else {
                            this.props.navigation.navigate("auth")
                        }
                    });
                });
            });
        } catch (error) {
            console.log("gg",error)
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
