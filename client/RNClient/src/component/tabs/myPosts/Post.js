import React, { Component } from 'react';
import {
    Text,
    View,
    TextInput,
    TouchableHighlight, TouchableOpacity
} from 'react-native';
import styles from '../../../styles/Style'
import CustomHeader from "../../CustomHeader";
import LoadImage from "../../LoadImage";
import {Root} from "native-base";
import {onClickAddImage, onSelectedImage} from '../../functions'
export default class Post extends Component {

    constructor(props) {
        super(props);
        this.state = {
            img: '',
            text: '',
            fileList: []
        }
    }
    render() {
        return (
            <Root>
                <View style={{flex: 1}}>
                    <CustomHeader title="Post" isHome={false} navigation={this.props.navigation}/>
                </View>
                {/*<View>*/}
                {/*    <TextInput*/}
                {/*        placeholder="Start writing your post here.. "*/}
                {/*        underlineColorAndroid='transparent'*/}
                {/*        onChangeText={(text) => this.setState({ text })} />*/}
                {/*</View>*/}
                {/*<TouchableOpacity onPress={onClickAddImage} >*/}
                {/*    <Text> Press Add Image</Text>*/}
                {/*</TouchableOpacity>*/}
                <LoadImage/>
            </Root>
        )
    }
}
