console.disableYellowBox = true

import DisplayCities from "./src/component/tabs/feed/DisplayCities";
import React from 'react';
import {Image, Dimensions, NativeModules} from 'react-native';
import {createAppContainer, createSwitchNavigator} from 'react-navigation';
import {createBottomTabNavigator} from 'react-navigation-tabs';
import {createStackNavigator} from 'react-navigation-stack';
import {createDrawerNavigator} from 'react-navigation-drawer';
import {IMAGE} from './src/constants/Image'
import CityDetail from './src/component/tabs/feed/CityDetail'
import SideMenu from './src/component/utils/SideMenu'
import Profile from './src/component/drawer/Profile'
import Settings from './src/component/drawer/Settings'
import Login from './src/component/auth/Login'
import Register from './src/component/auth/Register'
import RestorePassword from './src/component/auth/RestorePassword'
import WelcomePage from "./src/component/utils/WelcomePage";
import CreateCity from "./src/component/tabs/feed/CreateCity";
import PlaceDetail from "./src/component/tabs/feed/PlaceDetail";
import AsyncStorage from "@react-native-community/async-storage";
import SearchCity from "./src/component/tabs/feed/SearchCity";
import YourCityDetail from "./src/component/tabs/myCity/YourCityDetail";
import SpecialHeader from "./src/component/headers/SpecialHeader";

const {width: width} = Dimensions.get('window')
const {height: height} = Dimensions.get('window')
const navOptionHandler = (navigation) => ({
    header: null
})

const FeedStack = createStackNavigator({
    DisplayCities: {
        screen: DisplayCities,
        navigationOptions: navOptionHandler
    },
    DisplayCity: {
        screen: SearchCity,
        navigationOptions: navOptionHandler
    },
    SearchCity: {
        screen: SearchCity,
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
    PlaceDetail: {
        screen: PlaceDetail,
        navigationOptions: navOptionHandler
    },
})

const MyPostsStack = createStackNavigator({
    YourCityDetail: {
        screen: YourCityDetail,
        navigationOptions: navOptionHandler
    },
    SpecialHeader: {
        screen: SpecialHeader,
        navigationOptions: navOptionHandler
    },
})

const MainTabs = createBottomTabNavigator({
    Feed: {
        screen: FeedStack,
        navigationOptions: {
            tabBarLabel: 'Feed',
            tabBarIcon: ({tintColor}) => (
                <Image
                    source={IMAGE.ICON_FEED}
                    resizeMode="contain"
                    style={{width: width/20, height: height/20}}
                />
            )
        }
    },
    MyCity: {
        screen: MyPostsStack,
        navigationOptions: {
            tabBarLabel: 'My City',
            tabBarIcon: ({tintColor}) => (
                <Image
                    source={IMAGE.ICON_HOME}
                    resizeMode="contain"
                    style={{width: width/15, height: height/15}}
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
    // Logout: {
    //     screen: Logout,
    //     navigationOptions: navOptionHandler
    // },
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
        initialRouteName: 'loading'
    })

const AppNavigation = createAppContainer(MyApp);

export default class App extends React.Component {


    componentDidMount() {
        AsyncStorage.getAllKeys((err, keys) => {
            AsyncStorage.multiGet(keys, (err, stores) => {
                stores.map((result, i, store) => {
                    let email = store[i][0];
                    let token = store[i][1]

                    if (token !== null) {
                        NativeModules.ProfilesModule.getVisitedCities(
                            token,
                            email,
                            (err) => {
                                console.log(err)
                            },
                            (json) => {
                                let citiesMap = JSON.parse(json)
                                citiesMap.forEach(myFunction);

                                function myFunction(item, index) {
                                    global.visitedCityMap[item.cityId] = true
                                }
                            })
                        NativeModules.ProfilesModule.getVisitedPlaces(
                            token,
                            email,
                            (err) => {
                                console.log(err)
                            },
                            (placesJson) => {
                                let placesMap = JSON.parse(placesJson)
                                placesMap.forEach(myFunction)

                                function myFunction(item, index) {
                                    global.visitedPlaceMap[item.id] = true
                                }
                            })
                    }
                })
            })
        })

    }

    render() {
        return (
            <AppNavigation/>
        )
    }
}
