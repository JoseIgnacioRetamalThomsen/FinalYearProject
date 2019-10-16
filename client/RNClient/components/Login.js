import React, { Component } from 'react';
import {
    Text,
    View,
    TextInput,
    TouchableHighlight,
    Image,
    Alert
} from 'react-native';
import LoginGrpc from "./LoginGrpc"
//import { LoginButton } from 'react-native-fbsdk';
import styles from './Style'

export default class Login extends Component {

    constructor(props) {
        super(props);
        state = {
            email: '',
            password: '',
        }
    }

    onClickListener = (viewId) => {
        Alert.alert("Alert", "Button pressed " + viewId);
        LoginGrpc.check("one","two");
    }

    render() {
        return (
            <View style={styles.container}>
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

                <TouchableHighlight style={[styles.buttonContainer, styles.loginButton]} onPress={() => this.onClickListener('login')}>
                    <Text style={styles.loginText}>Login</Text>
                </TouchableHighlight>

                <TouchableHighlight style={styles.buttonContainer} onPress={() => this.onClickListener('restore_password')}>
                    <Text>Forgot your password?</Text>
                </TouchableHighlight>

                <TouchableHighlight style={styles.buttonContainer} onPress={() => this.onClickListener('register')}>
                    <Text>Register</Text>
                </TouchableHighlight>
{/* 
                <View>
                    <LoginButton
                        publishPermissions={["email"]}
                        onLoginFinished={
                            (error, result) => {
                                if (error) {
                                    alert("Login failed with error: " + error.message);
                                } else if (result.isCancelled) {
                                    alert("Login was cancelled");
                                } else {
                                    alert("Login was successful with permissions: " + result.grantedPermissions)
                                }
                            }
                        }
                        onLogoutFinished={() => alert("User logged out")} />
                </View> */}
            </View>
        );
    }
}
