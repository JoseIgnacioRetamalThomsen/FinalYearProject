import React from 'react';
import { Text, View, Image, Dimensions, SafeAreaView, ScrollView } from 'react-native';
import { createAppContainer } from 'react-navigation';
import { createBottomTabNavigator, createStackNavigator, createDrawerNavigator } from 'react-navigation-tabs';
import { Container, Header, Left, Body, Right, Button, Icon, Title, List, ListItem } from 'native-base';
import { IMAGE } from '../constants/Image'
// import Feed from './Feed'
// import Profile from './Profile'
// import Settings from './Settings'
class Profile extends React.Component {
  render() {
    return (
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        <Text>Profile!</Text>
      </View>
    );
  }
}

class Feed extends React.Component {
  render() {
    return (
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        <Text>Feed!</Text>
      </View>
    );
  }
}

class Settings extends React.Component {
    render() {
      return (
        <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
          <Text>Settings!</Text>
        </View>
      );
    }
  }

const TabNavigator = createBottomTabNavigator({
  Profile: {
    screen: Profile,
    navigationOptions: {
        tabBarLabel: 'Profile',
        tabBarIcon: ({ tintColor }) => (
            <Image
            source = {IMAGE.ICON_PROFILE}
            resizeMode= "contain"
            style={{width: 20, height: 20}}
            />
        )
    }
  },
  Feed: {
    screen: Feed,
    navigationOptions: {
        tabBarLabel: 'Feed',
        tabBarIcon: ({ tintColor }) => (
            <Image
            source = {IMAGE.ICON_FEED}
            resizeMode= "contain"
            style={{width: 20, height: 20}}
            />
        )
    }
  },
  Settings: {
    screen: Settings,
    navigationOptions: {
        tabBarLabel: 'Settings',
        tabBarIcon: ({ tintColor }) => (
            <Image
            source = {IMAGE.ICON_SETTINGS}
            resizeMode= "contain"
            style={{width: 20, height: 20}}
            />
        )
    }
  },
});

class SideMenu extends React.Component{
    render(){
        <SafeAreaView style ={{flex:1}}>
            <View style={{height: 150, alignItems: 'center', justifyContent:'center'}}>
            <Image source={IMAGE.ICON_USER_DEFAULT} style={{height:120, borderRadius:60}}/>
            </View>

        </SafeAreaView>
    }
}

// const MainStack = createStackNavigator({
//     HomeScreen: {
//         screen: HomeScreen,
//         navigationOptions: navOptionHandler
//     },
//     Settings: {
//         screen: Settings,
//         navigationOptions: navOptionHandler
//     },
//     Feed: {
//         screen: Feed,
//         navigationOptions: navOptionHandler
//     }
// }, {initialRouteName: 'Home'})
// const appDrawer = createDrawerNavigator({
//     drawer: MainStack
// },
// {
//     contentComponent: SideMenu,
//     drawerWidth: (Dimensions.get('window').width * 3/4)
// }
// )

export default createAppContainer(TabNavigator);