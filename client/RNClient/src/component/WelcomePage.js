import React, {Component} from 'react';
import SplashScreen from 'react-native-splash-screen'
import {
    NativeModules,
    View,
    Text,
} from 'react-native';
import AsyncStorage from '@react-native-community/async-storage'

export default class WelcomePage extends Component{
    constructor(props) {
        super(props);
        this.state = {
            isLoaded: false,
            email: '',
            token: ''
        }
        //this.compare = this.compare.bind(this);
    }

    componentDidMount() {
        // do stuff while splash screen is shown
        // After having done stuff (such as async tasks) hide the splash screen
        //SplashScreen.hide();
        // this.setState({isLoaded: true}, () => { SplashScreen.hide(); })
        setTimeout(() => {
            this.getSavedToken().then(r => alert(r));
            SplashScreen.hide();
           // this.props.navigation.replace("app");
        }, 1000);
    }

    // async compare(savedToken) {
    //     NativeModules.LoginModule.checkToken(
    //         this.state.email,
    //         this.state.token,
    //         (err) => {
    //             alert(err)
    //         },
    //         (isSuccess) => {
    //             isSuccess ? this.props.navigation.navigate('app') : this.props.navigation.navigate('auth')
    //         }
    //     )
    // }

    getSavedToken = async () => {
        try {
            const value = await AsyncStorage.getItem('isUser')
            const email = await AsyncStorage.getItem('email')

            if (value !== null) {
                // We have data!!
                console.log(value);
                if (value == "true") {
                    this.props.navigation.navigate("app", {email: email});
                } else {
                    this.props.navigation.navigate("auth");
                }
            } else {
                this.props.navigation.navigate("auth");
            }
        } catch (error) {
            // Error retrieving data
        }
    }

    render() {
        return (
            <View>
                <Text> Welcome Page</Text>
            </View>
        )
    }
}
