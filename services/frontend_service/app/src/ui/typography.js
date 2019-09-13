import styled from 'styled-components';

const Base = styled.div`
  white-space: pre-line;
  font-family: ${({ theme }) => theme.primaryFont}, sans-serif;
`;

export const Title = styled(Base)`
  margin-bottom: 32px;
  font-size: 40px;
  line-height: 48px;
  letter-spacing: 0.6px;
  font-weight: 300;
  color: ${({ theme }) => theme.primary};
`;

export const Subtitle = styled(Base)`
  margin-bottom: 26px;
  font-size: 32px;
  line-height: 38px;
  letter-spacing: 0.2px;
  font-weight: 500;
  color: ${({ theme }) => theme.primary};
`;

export const Body = styled(Base)`
  font-size: 15px;
  line-height: 24px;
  font-weight: 400;
  color: ${({ theme }) => theme.primary};
`;

export const BodyFaded = styled(Body)`
  color: ${({ theme }) => theme.foregroundSecondary};
`;

export const Text = styled(Base)`
  font-size: 13px;
  line-height: 21px;
  font-weight: 400;
  color: ${({ theme }) => theme.primary};
`;

export const TextFaded = styled(Text)`
  color: ${({ theme }) => theme.secondary};
`;
