import React from 'react';
import ReactDOM from 'react-dom';
import {ChakraProvider, ColorModeScript, theme} from '@chakra-ui/react';
import App from './App';
import * as serviceWorker from './serviceWorker';
import 'typeface-roboto';
import {ENABLE_SW} from "./environment";
import './index.css';


ReactDOM.render(
    <React.StrictMode>
      <ColorModeScript/>
      <ChakraProvider theme={theme}>
        <App/>
      </ChakraProvider>
    </React.StrictMode>,
    document.getElementById('root')
);

if (ENABLE_SW) {
  serviceWorker.register();
} else {
  serviceWorker.unregister();
}
