import { h, input } from 'react-hyperscript-helpers';
import { Component } from 'react';
import { connect } from 'react-redux';
import styled from 'styled-components';
import { get } from 'lodash';

import { DropzoneConn, Button } from '@portal/ui';
import { getTheme } from '@portal/theme';

import { fetchCreatePostAction } from '../actions';

const Container = styled.div`
  margin-top: 150px;
  display: flex;
  flex-direction: column;
  padding: 50px;
  justify-content: center;
`;

const ImageContainer = styled.img`
  max-width: 400px;
  max-height: 400px;
  object-fit: contain;
  align-self: center;
`;

const InputContainer = styled.div`
  margin: 30px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  input {
    max-width: 400px;
    width: 100%;
    margin: 10px 0;
    padding: 8px;
    font-size: 1em;
  }
`;

export class NewPostPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      caption: '',
      imageUrl: null,
    };
  }

  handleAttributeChange(e) {
    this.setState({
      [e.target.name]: e.target.value,
    });
  }

  render() {
    const { theme, fetchCreatePost } = this.props;
    const { caption, imageUrl, imageType } = this.state;

    return h(Container, [
      imageUrl === null
        ? h(DropzoneConn, {
            onDropFile: (file) =>
              this.setState({
                caption,
                imageUrl: URL.createObjectURL(file),
                imageType: get(file, 'type', ''),
              }),
          })
        : h(ImageContainer, { src: imageUrl }),
      h(InputContainer, [
        input({
          id: 'caption',
          type: 'text',
          name: 'caption',
          placeholder: 'Caption',
          value: caption,
          onChange: this.handleAttributeChange.bind(this),
        }),
        h(
          Button,
          {
            color: theme.primary,
            onClick: () => fetchCreatePost({ caption, imageUrl, imageType }),
          },
          ['Post'],
        ),
      ]),
    ]);
  }
}

const mapStateToProps = (state) => ({
  theme: getTheme(state),
});

const mapDispatchToProps = (dispatch) => ({
  fetchCreatePost: ({ caption, imageUrl, imageType }) =>
    dispatch(fetchCreatePostAction({ caption, imageUrl, imageType })),
});

export const NewPostPageConn = connect(
  mapStateToProps,
  mapDispatchToProps,
)(NewPostPage);
