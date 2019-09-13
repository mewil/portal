import { h } from 'react-hyperscript-helpers';
import { Component } from 'react';
import styled from 'styled-components';

const Group = styled.div`
  display: flex;
  flex-direction: row;
  padding: 0;
  margin: 0;
`;

const TabItem = styled.div`
  flex: 1;
  text-align: center;
  background-color: ${({ active, activeColor }) =>
    active ? activeColor : 'transparent'};
  border: 3px solid ${({ activeColor }) => activeColor};
  border-right: none;
  color: ${({ active, activeColor }) => (active ? 'white' : activeColor)};
  padding: 10px 0;
  font-size: 15px;
  font-weight: 500;
  &:first-child {
    border-top-left-radius: 3px;
    border-bottom-left-radius: 3px;
  }
  &:last-child {
    border-top-right-radius: 3px;
    border-bottom-right-radius: 3px;
    border-right: 3px solid ${({ activeColor }) => activeColor};
  }
  &:hover {
    background-color: ${({ activeColor }) => activeColor};
    color: white;
  }
`;

export class TabGroup extends Component {
  constructor(props) {
    super(props);
    this.state = {
      activeIndex: props.defaultIndex || 0,
    };
  }

  clickMiddleware(index, func) {
    this.setState({
      activeIndex: index,
    });
    func(index);
  }

  render() {
    const { tabs = [], primaryColor } = this.props;
    const { activeIndex } = this.state;
    return h(Group, [
      tabs.map((tab, i) =>
        h(
          TabItem,
          {
            key: i,
            onClick: () => this.clickMiddleware(i, tab.onClick),
            activeColor: primaryColor,
            active: i === activeIndex,
          },
          [`${tab.title}`],
        ),
      ),
    ]);
  }
}
