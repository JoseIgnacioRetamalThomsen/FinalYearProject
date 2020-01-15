import React, { Component } from 'react';
import { Header, Left, Body, Right, Button, Icon, Title } from 'native-base';

class CustomHeader extends Component {
    render() {
        let { title, isHome } = this.props;
        return (
            <Header>
                <Left>
                    {
                        isHome ?
                            <Button transparent onPress={() => this.props.navigation.openDrawer()}>
                                <Icon name='menu' />
                            </Button> :
                            <Button transparent>
                                <Icon name='arrow-back' onPress={() => this.props.navigation.goBack()} />
                            </Button>
                    }
                </Left>
                <Body>
                    <Title>{title}</Title>
                </Body>
                <Right>

                </Right>
            </Header>
        );
    }
}
export default CustomHeader