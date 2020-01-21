import React, { Component } from 'react';
import { View } from 'react-native';
import { Text } from 'native-base';
import CustomHeader from '../CustomHeader'

class Settings extends Component {
  render() {
    return (
      <View style={{ flex: 1 }}>
        <CustomHeader title="Settings" navigation={this.props.navigation} />
        <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
          <Text>Settings!</Text>
        </View>
      </View>
    );
  }
}
export default Settings