import React from 'react';
import {View, Text, default as AsyncStorage} from 'react-native';

class SplashScreen extends React.Component {
    performTimeConsumingTask = async() => {
        return new Promise((resolve) =>
            setTimeout(
                () => { resolve('result') },
                2000
            )
        )
    }

    async componentDidMount() {
        AsyncStorage.getItem("token", token).then(value => {
            if(value===null){
                AsyncStorage.setItem("token", true).then(() => {
                    this.props.navigation.replace('app')
                })
            }
            else{
                this.props.navigation.replace('auth')
            }
        })
    }

    render() {
        return (
            <View style={styles.viewStyles}>
                <Text style={styles.textStyles}>
                    Blitz Reading
                </Text>
            </View>
        );
    }
}

const styles = {
    viewStyles: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
        backgroundColor: 'orange'
    },
    textStyles: {
        color: 'white',
        fontSize: 40,
        fontWeight: 'bold'
    }
}

export default SplashScreen;
