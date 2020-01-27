import React, { Component } from 'react';
import {
    Text,
    View,
    TextInput,
    TouchableHighlight
} from 'react-native';
import styles from '../../../styles/Style'
import CustomHeader from "../../CustomHeader";
import LoadImage from "../../LoadImage";
import {Root} from "native-base";

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
            <Root>
                <View style={{flex: 1}}>
                    <CustomHeader title="Post" isHome={false} navigation={this.props.navigation}/>
                </View>
                <View>
                    <TextInput
                        placeholder="Start writing your post here.. "
                        underlineColorAndroid='transparent'
                        onChangeText={(text) => this.setState({ text })} />
                </View>

                    <LoadImage/>

                <View >
                    <TouchableHighlight style={[styles.buttonContainer, styles.loginButton]} onPress={() => this.onClickListener()}>
                        <Text style={styles.loginText}> Post</Text>
                    </TouchableHighlight>
                </View>

            </Root>
        )
    }
}
