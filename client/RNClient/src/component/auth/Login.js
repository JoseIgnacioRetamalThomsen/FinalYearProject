import React, { Component } from 'react';
import {
    Text,
    View,
    TextInput,
    TouchableHighlight,
    Image
} from 'react-native';
import LoginGrpc from "./LoginGrpc"
import styles from '../../styles/Style'

class Login extends Component {

    constructor(props) {
        super(props);
        this.state = {
            email: '',
            password: '',
        }
    }
    onClickListener = (viewId) => {
        Alert.alert("Alert", "Button pressed " + viewId);

        LoginGrpc.check(
            "one",
            "two", 
        (msg) => {
            Alert.alert(msg);
          },
          (x) => {
            Alert.alert("26"+x,x);

          },
        );

        switch (viewId) {
            case 'login':
                //TODO: CHECK  HERE if login correct
                this.props.navigation.navigate('app');
                break;

            case 'restore_password':
                this.props.navigation.navigate('RestorePassword');
                break;

            default:
                break;
        }
    }

    render() {
        return (
            <View style={styles.container}>
                <View style={styles.inputContainer}>
                    <Image style={styles.inputIcon} source={require('../../img/mail.png')} />
                    <TextInput style={styles.inputs}
                        placeholder="Email"
                        keyboardType="email-address"
                        underlineColorAndroid='transparent'
                        onChangeText={(email) => this.setState({ email })} />
                </View>

                <View style={styles.inputContainer}>
                    <Image style={styles.inputIcon} source={require('../../img/key.png')} />
                    <TextInput style={styles.inputs}
                        placeholder="Password"
                        secureTextEntry={true}
                        underlineColorAndroid='transparent' git
                        onChangeText={(password) => this.setState({ password })} />
                </View>

                <TouchableHighlight style={[styles.buttonContainer, styles.loginButton]} onPress={() => this.onClickListener('login')}>
                    <Text style={styles.loginText}>Login</Text>
                </TouchableHighlight>

                <TouchableHighlight style={styles.buttonContainer} onPress={() => this.onClickListener('restore_password')}>
                    <Text>Forgot your password?</Text>
                </TouchableHighlight>

                <TouchableHighlight style={styles.buttonContainer} onPress={() => this.props.navigation.navigate('Register')}>
                    <Text>Register</Text>
                </TouchableHighlight>
            </View>
        );
    }
}
export default Login