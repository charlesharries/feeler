import React, { useState, useEffect } from 'react';
import styled, { createGlobalStyle } from 'styled-components';
import Sentiment from './Sentiment';
import { Container, Cover, Input, Stack } from '../styles';
import Intro from './Intro';
import Footer from './Footer';
import { black } from '../styles/colors';

const GlobalStyle = createGlobalStyle`
  body {
    margin: 0;
    padding: 10px;
  }

  * {
    box-sizing: border-box;
    margin: 0;
  }
`;

const Wrapper = styled.main`
  font-family: -apple-system, BlinkMacSystemFont, avenir next, avenir,
    helvetica neue, helvetica, Ubuntu, roboto, noto, segoe ui, arial, sans-serif;
  background-color: ${black};
  color: #efefef;
  height: calc(100vh - 20px);
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;

  a {
    color: inherit;
  }
`;

function App() {
  const [phrase, setPhrase] = useState('');
  const [sentiment, setSentiment] = useState(null);

  useEffect(() => {
    async function getSentiment() {
      const s = await fetch('/sentiments', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ s: phrase }),
      }).then((r) => r.json());

      setSentiment(s);
    }

    getSentiment();
  }, [phrase]);

  return (
    <Wrapper>
      <GlobalStyle />
      <Cover>
        <Container className="cover__main" maxWidth={450}>
          <Stack align="center">
            <Intro />

            <Input>
              <input
                type="text"
                name="phrase"
                onChange={(e) => setPhrase(e.target.value)}
                placeholder="Enter some text..."
              />
            </Input>

            <Sentiment sentiment={sentiment} />
          </Stack>
        </Container>

        <Footer />
      </Cover>
    </Wrapper>
  );
}

export default App;
