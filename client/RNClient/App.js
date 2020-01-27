import React from 'react';
import {Image, Dimensions, View} from 'react-native';
import { createAppContainer, createSwitchNavigator } from 'react-navigation';
import { createBottomTabNavigator } from 'react-navigation-tabs';
import { createStackNavigator } from 'react-navigation-stack';
import { createDrawerNavigator } from 'react-navigation-drawer';
import { IMAGE } from './src/constants/Image'

import Feed from './src/component/tabs/feed/Feed'
import MyPosts from './src/component/tabs/myPosts/MyPosts'
import FeedDetail from './src/component/tabs/feed/FeedDetail'
import PostDetails from './src/component/tabs/myPosts/PostDetails'
import SideMenu from './src/component/SideMenu'
import Profile from './src/component/drawer/Profile'
import Settings from './src/component/drawer/Settings'
import Login from './src/component/auth/Login'
import Register from './src/component/auth/Register'
import RestorePassword from './src/component/auth/RestorePassword'
import LoadImage from './src/component/LoadImage'
const navOptionHandler = (navigation) => ({
  header: null
})

const FeedStack = createStackNavigator({
  Feed: {
    screen: Feed,
    navigationOptions: navOptionHandler
  },
  FeedDetail: {
    screen: FeedDetail,
    navigationOptions: navOptionHandler
  }
})
const MyPostsStack = createStackNavigator({
  MyPosts: {
    screen: MyPosts,
    navigationOptions: navOptionHandler
  },
  SearchDetail: {
    screen: PostDetails,
    navigationOptions: navOptionHandler
  },
  LoadImage: {
    screen: LoadImage,
    navigationOptions: navOptionHandler
  }
})

const MainTabs = createBottomTabNavigator({
  Feed: {
    screen: FeedStack,
    navigationOptions: {
      tabBarLabel: 'Feed',
      tabBarIcon: ({ tintColor }) => (
        <Image
          source={IMAGE.ICON_FEED}
          resizeMode="contain"
          style={{ width: 20, height: 20 }}
        />
      )
    }
  },
  MyPosts: {
    screen: MyPostsStack,
    navigationOptions: {
      tabBarLabel: 'MyPosts',
      tabBarIcon: ({ tintColor }) => (
          <Image
              source={IMAGE.ICON_DEFAULT_PROFILE}
              resizeMode="contain"
              style={{ width: 20, height: 20 }}
          />
      )
    }
  }
});

const MainStack = createStackNavigator({
  Home: {
    screen: MainTabs,
    navigationOptions: navOptionHandler
  },
  Profile: {
    screen: Profile,
    navigationOptions: navOptionHandler
  },
  Settings: {
    screen: Settings,
    navigationOptions: navOptionHandler
  },
}, { initialRouteName: 'Home' })

const appDrawer = createDrawerNavigator({
  drawer: MainStack,
},
  {
    contentComponent: SideMenu,
    drawerWidth: Dimensions.get('window').width * 3 / 4
  }
)

const authStack = createStackNavigator({
  Login: {
    screen: Login,
    navigationOptions: navOptionHandler
  },

  Register: {
    screen: Register,
    navigationOptions: navOptionHandler
  },
  RestorePassword: {
    screen: RestorePassword,
    navigationOptions: navOptionHandler
  },

})

const MyApp = createSwitchNavigator({
  app: appDrawer,
  auth: authStack
},
  {
    initialRouteName: 'auth'
  })

const AppNavigation =  createAppContainer(MyApp);

export default class App extends  React.Component{
  render(){
    return(
  <AppNavigation/>
    )
  }
}
