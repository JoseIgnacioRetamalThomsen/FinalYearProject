import React, {Component} from 'react'
import {
    Text,
    View,
    TextInput,
    TouchableHighlight,
    Image,
    NativeModules, TouchableOpacity
} from 'react-native'
import AsyncStorage from '@react-native-community/async-storage'
import styles from '../../styles/Style'

export default class Register extends Component {
    constructor(props) {
        super(props);
        this.state = {
            message: "",
            email: '',
            password: '',
            cpassword: '',
            token: '',
            isUser: false,
            hiddenPassword: true,
            type: 'input'
        }
    }

    setPasswordVisibility = () => {
        this.setState({hiddenPassword: !this.state.hiddenPassword});
    }

    async onClickListener() {
        let emailRegex = /^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i
        let isCorrectEmail = emailRegex.test(this.state.email)

        let passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{6,})/i
        let isCorrectPassword = passwordRegex.test(this.state.password)

        if (isCorrectEmail === false)
            this.setState({message: 'Email is not correct'})
        else if (isCorrectPassword == false)
            this.setState({
                message: "Password has to be at least 6 characters long, " +
                    "contain at least 1 lowercase, 1 uppercase alphabetical character, " +
                    "1 numeric character and 1 special symbol"
            })
        else if (this.state.cpassword == '')
            this.setState({message: "Please enter confirm password"})
        else if (this.state.password !== this.state.cpassword) {
            this.setState({message: "\nPassword did not match: Please try again..."})
            return false;
        } else {
            NativeModules.LoginModule.createUser(
                this.state.email,
                this.state.password,
                (err) => {
                    this.setState({message: err})
                },
                async (token) => {
                    if (token.includes("Duplicate")) {
                        this.setState({message: 'Email is already registered'})
                    } else if (token.includes("Exception")) {
                        this.setState({message: 'Server unavailable'})
                    } else {
                        console.log("token here " + token)
                        if (token !== "") {
                            try {
                                await AsyncStorage.setItem(this.state.email, token)
                                console.log("this.state.email, token" + this.state.email, token)
                                this.props.navigation.navigate('app')
                            } catch (e) {
                                console.log("Error ", e);
                                // try {
                                //     value = await AsyncStorage.getItem(this.state.email)
                                //     console.log("value in get" + value)
                                // } catch (e) {
                                // }
                                // console.log("asyncstorage " + await AsyncStorage.getItem(this.state.email))
                                // console.log("value " + value)
                                // console.log("token" + token)
                                //this.setState({message: 'Success'})
                                this.props.navigation.navigate('app')
                            }
                        } else this.setState({message: 'Email is already registered'})
                    }
                })

        }
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
                               clearTextOnFocus={false}
                               placeholder="Password"
                               secureTextEntry={this.state.hiddenPassword}
                               underlineColorAndroid='transparent'
                               value={this.state.password}
                               onChangeText={(password) => this.setState({password})}/>
                    <TouchableOpacity activeOpacity={0.8} style={styles.touchableButton}
                                      onPress={this.setPasswordVisibility}>
                        <Image
                            source={(this.state.hiddenPassword) ? require('../../img/hide.png') : require('../../img/show.png')}
                            style={styles.buttonImage}/>
                    </TouchableOpacity>
                </View>

                <View style={styles.inputContainer}>
                    <TouchableOpacity activeOpacity={0.8} style={styles.touchableButton}
                                      onPress={this.setPasswordVisibility}>
                        <Image
                            source={(this.state.hiddenPassword) ? require('../../img/hide.png') : require('../../img/show.png')}
                            style={styles.buttonImage}/>
                    </TouchableOpacity>

                    <Image style={styles.inputIcon} source={require('../../img/key.png')}/>
                    <TextInput style={styles.inputs}
                               clearTextOnFocus={false}
                               placeholder="Confirm Password"
                               secureTextEntry={this.state.hiddenPassword}
                               value={this.state.cpassword}
                               underlineColorAndroid='transparent'
                               onChangeText={(cpassword) => this.setState({cpassword})}/>
                </View>

                <TouchableHighlight style={[styles.buttonContainer, styles.loginButton]}
                                    onPress={() => this.onClickListener()}>
                    <Text style={styles.loginText}>Create Account</Text>
                </TouchableHighlight>

            </View>
        );
    }
}
