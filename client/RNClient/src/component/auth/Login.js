import React, {Component} from 'react'
import {
    Text,
    View,
    TextInput,
    TouchableHighlight,
    Image,
    NativeModules
} from 'react-native'
import styles from '../../styles/Style'
import AsyncStorage from '@react-native-community/async-storage'
import {logger} from "react-native-logger"

class Login extends Component {
    constructor(props) {
        super(props)
        this.state = {
            email: '',
            password: '',
            message: '',
            token: ''
        }
    }

    async onClickListenerLogin() {
        NativeModules.LoginModule.loginUser(
            this.state.email,
            this.state.password,
            (err) => {
                logger.log(err.message())
                this.setState({message: 'Incorrect email or password'})
            },
            async (token) => {
                let value
                if (token === "") {
                    this.setState({message: 'Email is not registered'})
                } else {
                    try {
                        await AsyncStorage.setItem(this.state.email, token)
                    } catch (e) {
                        logger.log(e)
                    }
                    try {
                        value = await AsyncStorage.getItem(this.state.email)
                    } catch (e) {
                    }
                    this.setState({token: value})
                    console.log("token " + token)
                    this.setState({message: 'Success'})
                    this.props.navigation.navigate('app')
                }
            }
        )
    }

    render() {
        return (
            <View style={styles.container}>
                <Text>
                    {this.state.message}
                </Text>
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
                                    onPress={() => this.props.navigation.navigate('restore_password')}>
                    <Text>Forgot your password?</Text>
                </TouchableHighlight>

                <TouchableHighlight style={styles.buttonContainer}
                                    onPress={() => this.props.navigation.navigate('Register')}>
                    <Text>Register</Text>
                </TouchableHighlight>
            </View>
        )
    }
}

export default Login
