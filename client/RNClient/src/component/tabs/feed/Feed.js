import React,  {Component} from 'react';
import { View } from 'react-native';
import { Button,  Text } from 'native-base';
import CustomHeader from '../../CustomHeader'

class Feed extends Component {
  render() {
    return (
      <View style={{ flex: 1 }}>
      <CustomHeader  title="Feed"  isHome={true} navigation={this.props.navigation}/>
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        <Text>Feed!</Text>
        <Button light onPress={() => this.props.navigation.navigate('FeedDetail')}>
            <Text> Go to FeedDetail page</Text>
        </Button>
      </View>
      </View>
    );
  }
}
export default Feed