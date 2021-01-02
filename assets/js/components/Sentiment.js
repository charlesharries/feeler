import React from 'react';
import styled, { createGlobalStyle } from 'styled-components';
import { Cluster, Stack, Text } from '../styles';
import Tag from './Tag';
import { green, red } from '../styles/colors';

const GlobalStyle = createGlobalStyle`
  body {
    transition: background-color 0.2s ease;
    background-color: ${({ score }) => {
      if (score > 0) return green;
      if (score < 0) return red;

      return 'transparent';
    }}
  }
`;

const SentimentResult = styled.div`
  transition: color 0.2s ease;
  color: ${({ score }) => {
    if (score === 0) return 'inherit';

    return score > 0 ? green : red;
  }};
`;

export default function Sentiment({ sentiment }) {
  if (!sentiment) return null;

  return (
    <>
      <GlobalStyle score={sentiment.score} />

      <Stack size="lg">
        <SentimentResult score={sentiment.score}>
          <Text fontWeight="700">{sentiment.verdict}</Text>
        </SentimentResult>

        <Cluster>
          {sentiment.positive_words.map((word, i) => (
            <Tag key={`pos-${word}-${i}`} positive>
              {word}
            </Tag>
          ))}

          {sentiment.negative_words.map((word, i) => (
            <Tag key={`neg-${word}-${i}`} positive={false}>
              {word}
            </Tag>
          ))}
        </Cluster>
      </Stack>
    </>
  );
}
