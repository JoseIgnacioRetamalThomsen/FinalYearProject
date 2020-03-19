import React, { Component } from 'react';
import { View } from 'react-native';
import { Text } from 'native-base';
import CustomHeader from '../../CustomHeader'

class PostDetails extends Component {
    render() {
        return (
            <View style={{ flex: 1 }}>
                <CustomHeader title="DisplayCityPosts Detail" navigation={this.props.navigation} />
                <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
                    <Text>SearchDetail!</Text>
                </View>
            </View>
        );
    }
}
export default PostDetails