import React from 'react';
import { Stack } from '../styles';

export default function Intro() {
  return (
    <Stack align="center">
      <h1>Sentiment Analysis</h1>
      <p>
        Ultra-basic sentiment analysis based on{' '}
        <a
          target="_blank"
          rel="noreferrer noopener"
          href="https://github.com/fnielsen/afinn/blob/master/afinn/data/AFINN-en-165.txt"
        >
          AFINN-165
        </a>
        .
      </p>
      <p>
        Check out the full source code{' '}
        <a
          target="_blank"
          rel="noreferrer noopener"
          href="https://github.com/charlesharries/feeler"
        >
          here
        </a>
        .
      </p>
    </Stack>
  );
}
