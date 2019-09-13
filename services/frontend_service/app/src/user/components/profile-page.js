import { div } from 'react-hyperscript-helpers';
import { connect } from 'react-redux';

export const ProfilePage = () => div('Post');

// import { Component } from 'react';
// import { h, div } from 'react-hyperscript-helpers';
// import InfiniteScroll from 'react-infinite-scroll-component';

// import { FeedPost } from './feed-post';

// export class ProfilePage extends Component {

//   render() {
//     const { fetchFeed } = this.props;
//     return div(
//       h(
//         InfiniteScroll,
//         {
//           dataLength: posts.length,
//           next: () => fetchFeed(page + 1),
//           hasMore: next !== '',
//         },
//         posts.map(({ postId }) => h(FeedPost, { key: postId, postId })),
//       ),
//     );
//   }
// }

// function mapDispatchToProps(dispatch) {
//   return {
//     fetchFeed: (page) =>
//       dispatch(
//         fetchFeedAction({
//           page,
//         }),
//       ),
//   };
// }

export const ProfilePageConn = connect()(ProfilePage);
// mapStateToProps,
// mapDispatchToProps,
