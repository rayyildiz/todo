import React from 'react';
import {BrowserRouter} from "react-router-dom";
import {Container, createStyles, CssBaseline, makeStyles, Theme, ThemeProvider} from "@material-ui/core";
import {AppRoutes} from "./App.Routes";
import {Header} from "./Components/Header";
import {createMuiTheme} from '@material-ui/core/styles';
import {deepOrange, red} from "@material-ui/core/colors";
import {ApolloClient, ApolloProvider, createHttpLink, InMemoryCache} from "@apollo/client";
import {BASE_API, LOCALSTORAGE_AUTH_KEY} from "./environment";
import {setContext} from "@apollo/client/link/context";


const theme = createMuiTheme({
  palette: {
    primary: red,
    secondary: deepOrange,
  },
});

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
      main: {
        paddingLeft: 0,
        paddingRight: 0
      },
      paper: {
        marginTop: theme.spacing(2),
        display: 'flex',
        flexDirection: 'column',
        marginRight: theme.spacing(4),
        marginLeft: theme.spacing(4),
      },
    }));

const httpLink = createHttpLink({
  uri: BASE_API,
});

const authLink = setContext((_, {headers}) => {
  const token = localStorage.getItem(LOCALSTORAGE_AUTH_KEY);
  return {
    headers: {
      ...headers,
      Authorization: token ? `${token}` : "",
    }
  }
});

const client = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: new InMemoryCache()
});
/*
 const token = localStorage.getItem(LOCALSTORAGE_TOKEN);
  if (token != null) {
    request.headers.set("Authorization", "Bearer " + token);
  }
 */

const App: React.FC = () => {
  const classes = useStyles();

  return (
      <ApolloProvider client={client}>
        <BrowserRouter>
          <Container component="main" className={classes.main} maxWidth="xl">
            <CssBaseline/>
            <ThemeProvider theme={theme}>
              <Header/>
              <div className={classes.paper}>
                <AppRoutes/>
              </div>
            </ThemeProvider>
          </Container>
        </BrowserRouter>
      </ApolloProvider>
  );
};

export default App;
