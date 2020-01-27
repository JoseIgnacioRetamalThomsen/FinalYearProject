import React, {Component} from 'react';
import {
    Text,
    View,
    TextInput,
    TouchableHighlight,
    Image,
} from 'react-native';
import styles from '../../styles/Style'
// import LoginGrpc from './LoginGrpc'
import { NativeModules } from 'react-native';

class Login extends Component {

    constructor(props) {
        super(props);
        this.state = {
            email: '',
            password: '',
        }
    }

    onClickListenerLogin = () => {
        NativeModules.LoginModule.loginUser(
            "email",
            "password",
            (msg) => {
                alert(msg);
            },
            (x) => {
                alert("Login message " + x)
            }
        )
    }
    onClickListener = (viewId) => {
        switch (viewId) {
            // case 'login':
            //     //TODO: CHECK  HERE if login correct
            //     NativeModules.LoginModule.check(
            //         "email",
            //         "password",
            //         (msg) => {
            //           // Alert.alert(msg);
            // },
            //     (x) => {
            //             //Alert.alert(" !!! ", + x, x)
            //     }
            //     )
            //     this.props.navigation.navigate('app');
            //     break;

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
                    <Image style={styles.inputIcon} source={require('../../img/mail.png')}/>
                    <TextInput style={styles.inputs}
                               placeholder="Email"
                               keyboardType="email-address"
                               underlineColorAndroid='transparent'
                               onChangeText={(email) => this.setState({email})}/>
                </View>

                <View style={styles.inputContainer}>
                    <Image style={styles.inputIcon} source={require('../../img/key.png')}/>
                    <TextInput style={styles.inputs}
                               placeholder="Password"
                               secureTextEntry={true}
                               underlineColorAndroid='transparent'
                               onChangeText={(password) => this.setState({password})}/>
                </View>

                <TouchableHighlight style={[styles.buttonContainer, styles.loginButton]}
                                    onPress={() => this.onClickListenerLogin()}>
                    <Text style={styles.loginText}>Login</Text>
                </TouchableHighlight>

                <TouchableHighlight style={styles.buttonContainer}
                                    onPress={() => this.onClickListener('restore_password')}>
                    <Text>Forgot your password?</Text>
                </TouchableHighlight>

                <TouchableHighlight style={styles.buttonContainer}
                                    onPress={() => this.props.navigation.navigate('Register')}>
                    <Text>Register</Text>
                </TouchableHighlight>
            </View>
        );
    }
}
export default Login
