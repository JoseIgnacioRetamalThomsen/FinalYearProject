import React, {Component} from 'react';
import {
    AsyncStorage,
    Text,
    View,
    TextInput,
    TouchableHighlight,
    Image,
    NativeModules
} from 'react-native';
import styles from '../../styles/Style'

class Login extends Component {

    constructor(props) {
        super(props);
        this.state = {
            email: '',
            password: '',
            message: '',
            isUser: false
        }
    }

    async onClickListenerLogin() {//is it async?
        NativeModules.LoginModule.loginUser(
            this.state.email,
            this.state.password,
            (err) => {
                this.setState({isUser: err})
                alert('err ' + err)
                this.setState({message: 'Incorrect email or password'})
                this.props.navigation.navigate('Register')
            },
            (isUser) => {
                alert('isUser ' + isUser)
                this.setState({message: 'Success'})
                this.props.navigation.navigate('RestorePassword')
            }
        )
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
