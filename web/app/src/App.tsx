import React from 'react';
import {BrowserRouter} from "react-router-dom";
import {AppRoutes} from "./App.Routes";
import {ApolloClient, ApolloProvider, createHttpLink, InMemoryCache} from "@apollo/client";
import {BASE_API, LOCALSTORAGE_AUTH_KEY} from "./environment";
import {setContext} from "@apollo/client/link/context";
import {Layout} from "./components/Layout";


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

const App: React.FC = () => {

  return (
      <ApolloProvider client={client}>
        <BrowserRouter>
          <Layout>
            <AppRoutes/>
          </Layout>
        </BrowserRouter>
      </ApolloProvider>
  );
};

export default App;
