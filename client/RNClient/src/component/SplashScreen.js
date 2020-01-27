import React, { Component } from 'react';
import {
    Platform,
    ActivityIndicator,
    StyleSheet,
    Text,
    View,
    StatusBar, default as AsyncStorage
} from 'react-native';

export default class SplashScreen extends Component {
    constructor() {
        super();
        this.bootstrapAsync();
    }
    bootstrapAsync = async () => {
        const userToken = await AsyncStorage.getItem('userToken');
        // This will switch to the App screen or Auth screen and this loading
        // screen will be unmounted and thrown away.
        this.props.navigation.navigate(userToken ? 'app' : 'auth');
    };
    render() {
        return (
            <View style={styles.container}>
                <ActivityIndicator />
                <StatusBar barStyle="default" />
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#4F6D7A',
    }
});
