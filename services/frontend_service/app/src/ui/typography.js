import styled from 'styled-components';

const Base = styled.div`
  white-space: pre-line;
  font-family: ${({ theme }) => theme.primaryFont}, sans-serif;
`;

export const Title = Base.extend`
  margin-bottom: 32px;
  font-size: 40px;
  line-height: 48px;
  letter-spacing: 0.6px;
  font-weight: 300;
  color: ${({ theme }) => theme.primary};
`;

export const Subtitle = Base.extend`
  margin-bottom: 26px;
  font-size: 32px;
  line-height: 38px;
  letter-spacing: 0.2px;
  font-weight: 500;
  color: ${({ theme }) => theme.primary};
`;

export const Body = Base.extend`
  font-size: 15px;
  line-height: 24px;
  font-weight: 400;
  color: ${({ theme }) => theme.primary};
`;

export const BodyFaded = Body.extend`
  color: ${({ theme }) => theme.foregroundSecondary};
`;

export const Text = Base.extend`
  font-size: 13px;
  line-height: 21px;
  font-weight: 400;
  color: ${({ theme }) => theme.primary};
`;

export const TextFaded = Text.extend`
  color: ${({ theme }) => theme.secondary};
`;
