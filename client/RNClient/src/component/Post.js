import React, { Component } from 'react';
import {
    Text,
    View,
    TextInput,
    TouchableHighlight
} from 'react-native';
import styles from '../styles/Style'
import CustomHeader from "./CustomHeader";
import LoadImage from "./LoadImage";

export default class Post extends Component {

    constructor(props) {
        super(props);
        this.state = {
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
                <View>
                    <CustomHeader title= 'Write a post' isHome={false} navigation={this.props.navigation}/>
                </View>
                <View>
                    <TextInput
                        placeholder="Start writing your post here.. "
                        underlineColorAndroid='transparent'
                        onChangeText={(text) => this.setState({ text })} />
                </View>
                <View>
                    <LoadImage/>
                </View>
                <View>
                    <TouchableHighlight style={[styles.buttonContainer, styles.loginButton]} onPress={() => this.onClickListener()}>
                        <Text style={styles.loginText}>Add Post</Text>
                    </TouchableHighlight>
                </View>
            </View>
        );
    }
}
