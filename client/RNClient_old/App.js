import Feed from './src/component/tabs/feed/Feed'
import Search from './src/component/tabs/search/Search'
import FeedDetail from './src/component/tabs/feed/FeedDetail'
import SearchDetail from './src/component/tabs/search/SearchDetail'
import SideMenu from './src/component/SideMenu'
import Profile from './src/component/drawer/Profile'
import Settings from './src/component/drawer/Settings'
import Login from './src/component/auth/Login'
import Register from './src/component/auth/Register'
import RestorePassword from './src/component/auth/RestorePassword'
import { Root } from 'native-base';

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
const SearchStack = createStackNavigator({
  Search: {
    screen: Search,
    navigationOptions: navOptionHandler
  },
  SearchDetail: {
    screen: SearchDetail,
    navigationOptions: navOptionHandler
  },
  LoadImage: {
    screen: LoadImage,
    navigationOptions: navOptionHandler
  }
})

const MainTabs = createBottomTabNavigator({
  Search: {
    screen: SearchStack,
    navigationOptions: {
      tabBarLabel: 'Search',
      tabBarIcon: ({ tintColor }) => (
        <Image
          source={IMAGE.ICON_DEFAULT_PROFILE}
          resizeMode="contain"
          style={{ width: 20, height: 20 }}
        />
      )
    }
  },
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
  // <Root>
  <AppNavigation/>

           // </Root>
    )
  }
}
