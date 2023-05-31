import React from 'react';
import { Provider } from 'react-redux';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { AppProps } from 'next/app';
import { NextComponentType, NextPageContext } from 'next'
import { Router } from 'next/router';

import store from '../store';
import { useMode } from '../styles/theme';

export default function MyApp({ Component, pageProps, router }: AppProps & { router: Router }) {
  return (
    <Provider store={store}>
      <AppWrapper Component={Component} pageProps={pageProps} router={router} />
    </Provider>
  );
}

interface AppWrapperProps {
  Component: NextComponentType<NextPageContext>;
  pageProps: any;
  router: Router;
}

function AppWrapper({ Component, pageProps, router }: AppWrapperProps) {
  const [theme, colorMode] = useMode();

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Component {...pageProps} router={router} />
    </ThemeProvider>
  );
}
