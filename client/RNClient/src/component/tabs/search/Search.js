import React,  {Component} from 'react';
import { View } from 'react-native';
import { Button,  Text } from 'native-base';
import CustomHeader from '../../CustomHeader'

class Search extends Component {
    render() {
      return (
        <View style={{ flex: 1 }}>
          <CustomHeader  title="Search"  isHome={true} navigation={this.props.navigation}/>
        <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        <Text>Search page</Text>
          <Button light onPress={() => this.props.navigation.navigate('SearchDetail')}>
              <Text> Go to Search page</Text>
          </Button>
        </View>
        </View>
      );
    }
  }

export default Search
  