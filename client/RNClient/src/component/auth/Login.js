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
            token: '',
            isUser: false,
        }
    }

    async onClickListenerLogin() {
       //  let emailRegex = /^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i
       //  let isCorrectEmail = emailRegex.test(this.state.email)
       //
       //  let passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{6,})/i
       //  let isCorrectPassword = passwordRegex.test(this.state.password)
       //
       //  if (isCorrectEmail === false)
       //      this.setState({message: 'Email is not correct'})
       //  else if (isCorrectPassword === false)
       //      this.setState({
       //          message: "Password has to be at least 6 characters long, " +
       //              "contain at least 1 lowercase, 1 uppercase alphabetical character, " +
       //              "1 numeric character and 1 special symbol"
       //      })
       // else {
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
   // }

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
