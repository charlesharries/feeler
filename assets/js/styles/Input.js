import styled from 'styled-components';
import { white, yellow, darkGrey } from './colors';

export default styled.div`
  width: 100%;

  label {
    display: block;
    font-weight: 700;

    & + input {
      margin-top: 5px;
    }
  }

  input {
    padding: 15px;
    color: ${white};
    border: 2px solid transparent;
    border-radius: 4px;
    width: 100%;
    display: block;
    font-size: 16px;
    background-color: ${darkGrey};
    transition: border 0.2s ease;

    &:active,
    &:focus {
      outline: none;
      border: 2px solid ${yellow};
    }
  }
`;
