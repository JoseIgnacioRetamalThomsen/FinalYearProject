import React, { Component } from 'react';
import { Header, Left, Body, Right, Button, Icon, Title } from 'native-base';

class CustomHeader extends Component {
    render() {
        let { title, isHome } = this.props;
        return (
            <Header style ={{backgroundColor:'#007AFF'}}>
                <Left>
                    {
                        isHome ?
                            <Button transparent onPress={() => this.props.navigation.openDrawer()}>
                                <Icon name='menu' />
                            </Button> :
                            <Button transparent onPress={() => this.props.navigation.goBack()}>
                                <Icon name='arrow-back' />
                            </Button>
                    }
                </Left>
                <Body>
                    <Title>{title}</Title>
                </Body>
                <Right>
                    <Button transparent onPress={() => this.props.navigation.navigate("DisplayCity")}>
                        <Icon name='search' />
                    </Button>
                </Right>
            </Header>
        );
    }

}
export default CustomHeader
