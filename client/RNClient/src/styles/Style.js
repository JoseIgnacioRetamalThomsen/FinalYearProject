import { StyleSheet } from 'react-native'
import { STYLES } from '../constants/Styles'

export default StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#DCDCDC',
    },
    inputContainer: {
        borderBottomColor: '#F5FCFF',
        backgroundColor: '#FFFFFF',
        borderRadius: 30,
        borderBottomWidth: 1,
        width: 250,
        height: 45,
        marginBottom: 20,
        flexDirection: 'row',
        alignItems: 'center'
    },
    inputs: {
        height: 45,
        marginLeft: 16,
        borderBottomColor: '#FFFFFF',
        flex: 1,
    },
    inputIcon: {
        width: 30,
        height: 30,
        marginLeft: 15,
        justifyContent: 'center'
    },
    buttonContainer: {
        height: 45,
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 20,
        width: 250,
        borderRadius: 30,
    },
    loginButton: {
        backgroundColor: "#00b5ec",
    },
    loginText: {
        color: 'white',
    },
    content:{
        flex: 1,
        alignItems:'center',
        marginTop:50,
        paddingLeft:30,
        paddingRight:30,
        marginBottom:30

    },
    btnPressStyle:{
        backgroundColor: '#0080ff',
        height:50,
        width: STYLES.width - 60,
        alignItems: 'center',
        justifyContent: 'center'
    },
    txtStyle:{
        color:'#ffffff'
    },
    itemImage:{
        backgroundColor:'#2f455c',
        height:150,
        width: STYLES.width-60,
        borderRadius:8,
        resizeMode:'contain'
    }
})