import { Component } from 'react';
import { h } from 'react-hyperscript-helpers';
import { connect } from 'react-redux';
import InfiniteScroll from 'react-infinite-scroll-component';
import styled from 'styled-components';
import { get } from 'lodash';

import { getFeedPosts, fetchFeedAction, FeedPostConn } from '@portal/posts';

const Container = styled.div`
  display: flex;
  flex-direction: column;
  padding: 50px;
  align-items: center;
`;

export class HomePage extends Component {
  componentDidMount() {
    const { fetchFeedPosts } = this.props;
    setTimeout(fetchFeedPosts, 100);
  }

  render() {
    const { fetchFeedPosts, posts } = this.props;
    return h(Container, [
      h(
        InfiniteScroll,
        {
          dataLength: posts.length,
          next: () => fetchFeedPosts(),
          hasMore: false,
        },
        posts.map((post) =>
          h(FeedPostConn, { key: get(post, 'postId'), post }),
        ),
      ),
    ]);
  }
}

const mapStateToProps = (state) => ({
  posts: getFeedPosts(state),
});

const mapDispatchToProps = (dispatch) => ({
  fetchFeedPosts: () => dispatch(fetchFeedAction()),
});

export const HomePageConn = connect(
  mapStateToProps,
  mapDispatchToProps,
)(HomePage);
