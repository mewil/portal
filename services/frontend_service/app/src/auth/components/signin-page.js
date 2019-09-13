import { h, input } from 'react-hyperscript-helpers';
import { Component } from 'react';
import styled from 'styled-components';
import { connect } from 'react-redux';

import { getTheme } from '@portal/theme';
import { Button, TabGroup } from '@portal/ui';

import { fetchSignInAction, fetchSignUpAction } from '../actions';

const Container = styled.div`
  width: 500px;
  max-width: calc(100% - 40px);
  min-height: calc(100vh - 30px - 2rem - 80px);
  padding: 20px 0 50px;
  margin: 0 auto;
`;

const TabContainer = styled.div`
  display: flex;
  flex-direction: column;
`;

const InputContainer = styled.div`
  margin: 30px 0;
  input {
    width: 100%;
    margin: 10px 0;
    padding: 8px;
    font-size: 1em;
  }
`;

const ButtonContainer = styled.div`
  display: flex;
  flex-direction: row;
  align-self: center;
`;

export class SignInPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: '',
      email: '',
      password: '',
      isSigningUp: true,
    };
    this.tabSelect = this.tabSelect.bind(this);
  }

  handleAttributeChange(e) {
    this.setState({
      [e.target.name]: e.target.value,
    });
  }

  tabSelect() {
    this.setState({
      isSigningUp: !this.state.isSigningUp,
    });
  }

  render() {
    const { theme, fetchSignUp, fetchSignIn } = this.props;
    const { username, email, password, isSigningUp } = this.state;
    return h(Container, [
      h(TabGroup, {
        tabs: [
          {
            title: 'Sign Up',
            onClick: this.tabSelect,
          },
          {
            title: 'Sign In',
            onClick: this.tabSelect,
          },
        ],
        primaryColor: theme.info,
      }),
      h(TabContainer, [
        h(InputContainer, [
          isSigningUp
            ? input({
                id: 'username',
                type: 'text',
                name: 'username',
                placeholder: 'Username',
                value: username,
                onChange: this.handleAttributeChange.bind(this),
              })
            : null,
          input({
            id: 'email',
            type: 'email',
            name: 'email',
            placeholder: 'Email',
            value: email,
            onChange: this.handleAttributeChange.bind(this),
          }),
          input({
            id: 'password',
            type: 'password',
            name: 'password',
            placeholder: 'Password',
            value: password,
            onChange: this.handleAttributeChange.bind(this),
          }),
        ]),
        h(ButtonContainer, [
          h(
            Button,
            {
              color: theme.primary,
              onClick: () =>
                isSigningUp
                  ? fetchSignUp({ username, email, password })
                  : fetchSignIn({ email, password }),
            },
            ['Confirm'],
          ),
        ]),
      ]),
    ]);
  }
}

const mapStateToProps = (state) => ({
  theme: getTheme(state),
});

const mapDispatchToProps = (dispatch) => ({
  fetchSignIn: ({ email, password }) =>
    dispatch(fetchSignInAction({ email, password })),
  fetchSignUp: ({ username, email, password }) =>
    dispatch(fetchSignUpAction({ username, email, password })),
});

export const SignInPageConn = connect(
  mapStateToProps,
  mapDispatchToProps,
)(SignInPage);
