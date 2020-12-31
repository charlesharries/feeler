import React from 'react';
import styled from 'styled-components';
import { darkGrey, green, red } from '../styles/colors';

const Pill = styled.span`
  border-radius: 1000px;
  background-color: ${darkGrey};
  color: ${({ positive }) => (positive ? green : red)};
  padding: 3px 10px 6px;
  font-size: 0.85rem;
`;

export default function Tag({ children, positive }) {
  return <Pill positive={positive}>{children}</Pill>;
}
