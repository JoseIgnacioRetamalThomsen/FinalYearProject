import React, {Component} from 'react';
import {
    Text,
    View,
    TextInput,
    TouchableHighlight,
    Image,
    NativeModules,
    AsyncStorage
} from 'react-native';
import styles from '../../styles/Style'

export default class Register extends Component {
    constructor(props) {
        super(props);
        this.state = {
            message: "",
            email: '',
            password: '',
            cpassword: '',
        }
    }

    storeToken = async (token) => {
        try {
            await AsyncStorage.setItem('userToken', token)
            alert('storeToken() token is ' + token)
        } catch (error) {
            console.log("error inside storeToken")
        }
    }

    async onClickListener  ()  {
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
                    alert("err"+err)
                },
                async (token) => {
                   //alert("token "+token)
                    if(token.includes("Duplicate")){
                        this.setState({message: 'Email is already registered'})
                    }else{
                        await AsyncStorage.setItem(JSON.stringify(this.state.email), token)
                            .then(() => console.log('saved email successfully'))
                            .catch(err => console.error('something went wrong', err));
                        alert("result "+this.state.email +  token)
                        // alert("asyncstorage " + AsyncStorage.setItem(this.state.email, token))
                        //
                        //alert("asyncstorage "+AsyncStorage.getItem(this.state.email, token))
                        this.setState({message: 'Success'})
                        this.props.navigation.navigate('app')
                    }
                }
            )

        }
    }

    render() {
        return (
            <View style={styles.container}>
                {/*<View>*/}
                    <Text style={styles.container}>
                        {this.state.message}
                    </Text>
                {/*</View>*/}
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
                <View style={styles.inputContainer}>
                    <Image style={styles.inputIcon} source={require('../../img/key.png')}/>
                    <TextInput style={styles.inputs}
                               placeholder="Confirm Password"
                               secureTextEntry={true}
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
