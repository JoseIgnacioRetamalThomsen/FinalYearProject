import React from 'react';
import HelloWorld from './HelloWorld';
HelloWorld.sayHello("Hello from native-react");
import { StyleSheet, Text, View } from 'react-native';
import { createAppContainer } from "react-navigation";
import { createStackNavigator} from "react-navigation-stack";
import Login from './components/Login';
import Register from './components/Register';
import HomeScreen from './components/HomeScreen';
import { logger } from 'react-native-logger'

export default class App extends React.Component {
  render() {
    return <AppContainer />;
  }
}

const AppNavigator = createStackNavigator({
  Login: {
    screen: Login
  },
  Register: {
    screen: Register
  },
  HomeScreen: {
    screen: HomeScreen
  }
});

const AppContainer = createAppContainer(AppNavigator);

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});