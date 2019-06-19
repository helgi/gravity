/*
Copyright 2019 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import React from 'react';
import styled from 'styled-components';
import cfg from 'app/config';
import history from 'app/services/history';
import { NavLink } from 'app/components/Router';
import { Text, ButtonOutlined, ButtonPrimary } from 'shared/components';
import Menu, { MenuItem} from 'shared/components/Menu';
import * as Icons from 'shared/components/Icon';

class ContainerMenu extends React.Component {

  static displayName = 'ContainerMenu';

  static defaultProps = {
    menuListCss: () => { },
  }

  constructor(props){
    super(props)
    this.state = {
      open: false,
      anchorEl: null,
    }
  }

  openTerminal(login){
    const {
      name: container,
      pod,
      namespace,
      serverId
    } = this.props.container;

    openNewWindow({
      serverId,
      pod,
      login,
      namespace,
      container,
    });
  }

  onOpen = () => {
    this.setState({ open: true });
  };

  onClose = () => {
    this.setState({ open: false });
  }

  onKeyPress = e => {
    if (e.key === 'Enter' && e.target.value) {
      this.openTerminal(e.target.value);
      this.onClose();
    }
  }

  setRef = e => {
    this.anchorEl = e;
  }

  render() {
    const {
      container,
      anchorOrigin,
      transformOrigin,
      ...styles
    } = this.props;

    const { open } = this.state;
    const { logUrl, name, sshLogins, pod, serverId, namespace } = container;

    return (
      <React.Fragment>
        <ButtonOutlined size="small" p="1" setRef={this.setRef} onClick={this.onOpen} {...styles}>
          {name}
          <Icons.CarrotDown ml="2" fontSize="3" color="text.onDark"/>
        </ButtonOutlined>
        <Menu
          menuListCss={menuListCss}
          anchorOrigin={anchorOrigin}
          transformOrigin={transformOrigin}
          anchorEl={this.anchorEl}
          open={open}
          onClose={this.onClose}
        >
          <LoginItemList
            serverId={serverId}
            logUrl={logUrl}
            title={name}
            namespace={namespace}
            pod={pod}
            container={name}
            logins={sshLogins}
            onKeyPress={this.onKeyPress}
          />
        </Menu>
      </React.Fragment>
    );
  }
}

export const LoginItemList = ({logins, title, serverId, logUrl, container, pod, namespace, onKeyPress}) => {
  logins = logins || [];
  const $menuItems = logins.map((login, key) => {
    const url = cfg.getConsoleInitPodSessionRoute({ login, serverId, container, pod, namespace });
    return (
      <MenuItem as={NavLink} to={url} key={key} target="_blank">
        {login}
      </MenuItem>
    )
  });

  return (
    <React.Fragment>
      <MenuHeader>SSH - {title}</MenuHeader>
      <Input onKeyPress={onKeyPress} type="text" autoFocus placeholder="Enter login name..."/>
      {$menuItems}
      <ButtonPrimary my="3" mx="3" size="small" as={NavLink} to={logUrl}>
        View Logs
      </ButtonPrimary>
    </React.Fragment>
  )
}

function openNewWindow(params){
  let url = cfg.getConsoleInitPodSessionRoute(params);
  url = history.ensureBaseUrl(url);
  window.open(url);//, "", "toolbar=yes,scrollbars=yes,resizable=yes");
}

const MenuHeader = styled.header`
  background: ${props => props.theme.colors.subtle};
  display: block;
  color: ${props => props.theme.colors.grey[400] };
  font-size: 11px;
  line-height: 24px;
  padding: 0 8px;
`

const menuListCss = props => `
  display: flex;
  flex-direction: column;

  ${MenuItem} {
    color: ${props.theme.colors.grey[400]};
    font-size: 12px;
    border-bottom: 1px solid ${props.theme.colors.subtle };
    line-height: 32px;
    margin: 0 8px;
    padding: 0 8px;

    &:hover {
      color: ${props.theme.colors.link};
    }

    &:last-child {
      border-bottom: none;
    }
  }
`

const Input = styled.input`
  background: ${props => props.theme.colors.subtle };
  border: 1px solid ${props => props.theme.colors.subtle };
  border-radius: 4px;
  box-sizing: border-box;
  color: #263238;
  padding: 0 8px;
  height: 32px;
  margin: 8px;
  outline: none;

  &:focus {
    background: ${props => props.theme.colors.light };
    border 1px solid ${props => props.theme.colors.link };
    box-shadow: inset 0 2px 4px rgba(0, 0, 0, .24);
  }

  ::-webkit-input-placeholder { /* Chrome/Opera/Safari */
  color: ${props => props.theme.colors.grey[100]};
  }
  ::-moz-placeholder { /* Firefox 19+ */
    color: ${props => props.theme.colors.grey[100]};
  }
  :-ms-input-placeholder { /* IE 10+ */
    color: ${props => props.theme.colors.grey[100]};
  }
  :-moz-placeholder { /* Firefox 18- */
    color: ${props => props.theme.colors.grey[100]};
  }
`

export default ContainerMenu;
