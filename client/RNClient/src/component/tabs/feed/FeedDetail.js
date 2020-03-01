import React, { Component } from 'react';
import { View } from 'react-native';
import { Text } from 'native-base';
import CustomHeader from '../../CustomHeader'

class FeedDetail extends Component {
    render() {
        return (
            <View style={{ flex: 1 }}>
                <CustomHeader title="CityPosts Detail" navigation={this.props.navigation} />
                <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
                    <Text>FeedDetail!</Text>
                </View>
            </View>
        );
    }
}
export default FeedDetail
