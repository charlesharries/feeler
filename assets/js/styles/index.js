import styled from 'styled-components';
import Input from './Input';

const Stack = styled.div`
  display: flex;
  flex-direction: column;
  ${({ align }) => {
    if (align === 'center') return 'text-align: center;';
    if (align === 'right') return 'text-align: right;';
  }}

  & > * {
    margin: 0;
  }

  & > * + * {
    margin-top: ${({ size }) => {
      const sizes = { sm: 5, md: 10, lg: 15, xl: 20 };

      return `${sizes[size || 'md']}px`;
    }};
  }
`;

const Cover = styled.div`
  --space: 15px;

  display: flex;
  flex-direction: column;
  min-height: 100%;
  padding: var(--space);

  & > * {
    margin-top: 1rem;
    margin-bottom: 1rem;
  }

  & > :first-child:not(.cover__main) {
    margin-top: 0;
  }

  & > :last-child:not(.cover__main) {
    margin-bottom: 0;
  }

  & > .cover__main {
    margin-top: auto;
    margin-bottom: auto;
  }
`;

const Container = styled.div`
  max-width: ${({ maxWidth }) => {
    if (!maxWidth) return '100%';
    return `${maxWidth}px`;
  }};
  padding-left: 20px;
  padding-right: 20px;
`;

const Cluster = styled.div`
  display: flex;
  flex-flow: row wrap;
  margin-right: -10px;
  margin-bottom: -10px;

  & > * {
    margin-right: 10px;
    margin-bottom: 10px;
  }
`;

const Text = styled.p`
  text-align: ${({ align }) => align};
  font-weight: ${({ fontWeight }) => fontWeight};
  font-size: ${({ size }) =>
    `${{ sm: 0.85, md: 1, lg: 1.25 }[size || 'md']}rem`};
`;

export { Cluster, Container, Cover, Input, Stack, Text };
