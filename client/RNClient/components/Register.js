import React, { Component } from 'react';
import {
    Text,
    View,
    TextInput,
    TouchableHighlight,
    Image,
    Alert
} from 'react-native';

import { createAppContainer } from "react-navigation";
import { createStackNavigator} from "react-navigation-stack";
import styles from './Style'

export default class Register extends Component {
    constructor(props) {
        super(props);
        state = {
            name:"",
            email: '',
            password: '',
            cpassword: '',
        }
    }

    onClickListener = (viewId) => {
        Alert.alert("Alert", "Button pressed " + viewId);
    }

    render() {
        return (
            <View style={styles.container}>
                <View style={styles.inputContainer}>
                    <Image style={styles.inputIcon}  source={require('../img/mail.png')} />
                    <TextInput style={styles.inputs}
                        placeholder="Name"
                        keyboardType="ascii-capable"
                        underlineColorAndroid='transparent'
                        onChangeText={(name) => this.setState({ name })} />
                </View>
                <View style={styles.inputContainer}>
                    <Image style={styles.inputIcon}  source={require('../img/mail.png')} />
                    <TextInput style={styles.inputs}
                        placeholder="Email"
                        keyboardType="email-address"
                        underlineColorAndroid='transparent'
                        onChangeText={(email) => this.setState({ email })} />
                </View>

                <View style={styles.inputContainer}>
                    <Image style={styles.inputIcon} source={require('../img/key.png')} />
                    <TextInput style={styles.inputs}
                        placeholder="Password"
                        secureTextEntry={true}
                        underlineColorAndroid='transparent'
                        onChangeText={(password) => this.setState({ password })} />
                </View>
                <View style={styles.inputContainer}>
                    <Image style={styles.inputIcon} source={require('../img/key.png')} />
                    <TextInput style={styles.inputs}
                        placeholder="Confirm Password"
                        secureTextEntry={true}
                        underlineColorAndroid='transparent'
                        onChangeText={(cpassword) => this.setState({ cpassword })} />
                </View>

                <TouchableHighlight style={[styles.buttonContainer, styles.loginButton]} onPress={() => this.onClickListener('register')}>
                    <Text style={styles.loginText}>Create Account</Text>
                </TouchableHighlight>

            </View>
        );
    }
}
