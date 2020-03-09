import DisplayCities from "./src/component/tabs/feed/DisplayCities";

console.disableYellowBox = true

import React from 'react';
import {Image, Dimensions, View} from 'react-native';
import {createAppContainer, createSwitchNavigator} from 'react-navigation';
import {createBottomTabNavigator} from 'react-navigation-tabs';
import {createStackNavigator} from 'react-navigation-stack';
import {createDrawerNavigator} from 'react-navigation-drawer';
import {IMAGE} from './src/constants/Image'

import DisplayCityPosts from './src/component/tabs/feed/DisplayCityPosts'
import MyPosts from './src/component/tabs/myPosts/MyPosts'
import CityDetail from './src/component/tabs/feed/CityDetail'
import PostDetails from './src/component/tabs/myPosts/PostDetails'
import SideMenu from './src/component/SideMenu'
import Profile from './src/component/drawer/Profile'
import Settings from './src/component/drawer/Settings'
import Login from './src/component/auth/Login'
import Register from './src/component/auth/Register'
import RestorePassword from './src/component/auth/RestorePassword'
import LoadImage from './src/component/LoadImage'
import WritePost from "./src/component/tabs/myPosts/WritePost";
import Post from "./src/component/tabs/myPosts/Post";
import WelcomePage from "./src/component/WelcomePage";
import MapInput from "./src/component/MapInput";
import CreateCity from "./src/component/tabs/feed/CreateCity";
import CreateCityPost from "./src/component/tabs/feed/CreateCityPost";
import Test from "./src/component/Test";

const navOptionHandler = (navigation) => ({
    header: null
})

const FeedStack = createStackNavigator({
    DisplayCities: {
        screen: DisplayCities,
        navigationOptions: navOptionHandler
    },
    Test:{
        screen: Test,
        navigationOptions: navOptionHandler
    },
    CreateCity: {
        screen: CreateCity,
        navigationOptions: navOptionHandler
    },
    CityDetail: {
        screen: CityDetail,
        navigationOptions: navOptionHandler
    },
    DisplayCityPosts: {
        screen: DisplayCityPosts,
        navigationOptions: navOptionHandler
    },
    CreateCityPost: {
        screen: CreateCityPost,
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
    MapInput: {
        screen: MapInput,
        navigationOptions: navOptionHandler
    },
    Post: {
        screen: Post,
        navigationOptions: navOptionHandler
    },
    WritePost: {
        screen: WritePost,
        navigationOptions: navOptionHandler
    },
    LoadImage: {
        screen: LoadImage,
        navigationOptions: navOptionHandler
    },
})

const MainTabs = createBottomTabNavigator({
    Feed: {
        screen: FeedStack,
        navigationOptions: {
            tabBarLabel: 'Cities',
            tabBarIcon: ({tintColor}) => (
                <Image
                    source={IMAGE.ICON_FEED}
                    resizeMode="contain"
                    style={{width: 20, height: 20}}
                />
            )
        }
    },
    MyPosts: {
        screen: MyPostsStack,
        navigationOptions: {
            tabBarLabel: 'Places',
            tabBarIcon: ({tintColor}) => (
                <Image
                    source={IMAGE.ICON_DEFAULT_PROFILE}
                    resizeMode="contain"
                    style={{width: 20, height: 20}}
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
}, {initialRouteName: 'Home'})

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

        loading: {
            screen: WelcomePage
        },
        app: appDrawer,
        auth: authStack
    },
    {
        initialRouteName: 'loading'//change back to loading
    })

const AppNavigation = createAppContainer(MyApp);

export default class App extends React.Component {
    render() {
        return (
            <AppNavigation/>
        )
    }
}
