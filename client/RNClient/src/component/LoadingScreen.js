import React, {Component, Fragment, useEffect} from 'react';
import {
    AsyncStorage,
    SafeAreaView,
    Platform,
    ActivityIndicator,
    StyleSheet,
    Text,
    View,
    StatusBar, TouchableHighlight,
} from 'react-native';
import SplashScreen from 'react-native-splash-screen'

import Login from "./auth/Login";

export const LoadingScreen = () => {
    useEffect(() => {
        SplashScreen.hide()
    }, [])

    // constructor() {
    //     super();
    //     this.bootstrapAsync();
    // }

    // const bootstrapAsync = async (token) => {
    //      // const userToken = await AsyncStorage.getItem('userToken', token);
    //      // this.props.navigation.navigate(userToken ? 'app' : 'auth');
    //     this.props.navigation.navigate('auth')
    //  }

    // render() {
    return (
        <Login></Login>
    )
}


