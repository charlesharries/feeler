import React, {useState, useEffect} from 'react';

function App() {
  const [phrase, setPhrase] = useState('');
  const [sentiment, setSentiment] = useState(null)

  useEffect(() => {
    getSentiment();

    async function getSentiment() {
      const s = await fetch('/sentiments', {
        method: 'POST',
        headers: {'Content-Type': "application/json"},
        body: JSON.stringify({ s: phrase }),
      }).then(r => r.json())

      setSentiment(s)
    }
  }, [phrase])

  return <>
    <input type="text" onChange={(e) => setPhrase(e.target.value)} />
    <p>Sentiment: {JSON.stringify(sentiment)}</p>
  </>
}

export default App