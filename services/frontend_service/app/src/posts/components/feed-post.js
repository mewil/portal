import { h } from 'react-hyperscript-helpers';
import { connect } from 'react-redux';
import styled from 'styled-components';
import { get } from 'lodash';

import { getTheme } from '@portal/theme';
import { getUser, fetchUserAction } from '@portal/user';
import { LikeButton } from '@portal/ui';

import { fetchLikePostAction } from '../actions';

const Container = styled.div`
  max-width: 600px;
  margin-bottom: 24px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  border-radius: 3px;
  padding: 4px;
`;

const UserContainer = styled.div`
  padding: 6px;
  font-size: 20px;
`;

const LikeContainer = styled.div`
  display: flex;
  flex-direction: row;
  padding: 10px 6px;
  margin-top: 4px;
`;

const LikeText = styled.div`
  display: flex;
  margin: 0px 16px;
  font-size: 14px;
`;

const CaptionContainer = styled.div`
  padding: 6px;
  font-size: 14px;
`;

const ImageContainer = styled.img`
  max-width: 600px;
  max-height: 600px;
  object-fit: contain;
  align-self: center;
  box-shadow: 0 2px 2px 0 rgba(133, 132, 155, 0.16),
    0 2px 2px 0 rgba(133, 132, 155, 0.4);
`;

export const FeedPost = ({ post, user, fetchUser, fetchLikePost, theme }) => {
  if (user === null) fetchUser({ userId: get(post, 'userId') });
  const { fileId, caption, likeCount, postId } = post;
  return h(Container, [
    h(UserContainer, [get(user, 'username')]),
    h(ImageContainer, { src: `/v1/file/${fileId}` }),
    h(LikeContainer, { onClick: () => fetchLikePost({ postId }) }, [
      h(LikeButton, { theme, active: true }),
      h(LikeText, [likeCount]),
    ]),
    h(CaptionContainer, [caption]),
  ]);
};

const mapStateToProps = (state, { post: { userId } }) => ({
  user: getUser(state, userId),
  theme: getTheme(state),
});

const mapDispatchToProps = (dispatch) => ({
  fetchUser: ({ userId }) => dispatch(fetchUserAction({ userId })),
  fetchLikePost: ({ postId }) => dispatch(fetchLikePostAction({ postId })),
});

export const FeedPostConn = connect(
  mapStateToProps,
  mapDispatchToProps,
)(FeedPost);
