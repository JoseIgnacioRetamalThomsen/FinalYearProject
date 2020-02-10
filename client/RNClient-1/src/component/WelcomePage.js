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

    // async compare() {
    //     NativeModules.LoginModule.checkToken(
    //         this.state.token,
    //         this.state.email,
    //         (err) => {
    //             logger.log(err)
    //             console.log("this.token, this.email " + this.state.token, this.state.email)
    //             this.props.navigation.navigate('auth')
    //         },
    //         (isSuccess) => {
    //             isSuccess ? this.props.navigation.navigate('app') : this.props.navigation.navigate('auth')
    //             logger.log(isSuccess)
    //         }
    //     )
    // }

    getSavedToken = async () => {
        try {
            AsyncStorage.getAllKeys((err, keys) => {
                AsyncStorage.multiGet(keys, (err, stores) => {
                    stores.map((result, i, store) => {
                        // get at each store's key/value so you can work with it
                        let key = store[i][0];
                        // this.setState({email: key}, () => {
                        //     console.log(this.state.email)
                        // })
                        let value = store[i][1]
                        // this.setState({token: value})
                        console.log("key/value " + key + value)
                        if (value !== null) {
                            //this.compare()
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
        } catch (error) {//catch StatusRunTimeException here?
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
