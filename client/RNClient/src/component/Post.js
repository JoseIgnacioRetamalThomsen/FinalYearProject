import React, { Component } from 'react';
import {
    Text,
    View,
    TextInput,
    TouchableHighlight,
    Image
} from 'react-native';
import styles from '../styles/Style'

export default class Post extends Component {

    constructor(props) {
        super(props);
        state = {
            img: '',
            text: '',
        }
    }

    onClickListener = (viewId) => {
        alert("submit post")
    }

    render() {
        return (
            <View style={styles.container}>
                {/*<View style={styles.inputContainer}>*/}
                {/*    /!*<Image source={require('../img/gmit.jpg')} style={{width:300, height:200, justifyContent: 'center'}}/>*!/*/}
                {/*    </View>*/}
                {/*    <View>*/}
                {/*    <TextInput*/}
                {/*               placeholder="Start writing your post here.. "*/}
                {/*               underlineColorAndroid='transparent'*/}
                {/*               onChangeText={(text) => this.setState({ text })} />*/}
                {/*</View>*/}


                {/*<TouchableHighlight style={[styles.buttonContainer, styles.loginButton]} onPress={() => this.onClickListener()}>*/}
                {/*    <Text style={styles.loginText}>Add Post</Text>*/}
                {/*</TouchableHighlight>*/}

            </View>
        );
    }
}
