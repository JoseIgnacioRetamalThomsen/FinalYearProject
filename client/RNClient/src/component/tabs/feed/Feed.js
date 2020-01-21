import React,  {Component} from 'react';
import { View, TextArea } from 'react-native';
import { Button,  Text } from 'native-base';
import CustomHeader from '../../CustomHeader'
import LoadImage from "../../LoadImage";
import SearchBar from 'react-native-search-bar';

class Feed extends Component {
  render() {
    return (
      <View style={{ flex: 1 }}>
      <CustomHeader  title="Feed"  isHome={true} navigation={this.props.navigation}/>
      <SearchBar
          ref="searchBar"
          placeholder="Search"
          // onChangeText={...}
          // onSearchButtonPress={...}
          // onCancelButtonPress={...}
      />




      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        {/*<Text>Add a post</Text>*/}

        {/*<Button  onPress={() => this.props.navigation.navigate('FeedDetail')}>*/}
        {/*    <Text> Add post</Text>*/}
        {/*</Button>*/}
      </View>
      </View>
    );
  }
}
export default Feed